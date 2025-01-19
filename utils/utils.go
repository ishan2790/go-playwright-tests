package utils

import (
        "fmt"
        "io/ioutil"
        "os"
        "path/filepath"
        "gopkg.in/yaml.v3"
        "github.com/playwright-community/playwright-go"
)

// BrowserOptions holds options to configure the browser launch
type BrowserOptions struct {
        ExecutablePath string `yaml:"executable_path"`
        Headless       bool   `yaml:"headless"`
        SlowMo         float64 `yaml:"slow_mo"`
		AppUrl         string `yaml:"app_url"`
}

// ReadBrowserOptionsFromYaml reads browser options from a YAML file
func ReadBrowserOptionsFromYaml(filePath string) (BrowserOptions, error) {
        var options BrowserOptions

        // Get the absolute path of the config file
        absFilePath, err := filepath.Abs(filePath)
        if err != nil {
                return options, fmt.Errorf("failed to get absolute path of config file: %w", err)
        }

        // Read the YAML file
        data, err := ioutil.ReadFile(absFilePath)
        if err != nil {
                return options, fmt.Errorf("failed to read YAML file: %w", err)
        }

        // Unmarshal YAML data into BrowserOptions struct
        err = yaml.Unmarshal(data, &options)
        if err != nil {
                return options, fmt.Errorf("failed to unmarshal YAML data: %w", err)
        }

        return options, nil
}

// LaunchPlaywrightBrowser launches a Chromium browser with the provided options
func LaunchPlaywrightBrowserAndOpenLoginPage() (playwright.Page, error) {
        // Get the absolute path of the config file
        _, filename := filepath.Split(os.Args[0]) // Use filepath.Split for this purpose
        configFilePath := filepath.Join(filepath.Dir(filename), "browser_config.yaml") 

        // Read browser options from YAML
        options, err := ReadBrowserOptionsFromYaml(configFilePath)
        if err != nil {
                return nil, fmt.Errorf("failed to read browser options: %w", err)
        }

        pw, err := playwright.Run()
        if err != nil {
                return nil, err
        }

        browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
                ExecutablePath: &options.ExecutablePath,
                Headless:        &options.Headless,
                SlowMo:          &options.SlowMo,
        })

        if err != nil {
                return nil, err
        }



		page, _ := browser.NewPage()
		

		if _, err := page.Goto(options.AppUrl); err != nil {
			return nil, fmt.Errorf("Failed to navigate to login page: %w", err)
		}

		return page, nil
}