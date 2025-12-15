
package main

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/chromedp/chromedp"
)

// worker processes profiles from the jobs channel
func worker(id int, client *LinkedInClient, ctx context.Context, jobs <-chan Profile, results chan<- Profile, wg *sync.WaitGroup) {
	defer wg.Done()
	for profile := range jobs {
		// 1. Visit and Connect
		err := client.ConnectWithProfile(ctx, profile)
		if err != nil {
			fmt.Printf("❌ Worker %d Failed on %s: %v\n", id, profile.Name, err)
			profile.Status = "Failed"
			results <- profile
			continue
		}

		// 2. Queue Message (Simulated)
		client.SendMessage(profile)
		profile.Status = "Success"
		results <- profile
	}
}

func main() {
	// --- CONFIGURATION ---
	// REPLACE THESE WITH A DUMMY ACCOUNT! TO CHECK ACTUAL WORKING
	cfg := Config{
		LinkedInEmail:    "your-dummy-email@gmail.com",
		LinkedInPassword: "your-dummy-password",
		HeadlessMode:     false, // False = Show Browser UI
		MockMode:         true,  // SET TO TRUE FOR SAFETY/TESTING. Set FALSE to really run it.
	}
	// ---------------------

	fmt.Println("🚀 Starting LinkedIn Automation Bot...")

	// 1. Setup Chrome Context
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", cfg.HeadlessMode),
		chromedp.Flag("disable-gpu", true),
		chromedp.WindowSize(1200, 800),
	)

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	// Create a browser context
	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	// 2. Initialize Client
	client := NewClient(cfg)

	// 3. Login (Only if not in Mock Mode)
	if !cfg.MockMode {
		if err := client.Login(ctx); err != nil {
			log.Fatalf("Critical Error: Could not log in. %v", err)
		}
	}

	// 4. Read Input File
	profiles, err := ReadLeads("leads.csv")
	if err != nil {
		log.Fatalf("Error reading CSV: %v", err)
	}

	// 5. Start Worker Pool
	// Note: With a real browser, use 1 worker to avoid account flags.
	// In MockMode, you can use more.
	numWorkers := 1
	if cfg.MockMode {
		numWorkers = 3
	}

	jobs := make(chan Profile, len(profiles))
	results := make(chan Profile, len(profiles))
	var wg sync.WaitGroup

	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		// We pass the SAME browser context to all workers
		go worker(w, client, ctx, jobs, results, &wg)
	}

	// 6. Send Jobs
	for _, p := range profiles {
		jobs <- p
	}
	close(jobs)

	// 7. Wait and Close
	wg.Wait()
	close(results)

	// 8. Final Report
	fmt.Println("\n--- 📊 Final Execution Report ---")
	for p := range results {
		fmt.Printf("[%s] %s\n", p.Status, p.Name)
	}
}