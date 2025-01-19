package main

import (
  "testing"

  "github.com/ishan2790/go-playwright-tests/pages"
  "github.com/ishan2790/go-playwright-tests/utils"
)

func TestLoginPass(t *testing.T) {

  page, err := utils.LaunchPlaywrightBrowserAndOpenLoginPage()
  defer page.Close()

  if err != nil {
    t.Fatalf("could not navigate to url: %v", err)
  }


  loginPage := pages.NewLoginPage(page)

  loginPage.EnterUsername("standard_user")
  loginPage.EnterPassword("secret_sauce")
  loginPage.ClickLoginButton()

  inventoryPage := pages.NewInventoryPage(page)
  isVisible := inventoryPage.Verify()
  if !isVisible {
    t.Fatalf("User login failed")
  }
  
}

func TestLoginFail(t *testing.T) {
	page, err := utils.LaunchPlaywrightBrowserAndOpenLoginPage()
	defer page.Close()
  
	if err != nil {
	  t.Fatalf("could not navigate to url: %v", err)
	}
  
  
	loginPage := pages.NewLoginPage(page)
  
	loginPage.EnterUsername("standard_user")
	loginPage.EnterPassword("bad_password")
	loginPage.ClickLoginButton()

   err = loginPage.VerifyLoginError()

   if err != nil {
     t.Fatalf("Failed to find login error: %v", err)
   }
}