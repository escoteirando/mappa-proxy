package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/gofiber/fiber/v2"
	// _ "github.com/joho/godotenv/autoload"
)

func init() {
	Routes["/tg/pub"] = RouteData{
		Name:      "TelegramPublisher",
		Method:    "POST",
		Handler:   TelegramPublisherHandler,
		CacheTime: 0,
		Mappa:     false,
	}
}

type BotRequestData struct {
	// Chat ID
	CId int64 `json:"cId"`
	// Message ID to respond to
	MId int64 `json:"mId"`
	// Message
	Msg string `json:"msg"`
}

const exception string = "Bot publisher error"

// TelegramPublisherHandler godoc
//	@Summary		Telegram Publisher handler
//	@Description	Publica mensagem em chat do Telegram
//	@Tags			telegram
//	@Accept			json
//	@Param			request	body	BotRequestData	true	"Bot request data"
//	@Produce		json
//	@Success		200	{object}	handlers.ReplyMessage
//	@Failure		400	{object}	handlers.ReplyMessage
//	@Router			/tg/pub [post]
func TelegramPublisherHandler(c *fiber.Ctx) error {
	botToken := os.Getenv("BOT_TOKEN")
	if len(botToken) == 0 {
		log.Println("Missing BOT_TOKEN environment variable")
		return reply_error(c, 400, "MISSING BOT TOKEN", nil)
	}

	var botData BotRequestData

	if err := json.Unmarshal(c.Body(), &botData); err != nil {
		log.Printf("Failed to serialize login request %s\n", err)
		return reply_error(c, 400, "Failed to serialize login request", err)
	}
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		return reply_error(c, 400, "Failed to instance BOT API", err)
	}

	msg := tgbotapi.NewMessage(botData.CId, botData.Msg)
	//msg.ReplyToMessageID = int(botData.MId)
	nmsg, err := bot.Send(msg)
	if err != nil {
		log.Printf("Failed to send message: %v\n", err)
		return reply_error(c, 400, "Failed to send message", err)
	}
	return reply_status(c, 200, fmt.Sprintf("Message sent: %v", nmsg))
}
