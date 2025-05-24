package main

import (
	"log"
	"os"
	"slices"
	"strings"

	"github.com/playwright-community/playwright-go"
)

func assertErrorToNilf(message string, err error) {
	if err != nil {
		log.Fatalf(message, err)
	}
}

func fillPage(page playwright.Page, selector string, value string) {
	err := page.Locator(selector).Fill(value)
	if err != nil {
		log.Fatalf("error during fill using selector %s and value %s: %s", selector, value, err.Error())
	}
}

func clickPage(page playwright.Page, selector string) {
	err := page.Locator(selector).Click()
	if err != nil {
		log.Fatalf("error during click using selector %s: %s", selector, err.Error())
	}
}

func setupPage() playwright.Page {
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

	return page
}

func setupRecipeFile(filePath string) {
	file, err := os.Create(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
}

func saveRecipe(page playwright.Page, title string, descr string) {
	fillPage(page, "[name='title']", title)
	fillPage(page, "[name='description']", descr)
	clickPage(page, "[name='submit']")
}

func setup(recipeFilePath string) playwright.Page {
	setupRecipeFile(recipeFilePath)
	page := setupPage()
	_, err := page.Goto("127.0.0.1:80")
	if err != nil {
		log.Fatal(err)
	}
	return page
}

func teardown(page playwright.Page, recipeFilePath string) {
	page.Close()
	setupRecipeFile(recipeFilePath)
}

func assertRecipeIsDisplayed(page playwright.Page, expectedTitle string, expectedDescr string) {
	titleLocator, _ := page.Locator("[name='displayTitle']").All()
	actualTitles := make([]string, 0)
	for _, title := range titleLocator {
		t, _ := title.TextContent()
		actualTitles = append(actualTitles, t)
	}
	containsTitle := slices.ContainsFunc(actualTitles, func(s string) bool {
		return strings.Contains(s, expectedTitle)
	})
	if !containsTitle {
		log.Fatalf("expected title %s was not found", expectedTitle)
	}

	descrLocator, _ := page.Locator("[name='displayDescription']").All()
	actualDescriptions := make([]string, 0)
	for _, descr := range descrLocator {
		d, _ := descr.TextContent()
		actualDescriptions = append(actualDescriptions, d)
	}
	containsDescr := slices.ContainsFunc(actualDescriptions, func(s string) bool {
		return strings.Contains(s, expectedDescr)
	})
	if !containsDescr {
		log.Fatalf("expected description %s was not found", expectedDescr)
	}
}
