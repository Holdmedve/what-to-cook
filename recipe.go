package main

import (
	"log"
	"os"
	"strings"
)

const RecipesFilePath = "recipes.txt"

type Recipe struct {
	Title       string
	Description string
}

func GetAllRecipes() []Recipe {
	recipeBytes, err := os.ReadFile(RecipesFilePath)
	if err != nil {
		log.Fatal(err)
	}

	var recipes []Recipe
	rawRecipes := strings.Split(string(recipeBytes), "\n")
	rawRecipes = rawRecipes[:len(rawRecipes)-1]
	for _, r := range rawRecipes {
		parts := strings.Split(r, "\t")
		recipes = append(recipes, Recipe{Title: parts[0], Description: parts[1]})
	}

	return recipes
}

func truncSaveRecipes(recipes []Recipe) {
	recipesString := ""
	for _, r := range recipes {
		recipesString += r.Title + "\t" + r.Description + "\n"
	}
	os.WriteFile(RecipesFilePath, []byte(recipesString), 0222)
}

func SaveRecipe(recipe Recipe) {
	recipes := GetAllRecipes()
	recipes = append(recipes, recipe)
	truncSaveRecipes(recipes)
}

func DeleteRecipe(titleOfRecipe string) {
	var recipes []Recipe = GetAllRecipes()
	var updatedRecipes []Recipe

	for _, r := range recipes {
		if r.Title == titleOfRecipe {
			continue
		}

		updatedRecipes = append(updatedRecipes, r)
	}

	truncSaveRecipes(updatedRecipes)
}

func UpdateRecipe(oldTitle string, newRecipe Recipe) {
	var recipes []Recipe = GetAllRecipes()

	for idx, r := range recipes {
		if r.Title == oldTitle {
			recipes[idx].Title = newRecipe.Title
			recipes[idx].Description = newRecipe.Description
		}
	}

	truncSaveRecipes(recipes)
}
