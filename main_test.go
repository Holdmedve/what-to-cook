package main

import (
	"log"
	"os"
	"testing"
)

// TODO create config logic that relies on an ENV env var
// if set to prod use recipes.txt otherwise test-recipes.txt
// ...or don't run the tests on prod
const recipeFilePath = "recipes.txt"

func TestSubmittedRecipeIsSaved(t *testing.T) {
	page := setup(recipeFilePath)
	defer teardown(page, recipeFilePath)

	saveRecipe(page, "test-recipe-name", "test-recipe-description")

	b, err := os.ReadFile(recipeFilePath)
	if err != nil {
		log.Fatal(err)
	}
	storedRecipe := string(b)
	expectedRecipe := "test-recipe-name\ttest-recipe-description\n"
	if storedRecipe != expectedRecipe {
		log.Fatalf("expected: %s, actual: %s", expectedRecipe, storedRecipe)
	}
}

func Test_Submit2Recipes_BothAreDisplayed(t *testing.T) {
	page := setup(recipeFilePath)
	defer teardown(page, recipeFilePath)
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
	page := setup(recipeFilePath)
	defer teardown(page, recipeFilePath)
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
	page := setup(recipeFilePath)
	defer teardown(page, recipeFilePath)
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
