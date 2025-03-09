package main

import (
	"log"
	"os"
	"strings"
	"testing"
)

const recipeFilePath = "test.txt"

func TestSubmittedRecipeIsSaved(t *testing.T) {
	setupRecipeFile(recipeFilePath)
	page := setupPage()
	_, err := page.Goto("127.0.0.1:80")
	assertErrorToNilf("could not goto: %w", err)
	fillPage(page, "[name='title']", "test-recipe-name")
	fillPage(page, "[name='description']", "test-recipe-description")

	clickPage(page, "[name='submit']")

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
	setupRecipeFile(recipeFilePath)
	page := setupPage()
	_, err := page.Goto("127.0.0.1:80")
	assertErrorToNilf("could not goto: %w", err)
	testRecipeTitleA := "testRecipeTitleA"
	testRecipeDescriptionA := "testRecipeDescriptioncA"
	testRecipeTitleB := "testRecipeTitleB"
	testRecipeDescriptionB := "testRecipeDescriptioncB"

	fillPage(page, "[name='title']", testRecipeTitleA)
	fillPage(page, "[name='description']", testRecipeDescriptionA)
	clickPage(page, "[name='submit']")

	fillPage(page, "[name='title']", testRecipeTitleB)
	fillPage(page, "[name='description']", testRecipeDescriptionB)
	clickPage(page, "[name='submit']")

	titleLocator, err := page.Locator("[name='displayTitle']").All()
	if len(titleLocator) != 2 {
		log.Fatalf("there should be exactly 2 titles, found: %d", len(titleLocator))
	}

	actualRecipeTitleA, _ := titleLocator[0].TextContent()
	actualRecipeTitleB, _ := titleLocator[1].TextContent()
	if strings.Contains(actualRecipeTitleA, testRecipeTitleA) == false {
		log.Fatalf("expected: %s, to contain: %s", actualRecipeTitleA, testRecipeTitleA)
	}
	if strings.Contains(actualRecipeTitleB, testRecipeTitleB) == false {
		log.Fatalf("expected: %s, to contain: %s", actualRecipeTitleB, testRecipeTitleB)
	}

	descriptionLocator, err := page.Locator("[name='displayDescription']").All()
	if len(descriptionLocator) != 2 {
		log.Fatalf("there should be exactly 2 titles, found: %d", len(descriptionLocator))
	}

	actualRecipeDescriptionA, _ := descriptionLocator[0].TextContent()
	actualRecipeDescriptionB, _ := descriptionLocator[1].TextContent()
	if strings.Contains(actualRecipeDescriptionA, testRecipeDescriptionA) == false {
		log.Fatalf("expected: %s, to contain: %s", actualRecipeDescriptionA, testRecipeDescriptionA)
	}
	if strings.Contains(actualRecipeDescriptionB, testRecipeDescriptionB) == false {
		log.Fatalf("expected: %s, to contain: %s", actualRecipeDescriptionB, testRecipeDescriptionB)
	}
}

func Test_Submit3RecipesDelete2nd_1stAnd3rdAreDisplayed(t *testing.T) {
	setupRecipeFile(recipeFilePath)
	page := setupPage()
	_, err := page.Goto("127.0.0.1:80")
	assertErrorToNilf("could not goto: %w", err)

	testRecipeTitleA := "testRecipeTitleA"
	testRecipeDescriptionA := "testRecipeDescriptioncA"
	testRecipeTitleB := "testRecipeTitleB"
	testRecipeDescriptionB := "testRecipeDescriptioncB"
	testRecipeTitleC := "testRecipeTitleC"
	testRecipeDescriptionC := "testRecipeDescriptioncC"

	fillPage(page, "[name='title']", testRecipeTitleA)
	fillPage(page, "[name='description']", testRecipeDescriptionA)
	clickPage(page, "[name='submit']")

	fillPage(page, "[name='title']", testRecipeTitleB)
	fillPage(page, "[name='description']", testRecipeDescriptionB)
	clickPage(page, "[name='submit']")

	fillPage(page, "[name='title']", testRecipeTitleC)
	fillPage(page, "[name='description']", testRecipeDescriptionC)
	clickPage(page, "[name='submit']")

	deleteBtns, err := page.Locator("[name='deleteBtn']").All()
	assertErrorToNilf("could not locate and fill: %w", err)
	deleteBtns[1].Click()

	titleLocator, err := page.Locator("[name='displayTitle']").All()
	if len(titleLocator) != 2 {
		log.Fatalf("there should be exactly 2 titles, found: %d", len(titleLocator))
	}

	firstRecipeTitleA, _ := titleLocator[0].TextContent()
	secondRecipeTitleC, _ := titleLocator[1].TextContent()
	if strings.Contains(firstRecipeTitleA, testRecipeTitleA) == false {
		log.Fatalf("expected: %s, to contain: %s", firstRecipeTitleA, testRecipeTitleA)
	}
	if strings.Contains(secondRecipeTitleC, testRecipeTitleC) == false {
		log.Fatalf("expected: %s, to contain: %s", secondRecipeTitleC, testRecipeTitleC)
	}

	descriptionLocator, err := page.Locator("[name='displayDescription']").All()
	if len(descriptionLocator) != 2 {
		log.Fatalf("there should be exactly 2 titles, found: %d", len(descriptionLocator))
	}

	firstRecipeDescriptionA, _ := descriptionLocator[0].TextContent()
	secondRecipeDescriptionC, _ := descriptionLocator[1].TextContent()
	if strings.Contains(firstRecipeDescriptionA, testRecipeDescriptionA) == false {
		log.Fatalf("expected: %s, to contain: %s", firstRecipeDescriptionA, testRecipeDescriptionA)
	}
	if strings.Contains(secondRecipeDescriptionC, testRecipeDescriptionC) == false {
		log.Fatalf("expected: %s, to contain: %s", secondRecipeDescriptionC, testRecipeDescriptionC)
	}
}
