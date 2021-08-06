package tg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type BotRequestData struct {
	CId int64  `json:"cId"`
	MId int64  `json:"mId"`
	Msg string `json:"msg"`
}

const exception string = "Bot publisher error"

func Publish(c *gin.Context) {
	botToken := "1906817161:AAGXe-HSMfvmBOOUoqYEZiKZy53KoBARJE8" //os.Getenv("BOT_TOKEN") ||
	if len(botToken) == 0 {
		log.Println("Missing BOT_TOKEN environment variable")
		c.JSON(400, gin.H{"message": exception, "error": fmt.Errorf("Missing BOT_TOKEN")})
		return
	}
	requestBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Printf("Failed to get request body: %s\n", err)
		c.JSON(400, gin.H{"message": exception, "error": err.Error()})
		return
	}

	var botData BotRequestData

	if err = json.Unmarshal(requestBody, &botData); err != nil {
		log.Printf("Failed to serialize login request %s\n", err)
		c.JSON(400, gin.H{"message": exception, "error": err.Error()})
		return
	}
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Printf("Failed to instance Bot API: %v\n", err)
		c.JSON(400, gin.H{"message": exception, "error": err})
		return
	}

	msg := tgbotapi.NewMessage(botData.CId, botData.Msg)
	//msg.ReplyToMessageID = int(botData.MId)
	nmsg, err := bot.Send(msg)
	if err != nil {
		log.Printf("Failed to send message: %v\n", err)
		c.JSON(400, gin.H{"message": exception, "error": err})
	} else {
		c.JSON(200, gin.H{"message": fmt.Sprintf("Message sent: %v", nmsg)})
	}
}
