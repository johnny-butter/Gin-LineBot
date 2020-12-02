package main

import (
	"errors"

	"github.com/line/line-bot-sdk-go/linebot"
)

type BotResponse interface {
	Response(string) linebot.SendingMessage
}

const (
	RandomText    string = "RandomText"
	RandomSticker string = "RandomSticker"
)

func NewBotResponse(t string) (BotResponse, error) {
	switch t {
	case RandomText:
		return NewRandomText(), nil
	case RandomSticker:
		return NewRandomSticker(), nil
	default:
		return nil, errors.New("unknown response type")
	}
}
