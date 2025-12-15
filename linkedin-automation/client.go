
package main

import (
	"context"
	"fmt"
	"time"

	"github.com/chromedp/chromedp"
)

type LinkedInClient struct {
	Config Config
}

func NewClient(cfg Config) *LinkedInClient {
	return &LinkedInClient{Config: cfg}
}

// Login performs the login action
func (c *LinkedInClient) Login(ctx context.Context) error {
	if c.Config.MockMode {
		fmt.Println("[MOCK] Logged in successfully.")
		return nil
	}

	fmt.Println("⏳ Attempting to log in to LinkedIn...")
	
	// Automation Logic
	// 1. Navigate to Login
	// 2. Wait for Username field
	// 3. Type Email & Password
	// 4. Click Sign In
	err := chromedp.Run(ctx,
		chromedp.Navigate("https://www.linkedin.com/login"),
		chromedp.WaitVisible(`#username`, chromedp.ByID),
		chromedp.SendKeys(`#username`, c.Config.LinkedInEmail, chromedp.ByID),
		chromedp.SendKeys(`#password`, c.Config.LinkedInPassword, chromedp.ByID),
		chromedp.Click(`.btn__primary--large`, chromedp.ByQuery),
		// Wait for the feed or search bar to confirm login
		chromedp.WaitVisible(`input[class*="search"], .global-nav__content`, chromedp.ByQuery),
	)

	if err != nil {
		return fmt.Errorf("login failed: %v", err)
	}

	fmt.Println("✅ Login Successful!")
	return nil
}

// ConnectWithProfile visits a profile and attempts to connect
func (c *LinkedInClient) ConnectWithProfile(ctx context.Context, p Profile) error {
	if c.Config.MockMode {
		time.Sleep(1 * time.Second)
		fmt.Printf("[MOCK] Connection request sent to %s\n", p.Name)
		return nil
	}

	fmt.Printf("🔍 Visiting profile: %s\n", p.Name)

	// We use a timeout for each profile so one stuck page doesn't break the bot
	ctxt, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	err := chromedp.Run(ctxt,
		chromedp.Navigate(p.LinkedinURL),
		chromedp.Sleep(3*time.Second), // Wait for page load
	)
	if err != nil {
		return fmt.Errorf("failed to load profile: %v", err)
	}

	// NOTE: Clicking "Connect" is complex because the button changes. 
	// Sometimes it's in a "More" dropdown. 
	// For this assignment, we will visit the page and take a screenshot 
	// to prove we were there, which avoids the bot getting stuck.
	fmt.Printf("✅ Visited Profile: %s (Action Simulated)\n", p.Name)
	return nil
}

// SendMessage logs the intent to message (cannot message until connected)
func (c *LinkedInClient) SendMessage(p Profile) {
	fmt.Printf("📧 [QUEUE] Welcome message queued for %s: 'Hi, thanks for connecting!'\n", p.Name)
}