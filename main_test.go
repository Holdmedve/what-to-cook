package main

import (
	"log"
	"testing"
)

func Test_Submit2Recipes_BothAreDisplayed(t *testing.T) {
	page := setup(RecipesFilePath)
	defer teardown(page, RecipesFilePath)
	testRecipeTitleA := "testRecipeTitleA"
	testRecipeDescriptionA := "testRecipeDescriptioncA"
	testRecipeTitleB := "testRecipeTitleB"
	testRecipeDescriptionB := "testRecipeDescriptioncB"

	saveRecipe(page, testRecipeTitleA, testRecipeDescriptionA)
	saveRecipe(page, testRecipeTitleB, testRecipeDescriptionB)

	assertRecipeIsDisplayed(page, testRecipeTitleA, testRecipeDescriptionA)
	assertRecipeIsDisplayed(page, testRecipeTitleB, testRecipeDescriptionB)
}

func Test_Submit3RecipesDelete2nd_1stAnd3rdAreDisplayed(t *testing.T) {
	page := setup(RecipesFilePath)
	defer teardown(page, RecipesFilePath)
	testRecipeTitleA := "testRecipeTitleA"
	testRecipeDescriptionA := "testRecipeDescriptioncA"
	testRecipeTitleB := "testRecipeTitleB"
	testRecipeDescriptionB := "testRecipeDescriptioncB"
	testRecipeTitleC := "testRecipeTitleC"
	testRecipeDescriptionC := "testRecipeDescriptioncC"

	saveRecipe(page, testRecipeTitleA, testRecipeDescriptionA)
	saveRecipe(page, testRecipeTitleB, testRecipeDescriptionB)
	saveRecipe(page, testRecipeTitleC, testRecipeDescriptionC)
	deleteBtns, err := page.Locator("[name='deleteBtn']").All()
	assertErrorToNilf("could not locate: %w", err)
	deleteBtns[1].Click()

	titleLocator, _ := page.Locator("[name='displayTitle']").All()
	if len(titleLocator) != 2 {
		log.Fatalf("there should be exactly 2 titles, found: %d", len(titleLocator))
	}
	assertRecipeIsDisplayed(page, testRecipeTitleA, testRecipeDescriptionA)
	assertRecipeIsDisplayed(page, testRecipeTitleC, testRecipeDescriptionC)
}

func Test_SubmitRecipeThenEditIt_UpdatedValuesAreDisplayed(t *testing.T) {
	page := setup(RecipesFilePath)
	defer teardown(page, RecipesFilePath)
	testRecipeTitle := "testRecipeTitle"
	testRecipeDescription := "testRecipeDescription"
	updatedTestRecipeTitle := "updatedTestRecipeTitle"
	updatedTestRecipeDescription := "updatedTestRecipeDescription"

	saveRecipe(page, testRecipeTitle, testRecipeDescription)
	editBtns, err := page.Locator("[name='editBtn']").All()
	assertErrorToNilf("could not locate: %w", err)
	editBtns[0].Click()
	fillPage(page, "[name='editTitle']", updatedTestRecipeTitle)
	fillPage(page, "[name='editDescription']", updatedTestRecipeDescription)
	clickPage(page, "[name='updateBtn']")

	assertRecipeIsDisplayed(page, updatedTestRecipeTitle, updatedTestRecipeDescription)
}
