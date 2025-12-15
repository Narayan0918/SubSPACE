# 🤖 LinkedIn Automation Bot (Golang)

A high-performance, concurrent automation tool designed to manage LinkedIn connection requests and messaging at scale.

Built using **Golang** to leverage its native concurrency primitives (Goroutines and Channels), this project demonstrates a robust **Worker Pool architecture** capable of handling multiple profiles simultaneously while respecting rate limits.

---

## 🏗️ Architecture & Design

The core of this application is the **Worker Pool Pattern**. Instead of processing profiles sequentially (one by one), the system spawns a configurable number of "Workers" that pull jobs from a queue. This ensures:
1.  **Scalability:** You can increase throughput simply by adding more workers.
2.  **Resource Management:** Prevents the system from crashing by limiting active browser instances.
3.  **Efficiency:** I/O operations (network requests) are handled concurrently.



### Key Components
* **The Orchestrator (`main.go`):** Initializes the browser context and distributes work.
* **The Worker:** A goroutine that picks a `Profile` from the `jobs` channel, processes it, and sends the result to the `results` channel.
* **The Client (`client.go`):** A modular interface for LinkedIn interactions. It features a **Hybrid Design**:
    * **Mock Mode (Default):** Simulates network delays and API responses for safe development and testing.
    * **Live Mode:** Uses `chromedp` (Headless Chrome) to perform actual browser automation.

---

## 🚀 Features

* **⚡ Concurrent Execution:** Processes multiple leads in parallel using Go routines.
* **🛡️ Safety-First Design:** Defaults to "Mock Mode" to prevent accidental IP bans or account flagging during development.
* **🌐 Real Browser Automation:** Integrated with [Chromedp](https://github.com/chromedp/chromedp) to drive a real Chrome instance for login and navigation.
* **📄 Dynamic Input:** Parses target profiles dynamically from a `leads.csv` file.
* **📊 Execution Reporting:** Provides a summary of successful and failed operations at the end of the run.

---

## 🛠️ Prerequisites

* **Golang:** v1.20 or higher ([Download Here](https://go.dev/dl/))
* **Google Chrome:** Installed on the host machine (for Live Mode).
* **Git:** ([Download Here](https://git-scm.com/downloads))

---

## 📥 Installation & Setup

1.  **Clone the Repository**
    ```bash
    git clone [https://github.com/YOUR_USERNAME/linkedin-automation-assignment.git](https://github.com/YOUR_USERNAME/linkedin-automation-assignment.git)
    cd linkedin-automation-assignment
    ```

2.  **Install Dependencies**
    This project uses `chromedp` for browser control.
    ```bash
    go mod tidy
    ```

3.  **Prepare Input Data**
    Ensure `leads.csv` is present in the root directory. Format:
    ```csv
    ID,Name,LinkedinURL
    1,Satya Nadella,[https://www.linkedin.com/in/satya-nadella/](https://www.linkedin.com/in/satya-nadella/)
    2,Bill Gates,[https://www.linkedin.com/in/williamhgates/](https://www.linkedin.com/in/williamhgates/)
    ```

---

## ⚙️ Configuration

You can configure the bot's behavior by modifying the `Config` struct in `main.go`:

```go
cfg := Config{
    LinkedInEmail:    "your-email@example.com",
    LinkedInPassword: "your-password",
    HeadlessMode:     false, // Set 'true' to hide the browser window
    MockMode:         true,  // Set 'false' to perform REAL actions
}

---

🏃 Usage
You can run the automation tool directly from your terminal. By default, the project comes with a leads.csv file for testing.

1. Standard Run (Simulation Mode) This mode validates your logic without launching a browser, making it safe and fast for testing.

Bash

go run .
2. Live Browser Mode To see the bot control the Chrome browser in real-time:

1. Open main.go.

2. Change the configuration struct:

MockMode: false,     // Enables real browser actions
HeadlessMode: false, // Shows the browser UI
Run the command again: go run .

---
💻 Sample Output
The output demonstrates the Worker Pool in action. Notice how multiple profiles are processed simultaneously (concurrently) rather than one by one.

🚀 Starting LinkedIn Automation Bot...
⏳ Attempting to log in to LinkedIn...
[MOCK] Logged in successfully.

Worker 1 started.
Worker 2 started.
Worker 3 started.

🔍 Visiting profile: Satya Nadella
🔍 Visiting profile: Bill Gates
[MOCK] Connection request sent to Satya Nadella
📧 [QUEUE] Welcome message queued for Satya Nadella: 'Hi, thanks for connecting!'
[MOCK] Connection request sent to Bill Gates
📧 [QUEUE] Welcome message queued for Bill Gates: 'Hi, thanks for connecting!'

--- 📊 Final Execution Report ---
[Success] Satya Nadella
[Success] Bill Gates

---
📂 Project Structure
This project follows a clean, modular architecture to separate data handling, business logic, and orchestration.

linkedin-automation/
├── main.go           # The Manager: Initializes the browser and orchestrates the Worker Pool.
├── client.go         # The Worker: Contains the 'chromedp' logic for controlling the browser.
├── csv_reader.go     # Data Layer: Handles parsing and validation of the input CSV file.
├── types.go          # Blueprints: Defines the data structures (Profile, Config) used across the app.
├── leads.csv         # Input: A list of target profiles (ID, Name, URL).
└── go.mod            # Dependencies: Manages external libraries like 'chromedp'.

---
⚠️ Disclaimer
Please Read Carefully: This software is developed strictly for educational and assessment purposes as part of an internship application.

Terms of Service: Automated interaction with LinkedIn (scraping, botting) may violate their User Agreement.

Liability: The author is not responsible for any account restrictions, bans, or legal consequences resulting from the use of this tool on live accounts.

Recommendation: Use the provided MockMode for all demonstrations to ensure compliance and safety.

---
🔗 References & Resources
Browser Automation: Chromedp GitHub Repository - The engine used to drive the Chrome browser.

Concurrency in Go: Go by Example: Worker Pools - The pattern used to manage scalable requests.

CSV Handling: Go Standard Library: encoding/csv - Documentation for the file parser.