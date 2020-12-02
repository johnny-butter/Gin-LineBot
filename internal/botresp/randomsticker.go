package botresp

import (
	"math/rand"
	"strconv"

	"github.com/line/line-bot-sdk-go/linebot"
)

type RandomStickerClass struct{}

func NewRandomSticker() *RandomStickerClass {
	return new(RandomStickerClass)
}

func (r *RandomStickerClass) Response(msg string) linebot.SendingMessage {
	stickerID := RandomIntBetween(52002734, 52002773)

	return linebot.NewStickerMessage("11537", strconv.FormatInt(int64(stickerID), 10))
}

func RandomIntBetween(min, max int) int {
	return rand.Intn(max-min) + min
}
