# FR3131

**FR3131** is a Go-based subdomain reconnaissance tool that wraps multiple enumeration tools and automates the process of discovering and combining subdomains.

## 🚀 Features

- Runs multiple tools in parallel:
  - `assetfinder`
  - `subfinder`
  - `Sublist3r`
  - `findomain`
- Combines and deduplicates all results
- Filters unwanted characters and formatting
- Sends real-time status notifications via Telegram

## 🛠 Requirements

- Go (v1.20+ recommended)
- Python3
- Tools installed in system `$PATH`:
  - `assetfinder`
  - `subfinder`
  - `findomain`
- Sublist3r cloned at: `/opt/Sublist3r`
- Telegram bot setup in `summary/alert.go` (send messages to your own Telegram chat)

## 📦 Installation

Clone the repository:

```bash
git clone https://github.com/your-username/FR3131.git
cd FR3131
go mod tidy

⚙️ Usage

go run main.go -t example.com
The tool will:

Run all recon tools

Log how many domains each tool found

Clean and combine results

Save final output in all-sub-finder.txt
