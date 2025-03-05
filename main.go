package main

import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
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

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.FormValue("title") != "" && r.FormValue("description") != "" {
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

	http.ListenAndServe("127.0.0.1:80", nil)
}
