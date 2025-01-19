package pages

import (
	"github.com/playwright-community/playwright-go"
)

type InventoryPage struct {
	page playwright.Page
}

func NewInventoryPage(page playwright.Page) *InventoryPage {
	return &InventoryPage{
			page: page,
	}
}

func (l InventoryPage) Verify() bool {
	isVisible,_ := l.page.Locator("//span[.='Products']").IsVisible()
	return isVisible
}