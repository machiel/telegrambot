package telegrambot

type Update struct {
	Ok     bool         `json:"ok"`
	Result []UpdateItem `json:"result"`
}

type UpdateItem struct {
	Message Message `json:"message"`
	ID      int64   `json:"update_id"`
}

type Audio struct {
	Duration int    `json:"duration"`
	FileID   string `json:"file_id"`
	FileSize int    `json:"file_size"`
	MimeType string `json:"mime_type"`
}

type Contact struct {
	FirstName   string `json:"first_name"`
	PhoneNumber string `json:"phone_number"`
}

type Message struct {
	Chat     Chat        `json:"chat"`
	Date     int64       `json:"date"`
	From     User        `json:"from"`
	ID       int64       `json:"message_id"`
	Text     string      `json:"text"`
	Sticker  Sticker     `json:"sticker"`
	Location Location    `json:"location"`
	Document Document    `json:"document"`
	Photo    []PhotoSize `json:"photo"`
	Video    Video       `json:"video"`
	Contact  Contact     `json:"contact"`
	Audio    Audio       `json:"audio"`
}

type Document struct {
	FileID   string    `json:"file_id"`
	FileName string    `json:"file_name"`
	FileSize int       `json:"file_size"`
	MimeType string    `json:"mime_type"`
	Thumb    PhotoSize `json:"thumb"`
}

type PhotoSize struct {
	FileID   string `json:"file_id"`
	FileSize int    `json:"file_size"`
	Height   int    `json:"height"`
	Width    int    `json:"width"`
}

type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Sticker struct {
	FileID   string    `json:"file_id"`
	FileSize int       `json:"file_size"`
	Height   int       `json:"height"`
	Thumb    PhotoSize `json:"thumb"`
	Width    int       `json:"width"`
}

type Video struct {
	Caption  string    `json:"caption"`
	Duration int       `json:"duration"`
	FileID   string    `json:"file_id"`
	FileSize int       `json:"file_size"`
	Height   int       `json:"height"`
	Thumb    PhotoSize `json:"thumb"`
	Width    int       `json:"width"`
}

type User struct {
	FirstName string `json:"first_name"`
	ID        int64  `json:"id"`
	Username  string `json:"username"`
}

type Chat struct {
	FirstName string `json:"first_name"`
	ID        int64  `json:"id"`
	Username  string `json:"username"`
}
