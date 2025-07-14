package alert

import (
	"bytes"
	"fmt"
	"net/http"
	"time"
)

// Leave botToken and chatID empty to disable Telegram alerts
const (
	botToken = "" // Set your bot token here, or leave blank to disable
	chatID   = "" // Set your chat ID here, or leave blank to disable
)

// SendTelegramMessage sends a message to a Telegram chat via bot API.
// If botToken or chatID is empty, it does nothing and returns silently.
func SendTelegramMessage(message string) error {
	if botToken == "" || chatID == "" {
		// Telegram alerts are disabled
		return nil
	}

	currentTime := time.Now()
	timeStamp := currentTime.Format("2006-01-02 15:04:05")
	fullMessage := fmt.Sprintf("%s\n%s", message, timeStamp)

	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", botToken)
	data := []byte(fmt.Sprintf("chat_id=%s&text=%s", chatID, fullMessage))

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		return fmt.Errorf("Error creating request: %s", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("Error sending request: %s", err)
	}
	defer resp.Body.Close()

	return nil
}
