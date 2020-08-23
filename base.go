package main

import "github.com/line/line-bot-sdk-go/linebot"

type BotResponse interface {
	Response(string) linebot.SendingMessage
}

const (
	RandomText    string = "RandomText"
	RandomSticker string = "RandomSticker"
)

func NewBotResponse(t string) BotResponse {
	switch t {
	case RandomText:
		return NewRandomText()
	case RandomSticker:
		return NewRandomSticker()
	default:
		panic("Unknown response type")
	}
}
