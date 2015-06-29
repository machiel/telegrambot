package telegrambot

import "net/url"

// ForwardMessage allows you to forward a message from one conversation to another
func (b TelegramBot) ForwardMessage(chatID int64, fromChatID int64, messageID int64) (string, error) {
	v := url.Values{}
	v.Add("chat_id", asString(chatID))
	v.Add("from_chat_id", asString(fromChatID))
	v.Add("message_id", asString(messageID))

	var err error
	var resp []byte

	if resp, err = b.executeRequest("/forwardMessage", v); err != nil {
		return "", err
	}

	return string(resp), nil
}
