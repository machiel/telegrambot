package telegrambot

import (
	"encoding/json"
	"net/url"
)

// GetUpdates retrieves the latest updates from the Telegram servers
func (b TelegramBot) GetUpdates(offset int64) (Update, error) {

	v := url.Values{}
	v.Add("offset", asString(offset))

	var u Update
	resp, err := b.executeRequest("/getUpdates", v)

	if err != nil {
		return u, err
	}

	err = json.Unmarshal(resp, &u)

	if err != nil {
		return u, err
	}

	return u, nil
}
