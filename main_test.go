package main

import (
	"log"
	"os"
	"testing"
)

const recipeFilePath = "test.txt"

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
	assertErrorToNilf("could not locate and fill: %w", err)
	deleteBtns[1].Click()

	titleLocator, err := page.Locator("[name='displayTitle']").All()
	if len(titleLocator) != 2 {
		log.Fatalf("there should be exactly 2 titles, found: %d", len(titleLocator))
	}
	assertRecipeIsDisplayed(page, testRecipeTitleA, testRecipeDescriptionA)
	assertRecipeIsDisplayed(page, testRecipeTitleC, testRecipeDescriptionC)
}
