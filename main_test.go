package main

import (
	"log"
	"os"
	"strings"
	"testing"

	"github.com/playwright-community/playwright-go"
)

func assertErrorToNilf(message string, err error) {
	if err != nil {
		log.Fatalf(message, err)
	}
}

func TestSubmittedRecipeIsSaved(t *testing.T) {
	recipeFilePath := "test.txt"
	file, err := os.Create(recipeFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	pw, err := playwright.Run()
	assertErrorToNilf("could not launch playwright: %w", err)
	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(false),
	})
	assertErrorToNilf("could not launch Chromium: %w", err)
	context, err := browser.NewContext()
	assertErrorToNilf("could not create context: %w", err)
	page, err := context.NewPage()
	assertErrorToNilf("could not create page: %w", err)
	_, err = page.Goto("127.0.0.1:80")
	assertErrorToNilf("could not goto: %w", err)
	err = page.Locator("[name='title']").Fill("test-recipe-name")
	assertErrorToNilf("could not locate and fill: %w", err)
	err = page.Locator("[name='description']").Fill("test-recipe-description")
	assertErrorToNilf("could not locate and fill: %w", err)

	err = page.Locator("[name='submit']").Click()
	assertErrorToNilf("could not locate and click: %w", err)

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
	recipeFilePath := "test.txt"
	file, err := os.Create(recipeFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	defer os.Create(recipeFilePath)
	pw, err := playwright.Run()
	assertErrorToNilf("could not launch playwright: %w", err)
	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(false),
	})
	assertErrorToNilf("could not launch Chromium: %w", err)
	context, err := browser.NewContext()
	assertErrorToNilf("could not create context: %w", err)
	page, err := context.NewPage()
	assertErrorToNilf("could not create page: %w", err)
	_, err = page.Goto("127.0.0.1:80")
	assertErrorToNilf("could not goto: %w", err)
	testRecipeTitleA := "testRecipeTitleA"
	testRecipeDescriptionA := "testRecipeDescriptioncA"
	testRecipeTitleB := "testRecipeTitleB"
	testRecipeDescriptionB := "testRecipeDescriptioncB"

	err = page.Locator("[name='title']").Fill(testRecipeTitleA)
	assertErrorToNilf("could not locate and fill: %w", err)
	err = page.Locator("[name='description']").Fill(testRecipeDescriptionA)
	assertErrorToNilf("could not locate and fill: %w", err)
	err = page.Locator("[name='submit']").Click()
	assertErrorToNilf("could not locate and click: %w", err)

	err = page.Locator("[name='title']").Fill(testRecipeTitleB)
	assertErrorToNilf("could not locate and fill: %w", err)
	err = page.Locator("[name='description']").Fill(testRecipeDescriptionB)
	assertErrorToNilf("could not locate and fill: %w", err)
	err = page.Locator("[name='submit']").Click()
	assertErrorToNilf("could not locate and click: %w", err)

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
	recipeFilePath := "test.txt"
	file, err := os.Create(recipeFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	defer os.Create(recipeFilePath)
	pw, err := playwright.Run()
	assertErrorToNilf("could not launch playwright: %w", err)
	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(false),
	})
	assertErrorToNilf("could not launch Chromium: %w", err)
	context, err := browser.NewContext()
	assertErrorToNilf("could not create context: %w", err)
	page, err := context.NewPage()
	assertErrorToNilf("could not create page: %w", err)
	_, err = page.Goto("127.0.0.1:80")
	assertErrorToNilf("could not goto: %w", err)

	testRecipeTitleA := "testRecipeTitleA"
	testRecipeDescriptionA := "testRecipeDescriptioncA"
	testRecipeTitleB := "testRecipeTitleB"
	testRecipeDescriptionB := "testRecipeDescriptioncB"
	testRecipeTitleC := "testRecipeTitleC"
	testRecipeDescriptionC := "testRecipeDescriptioncC"

	err = page.Locator("[name='title']").Fill(testRecipeTitleA)
	assertErrorToNilf("could not locate and fill: %w", err)
	err = page.Locator("[name='description']").Fill(testRecipeDescriptionA)
	assertErrorToNilf("could not locate and fill: %w", err)
	err = page.Locator("[name='submit']").Click()
	assertErrorToNilf("could not locate and click: %w", err)

	err = page.Locator("[name='title']").Fill(testRecipeTitleB)
	assertErrorToNilf("could not locate and fill: %w", err)
	err = page.Locator("[name='description']").Fill(testRecipeDescriptionB)
	assertErrorToNilf("could not locate and fill: %w", err)
	err = page.Locator("[name='submit']").Click()
	assertErrorToNilf("could not locate and click: %w", err)

	err = page.Locator("[name='title']").Fill(testRecipeTitleC)
	assertErrorToNilf("could not locate and fill: %w", err)
	err = page.Locator("[name='description']").Fill(testRecipeDescriptionC)
	assertErrorToNilf("could not locate and fill: %w", err)
	err = page.Locator("[name='submit']").Click()
	assertErrorToNilf("could not locate and click: %w", err)

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
