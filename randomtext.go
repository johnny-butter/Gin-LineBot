package main

import "github.com/line/line-bot-sdk-go/linebot"

type RandomTextClass struct{}

func NewRandomText() *RandomTextClass {
	return new(RandomTextClass)
}

func (r *RandomTextClass) Response(msg string) linebot.SendingMessage {
	return linebot.NewTextMessage(msg)
}
