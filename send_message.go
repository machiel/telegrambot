package telegrambot

import (
	"encoding/json"
	"log"
	"net/url"
	"strconv"
)

// Keyboard allows you to specify a specific keyboard layout to respond with
type Keyboard struct {
	Layout     [][]string `json:"keyboard"`
	Consumable bool       `json:"one_time_keyboard"`
}

// SendMessageOptions allows you to define optional parameters
type SendMessageOptions struct {
	ReplyMarkup *Keyboard
}

// SendMessage allows you to send a message to a specific target
func (b TelegramBot) SendMessage(id int64, text string, options SendMessageOptions) (string, error) {
	v := url.Values{}
	v.Add("chat_id", asString(id))
	v.Add("text", text)

	if options.ReplyMarkup != nil {
		json, err := json.Marshal(options.ReplyMarkup)

		if err != nil {
			log.Printf("Could not marshall keyboard: %s", err)
		} else {
			v.Add("reply_markup", string(json))
		}
	}

	resp, err := b.executeRequest("/sendMessage", v)

	if err != nil {
		return "", err
	}

	return string(resp), nil

}

// SendPhoto allows you to send a photo to a specific target
func (b TelegramBot) SendPhoto(id int64, fileID string, options SendMessageOptions) (string, error) {
	return b.sendData(id, fileID, options, "photo", "/sendPhoto")
}

// SendVoice allows you to send a voicemessage to a specific target
func (b TelegramBot) SendVoice(id int64, fileID string, options SendMessageOptions) (string, error) {
	return b.sendData(id, fileID, options, "voice", "/sendVoice")
}


// SendSticker allows you to send a sticker to a specific target
func (b TelegramBot) SendSticker(id int64, fileID string, options SendMessageOptions) (string, error) {
	return b.sendData(id, fileID, options, "sticker", "/sendSticker")
}

// SendAudio allows you to send audio to a specific target
func (b TelegramBot) SendAudio(id int64, fileID string, options SendMessageOptions) (string, error) {
	return b.sendData(id, fileID, options, "audio", "/sendAudio")
}

// SendVideo allows you to send a video to a specific target
func (b TelegramBot) SendVideo(id int64, fileID string, options SendMessageOptions) (string, error) {
	return b.sendData(id, fileID, options, "video", "/sendVideo")
}

// SendDocument allows you to send a document to a specific target
func (b TelegramBot) SendDocument(id int64, fileID string, options SendMessageOptions) (string, error) {
	return b.sendData(id, fileID, options, "document", "/sendDocument")
}

// SendLocation allows you to send a location to a chat
func (b TelegramBot) SendLocation(id int64, latitude float64, longitude float64, options SendMessageOptions) (string, error) {

	v := url.Values{}
	v.Add("chat_id", asString(id))
	v.Add("longitude", strconv.FormatFloat(longitude, 'f', -1, 64))
	v.Add("latitude", strconv.FormatFloat(latitude, 'f', -1, 64))
	return b.finalizeRequest("/sendLocation", v)
}

func (b TelegramBot) sendData(id int64, data string, options SendMessageOptions, dataKey string, target string) (string, error) {
	v := url.Values{}
	v.Add("chat_id", asString(id))
	v.Add(dataKey, data)

	return b.finalizeRequest(target, v)
}

func (b TelegramBot) finalizeRequest(target string, v url.Values) (string, error) {
	resp, err := b.executeRequest(target, v)

	if err != nil {
		return "", err
	}

	return string(resp), nil

}
