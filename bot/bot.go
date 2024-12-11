package main

import (
	"fmt"
	"golang_practice/bot/utils"
	"github.com/mymmrac/telego"
)

type MyBot struct {
	Token string
	ModeratorChatId telego.ChatID
}

func (m *MyBot) Bot() *telego.Bot{
	bot, err := telego.NewBot(m.Token)
	if err != nil{
		fmt.Println(err)
	}
	return bot
}

func (m *MyBot) SendAlert (text string){
	if text == ""{
		text = "Bot started"
	}
	msg := &telego.SendMessageParams{ 
		ChatID: m.ModeratorChatId, 
		Text: text,
	}

	bot := *m.Bot()
	bot.SendMessage(msg)
}

func (m *MyBot) SendMsg (chatId telego.ChatID, text string){
	param := &telego.SendMessageParams{
		ChatID: chatId,
		Text: text,
	}
	m.Bot().SendMessage(param)
}

func main(){
	bot := MyBot{Token: utils.ApiKey, ModeratorChatId: telego.ChatID{ID: utils.ModeratorId}}
	fmt.Println(bot.Bot())
	ollama := utils.Ollama{Model: "llama3.2:1b", Stream: false}
	ollama.StartServer()
	bot.SendAlert("")

	updates, _ := bot.Bot().UpdatesViaLongPolling(nil)
	defer bot.Bot().StopLongPolling()
	for update := range updates {
		if update.Message != nil {
			switch update.Message.Text{
			case "/start":
				bot.SendMsg(update.Message.Chat.ChatID(), "Hello")
			default:
				query := utils.Query{Ollama: &ollama, Prompt: update.Message.Text}
				answer :=  query.Ask()
				bot.SendMsg(update.Message.Chat.ChatID(), string(answer))
			}

			
		}
	}
}