package pages

import (
        "fmt"
        "github.com/playwright-community/playwright-go"
)

type LoginPage struct {
        page playwright.Page
}

func NewLoginPage(page playwright.Page) *LoginPage {
        return &LoginPage{
                page: page,
        }
}

func (l LoginPage) NavigateTo(url string) error {
        _, err := l.page.Goto(url) 
        if err != nil {
            return fmt.Errorf("Failed to navigate to URL: %w", err) 
        }
        return nil
}

func (l LoginPage) EnterUsername(username string) error {
        err := l.page.Fill("#user-name", username) // Recommended approach
        if err != nil {
            // Handle the error
            return fmt.Errorf("Failed to enter username: %w", err)
        }
        return nil
}

func (l LoginPage) EnterPassword(password string) error {
        err := l.page.Fill("#password", password)
        if err != nil {
                return fmt.Errorf("Failed to fill password: %w", err) 
        }
        return nil
}

func (l LoginPage) ClickLoginButton() error {
        err := l.page.Click("#login-button")
        if err != nil {
                return fmt.Errorf("Failed to click login: %w", err) 
        }
        return nil    
}

func (l LoginPage) VerifyLoginError() error {
        
        isVisible,_ := l.page.Locator(`//h3[@data-test='error' 
        and normalize-space(.)='Epic sadface: Username and password do not match any user in this service']`).IsVisible()

	if !(isVisible) {
                return fmt.Errorf("Login Failure error message not visible")
        }
        return nil


}