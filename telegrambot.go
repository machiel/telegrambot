package telegrambot

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

const endpoint = "https://api.telegram.org/bot"

type TelegramBot struct {
	APIKey string
}

func New(key string) TelegramBot {
	return TelegramBot{APIKey: key}
}

func (b TelegramBot) executeRequest(target string, v url.Values) ([]byte, error) {
	resp, err := http.Get(b.getURL(target, v))

	if err != nil {
		return []byte{}, err
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return body, err
	}

	return body, nil
}

func (b TelegramBot) getURL(target string, v url.Values) string {
	return endpoint + b.APIKey + target + "?" + v.Encode()
}

func asString(i int64) string {
	return strconv.FormatInt(i, 10)
}
