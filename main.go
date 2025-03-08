package main

import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
)

type Recipe struct {
	Title       string
	Description string
}

type FormData struct {
	Recipes []Recipe
}

func main() {
	filename := "./test.txt"
	if _, err := os.Stat(filename); errors.Is(err, os.ErrNotExist) {
		os.Create(filename)
	}

	tmpl := template.Must(template.ParseFiles("./static/forms.html"))

	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost && r.FormValue("title") != "" && r.FormValue("description") != "" {
			currentRecipes, err := os.ReadFile(filename)
			if err != nil {
				log.Fatal(err)
			}

			recipe := Recipe{
				Title:       r.FormValue("title"),
				Description: r.FormValue("description"),
			}
			newRecipe := fmt.Sprintf("%s\t%s\n", recipe.Title, recipe.Description)
			updatedRecipes := append(currentRecipes, []byte(newRecipe)...)
			os.WriteFile(filename, []byte(updatedRecipes), 0644)
			http.Redirect(w, r, "/", 303)
		}

		b, err := os.ReadFile(filename)
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

		tmpl.Execute(w, FormData{Recipes: recipes})
	})

	r.HandleFunc("/recipes/{title}", func(w http.ResponseWriter, r *http.Request) {
		if r.FormValue("_method") != "delete" {
			w.Write([]byte("I was expecting _method to be equal to delete"))
			return
		}

		vars := mux.Vars(r)
		titleOfRecipeToDelete := vars["title"]

		b, err := os.ReadFile(filename)
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

		os.WriteFile(filename, []byte(updatedRecipes), 0222)

		http.Redirect(w, r, "/", 303)
	})

	http.ListenAndServe("127.0.0.1:80", r)
}
