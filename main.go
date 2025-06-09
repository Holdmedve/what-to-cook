package main

import (
	"errors"
	"html/template"
	"net/http"
	"os"

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

		var recipes []Recipe = GetAllRecipes()

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

			DeleteRecipe(titleOfRecipeToDelete)

			http.Redirect(w, r, "/", http.StatusSeeOther)
		} else if r.FormValue("editBtn") != "" {
			var recipes []Recipe = GetAllRecipes()

			vars := mux.Vars(r)
			titleOfRecipeToEdit := vars["title"]

			var recipeViews []RecipeView
			for _, r := range recipes {
				edit := r.Title == titleOfRecipeToEdit
				recipeViews = append(recipeViews, RecipeView{Recipe: r, Edit: edit})
			}

			tmpl.Execute(w, FormData{RecipeViews: recipeViews})
		} else if r.FormValue("updateBtn") != "" {

			oldTitleOfRecipeToUpdate := r.FormValue("oldTitle")
			DeleteRecipe(oldTitleOfRecipeToUpdate)

			updatedRecipe := Recipe{Title: r.FormValue("editTitle"), Description: r.FormValue("editDescription")}
			SaveRecipe(updatedRecipe)

			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	})

	http.ListenAndServe("127.0.0.1:80", r)
}
