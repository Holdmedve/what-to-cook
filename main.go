package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

// start version controlling!
// use playwright

type Recipe struct {
	Title       string
	Description string
}

type TodoPageData struct {
	PageTitle string
	Todos     []Recipe
}

type ContactDetails struct {
	Email   string
	Subject string
	Message string
}

func main() {
	tmpl := template.Must(template.ParseFiles("./static/forms.html"))
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	data := TodoPageData{
	// 		PageTitle: "My TODO list",
	// 		Todos: []Recipe{
	// 			{Name: "Rec 1", Description: "om nom nom"},
	// 			{Name: "Rec 2", Description: "om nom nom"},
	// 			{Name: "Rec 3", Description: "om nom nom"},
	// 		},
	// 	}
	// 	tmpl.Execute(w, data)
	// })

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}

		recipe := Recipe{
			Title:       r.FormValue("title"),
			Description: r.FormValue("description"),
		}

		filename := "./test.txt"

		formattedRecipe := fmt.Sprintf("%s\t%s\n", recipe.Title, recipe.Description)
		dataToSave := []byte(formattedRecipe)

		f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		if err != nil {
			panic(err)
		}

		defer f.Close()

		if _, err = f.WriteString(string(dataToSave)); err != nil {
			panic(err)
		}

		tmpl.Execute(w, struct{ Success bool }{true})
	})

	http.ListenAndServe("127.0.0.1:80", nil)
}
