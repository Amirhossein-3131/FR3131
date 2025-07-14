# FR3131

**FR3131** is a Go-based subdomain reconnaissance tool that wraps multiple enumeration tools and automates the process of discovering, cleaning, and combining subdomains.

---

## 🚀 Features

- Runs multiple tools in parallel:
  - `assetfinder`
  - `subfinder`
  - `Sublist3r`
  - `findomain`
- Combines and deduplicates all results
- Filters color codes and unwanted formatting
- Saves final output to `all-sub-finder.txt`
- (Optional) Sends real-time status updates via Telegram

---

## 🛠 Requirements

- **Go** (v1.20+ recommended)
- **Python3**
- Tools must be installed and accessible from system `$PATH`:
  - [`assetfinder`](https://github.com/tomnomnom/assetfinder)
  - [`subfinder`](https://github.com/projectdiscovery/subfinder)
  - [`findomain`](https://github.com/findomain/findomain)
- **Sublist3r** should be cloned at: `/opt/Sublist3r`
- *(Optional)* Telegram bot configuration in `summary/alert.go` (see below)

---

## 📦 Installation

Clone the repository and tidy Go dependencies:

```bash
git clone https://github.com/َAmirhossein-3131/FR3131.git
cd FR3131
go mod tidy
````

---

## ⚙️ Usage

Run the tool by specifying a target domain:

```bash
go run main.go -t example.com
```

---

## 🧠 What It Does

1. Runs all four tools to find subdomains
2. Logs how many domains each tool found
3. Cleans and filters output (removes ANSI escape sequences, etc.)
4. Combines all results and removes duplicates
5. Saves final list to: `all-sub-finder.txt`
6. *(Optional)* Sends progress updates to Telegram (if enabled)

---

## 📁 Output Files

| File name                | Description                  |
| ------------------------ | ---------------------------- |
| `assetfinder-output.txt` | Raw output from Assetfinder  |
| `subfinder-output.txt`   | Raw output from Subfinder    |
| `findomain-output.txt`   | Raw output from Findomain    |
| `sublist3r-output.txt`   | Raw output from Sublist3r    |
| `all-sub-finder.txt`     | Final cleaned & deduplicated |

---

## 🔔 Telegram Integration (Optional)

To enable Telegram alerts:

1. Open `summary/alert.go`
2. Set your bot token and chat ID:

```go
const (
  botToken = "your_bot_token"
  chatID   = "your_chat_id"
)
```

> If you leave both values empty (`""`), Telegram alerts will be disabled automatically.

To get your chat ID, use:

```
https://api.telegram.org/bot<your_bot_token>/getUpdates
```

---

 **FR3131** – Fast and simple multi-tool subdomain reconnaissance in Go
