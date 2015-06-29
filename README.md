# Overview
[![Twitter](https://img.shields.io/badge/author-%40MachielMolenaar-blue.svg)](https://twitter.com/MachielMolenaar)
[![GoDoc](https://godoc.org/github.com/Machiel/telegrambot?status.svg)](https://godoc.org/github.com/Machiel/telegrambot)

TelegramBot helps you to easily access the Telegram Bot API

# License
Slugify is licensed under a MIT license.

# Warning
This package has no tests yet, so use at your own risk. This was quickly hacked
together!

# Installation
A simple `go get github.com/Machiel/telegrambot` should suffice.

# Usage

## Example

```go
package main

import (
	"math/rand"
	"time"

	"github.com/Machiel/telegrambot"
)

var answers = []string{
	"Hodor!",
	"Hodor, hodor... hodor!",
	"Hodor?",
	"Hodor...",
	"Hodor...! Hodor...!",
	"Hooodorrrr!!!",
	"Hodor",
	"Hodor?!",
}

func main() {

	telegram := telegrambot.New("api-key")

	offset := new(int64)
	for {

		updates, _ := telegram.GetUpdates(*offset)

		for _, update := range updates.Result {

			if update.ID >= *offset {
				*offset = (update.ID + 1)
			}

			message := update.Message
			answer := answers[rand.Intn(len(answers))]

			telegram.SendMessage(message.Chat.ID, answer, telegrambot.SendMessageOptions{})
		}

		time.Sleep(500)
	}

}
```
