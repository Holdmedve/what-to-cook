package main

import (
	"errors"
	"html/template"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
)

type RecipeView struct {
	Recipe
	Edit bool
}

type FormData struct {
	RecipeViews []RecipeView
}

func main() {
	if _, err := os.Stat(RecipesFilePath); errors.Is(err, os.ErrNotExist) {
		os.Create(RecipesFilePath)
	}

	tmpl := template.Must(template.ParseFiles("./static/forms.html"))

	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost && r.FormValue("title") != "" && r.FormValue("description") != "" {
			recipe := Recipe{
				Title:       r.FormValue("title"),
				Description: r.FormValue("description"),
			}
			SaveRecipe(recipe)
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}

		b, err := os.ReadFile(RecipesFilePath)
		if err != nil {
			panic(err)
		}
		rawRecipes := strings.Split(string(b), "\n")
		rawRecipes = rawRecipes[:len(rawRecipes)-1]
		var recipes []Recipe
		for _, raw := range rawRecipes {
			rawParts := strings.Split(raw, "\t")
			recipes = append(recipes, Recipe{Title: rawParts[0], Description: rawParts[1]})
		}
		var recipeViews []RecipeView
		for _, r := range recipes {
			recipeViews = append(recipeViews, RecipeView{Recipe: r, Edit: false})
		}

		tmpl.Execute(w, FormData{RecipeViews: recipeViews})
	})

	r.HandleFunc("/recipes/{title}", func(w http.ResponseWriter, r *http.Request) {
		if r.FormValue("deleteBtn") != "" {
			vars := mux.Vars(r)
			titleOfRecipeToDelete := vars["title"]

			b, err := os.ReadFile(RecipesFilePath)
			if err != nil {
				panic(err)
			}

			recipes := strings.Split(string(b), "\n")
			recipes = recipes[:len(recipes)-1]
			var updatedRecipes string
			for _, r := range recipes {
				title := strings.Split(r, "\t")[0]
				if title == titleOfRecipeToDelete {
					continue
				}
				updatedRecipes += r + "\n"
			}

			os.WriteFile(RecipesFilePath, []byte(updatedRecipes), 0222)

			http.Redirect(w, r, "/", http.StatusSeeOther)
		} else if r.FormValue("editBtn") != "" {
			b, err := os.ReadFile(RecipesFilePath)
			if err != nil {
				panic(err)
			}

			rawRecipes := strings.Split(string(b), "\n")
			rawRecipes = rawRecipes[:len(rawRecipes)-1]
			var recipes []Recipe
			for _, raw := range rawRecipes {
				rawParts := strings.Split(raw, "\t")
				recipes = append(recipes, Recipe{Title: rawParts[0], Description: rawParts[1]})
			}

			vars := mux.Vars(r)
			titleOfRecipeToEdit := vars["title"]
			var recipeViews []RecipeView
			for _, r := range recipes {
				edit := r.Title == titleOfRecipeToEdit
				recipeViews = append(recipeViews, RecipeView{Recipe: r, Edit: edit})
			}

			tmpl.Execute(w, FormData{RecipeViews: recipeViews})
		} else if r.FormValue("updateBtn") != "" {
			b, err := os.ReadFile(RecipesFilePath)
			if err != nil {
				panic(err)
			}

			rawRecipes := strings.Split(string(b), "\n")
			rawRecipes = rawRecipes[:len(rawRecipes)-1]
			var recipes []Recipe
			for _, raw := range rawRecipes {
				rawParts := strings.Split(raw, "\t")
				recipes = append(recipes, Recipe{Title: rawParts[0], Description: rawParts[1]})
			}

			oldTitleOfRecipeToUpdate := r.FormValue("oldTitle")
			for idx, rec := range recipes {
				if rec.Title == oldTitleOfRecipeToUpdate {
					recipes[idx].Title = r.FormValue("editTitle")
					recipes[idx].Description = r.FormValue("editDescription")
					break
				}
			}

			updatedRecipes := ""
			for _, r := range recipes {
				updatedRecipes += r.Title + "\t" + r.Description + "\n"
			}
			os.WriteFile(RecipesFilePath, []byte(updatedRecipes), 0222)

			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	})

	http.ListenAndServe("127.0.0.1:80", r)
}
