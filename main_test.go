package main

import (
	"log"
	"os"
	"testing"

	"github.com/playwright-community/playwright-go"
)

func assertErrorToNilf(message string, err error) {
	if err != nil {
		log.Fatalf(message, err)
	}
}

func TestThatFails(t *testing.T) {
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
