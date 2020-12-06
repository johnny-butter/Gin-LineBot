package main

import (
	"fmt"
	"log"
	"os"

	"internal/botresp"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	router := gin.Default()

	db := GetDBConnect(os.Getenv("ENV"))

	bot, err := linebot.New(
		os.Getenv("LINE_CHANNEL_SECRET"),
		os.Getenv("LINE_CHANNEL_ACCESS_TOKEN"),
	)

	if err != nil {
		log.Fatal(err)
	}

	router.POST("/callback", func(c *gin.Context) {
		fmt.Println(c.Request)
		events, err := bot.ParseRequest(c.Request)
		if err != nil {
			if err == linebot.ErrInvalidSignature {
				c.Writer.WriteHeader(400)

				c.JSON(http.StatusBadRequest, struct {
					Message string `json:"message"`
				}{"Invalid signature error"})
				return
			} else {
				c.Writer.WriteHeader(500)

				c.JSON(http.StatusInternalServerError, struct {
					Message string `json:"message"`
				}{"Events parse error"})
				return
			}
		}
		for _, event := range events {
			if event.Type == linebot.EventTypeMessage {
				var err error
				var r botresp.BotResponse

				switch message := event.Message.(type) {
				case *linebot.TextMessage:
					keywords := GoKeyword{}
					db.Where("keyword = ?", message.Text).First(&keywords)

					r, err = botresp.NewBotResponse(keywords.ResponseCls)
					if err != nil {
						c.JSON(http.StatusBadRequest, struct {
							Message string `json:"message"`
						}{Message: err.Error()})
						return
					}

					if _, err = bot.ReplyMessage(event.ReplyToken, r.Response(message.Text)).Do(); err != nil {
						log.Print(err)
					}

				case *linebot.StickerMessage:
					r, err = botresp.NewBotResponse("RandomSticker")
					if err != nil {
						c.JSON(http.StatusBadRequest, struct {
							Message string `json:"message"`
						}{Message: err.Error()})
						return
					}

					if _, err = bot.ReplyMessage(event.ReplyToken, r.Response("")).Do(); err != nil {
						log.Print(err)
					}
				}
			}
		}

		c.String(http.StatusOK, "ok")
	})

	if p, ok := os.LookupEnv("PORT"); ok {
		router.Run(":" + p)
	} else {
		router.Run(":8080")
	}
}
