package main
import (
 "testing"
 "github.com/ishan2790/go-playwright-tests/pages"
 "github.com/playwright-community/playwright-go"
)

func TestLoginPass(t *testing.T) {
	pw, err := playwright.Run()
	if err != nil {
		t.Fatalf("could not start playwright: %v", err)
	}
	
	chromePath := "C:/Program Files/Google/Chrome/Application/chrome.exe"
	chromePathPtr := &chromePath
	//Set headless to true and remove time_slow for fast execution
	headless := false
	time_slow := 1000.0
	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		ExecutablePath: chromePathPtr,
		Headless: &headless,
		SlowMo: &time_slow,
	})
	if err != nil {
		t.Fatalf("could not launch browser: %v", err)
	}
	defer browser.Close()

	page, _ := browser.NewPage()
	defer page.Close()
	//page.Goto("https://www.saucedemo.com/")
	loginPage := pages.NewLoginPage(page)
	err = loginPage.NavigateTo("https://www.saucedemo.com/")
	if err != nil {
			t.Fatalf("Failed to navigate to login page: %v", err)
	}

	loginPage.EnterUsername("standard_user")
	loginPage.EnterPassword("secret_sauce")
	loginPage.ClickLoginButton()

	inventoryPage := pages.NewInventoryPage(page)
	isVisible := inventoryPage.Verify()
	if !(isVisible)  {
		t.Fatalf("User login failed")
	}
}

func TestLoginFail(t *testing.T) {
	pw, err := playwright.Run()
	if err != nil {
		t.Fatalf("could not start playwright: %v", err)
	}
	
	chromePath := "C:/Program Files/Google/Chrome/Application/chrome.exe"
	chromePathPtr := &chromePath
	headless := false
	time_slow := 1000.0
	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		ExecutablePath: chromePathPtr,
		Headless: &headless,
		SlowMo: &time_slow,
	})
	if err != nil {
		t.Fatalf("could not launch browser: %v", err)
	}
	defer browser.Close()

	page, _ := browser.NewPage()
	defer page.Close()
	//page.Goto("https://www.saucedemo.com/")
	loginPage := pages.NewLoginPage(page)
	err = loginPage.NavigateTo("https://www.saucedemo.com/")
	if err != nil {
			t.Fatalf("Failed to navigate to login page: %v", err)
	}

	loginPage.EnterUsername("standard_user")
	loginPage.EnterPassword("bad_password")
	loginPage.ClickLoginButton()

	err = loginPage.VerifyLoginError()

	if err != nil {
		t.Fatalf("Failed to find login error: %v", err)
	}
}