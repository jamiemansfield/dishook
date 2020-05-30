package dishook

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// Sends the webhook to Discord. *waves*
func SendWebhook(client *http.Client, url string, hook *Webhook) error {
	if client == nil {
		client = http.DefaultClient
	}

	// Encode body as JSON
	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false)
	if err := enc.Encode(hook); err != nil {
		return err
	}

	// Create the request
	req, err := http.NewRequest(http.MethodPost, url, buf)
	if err != nil {
		return err
	}
	req.Header.Set("User-Agent", "Dishook/0.1.0")
	req.Header.Set("Content-Type", "application/json")

	// Do the request
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// TODO: see if there any issues

	return nil
}
