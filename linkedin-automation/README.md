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
