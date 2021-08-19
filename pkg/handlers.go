package pkg

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	commandStart = "start"
	commandTemp = "temp"
	commandHelp = "help"
	commandError = "Введите команду, начинающаяся с символа '/'"
)
func (b *Bot) handleCommand(message *tgbotapi.Message) error{
	switch message.Command(){
	case commandStart:
		return b.handleStartCommand(message)
	case commandTemp:
		return b.handleTempCommand(message)
	case commandHelp:
		return b.handleHelpCommand(message)
	default:
		return b.handleUnknownCommand(message)
	}
}

func (b *Bot) handleMessage (message *tgbotapi.Message) {
	log.Printf("[%s] %s", message.From.UserName, message.Text)
	msg := tgbotapi.NewMessage(message.Chat.ID, commandError)
	b.bot.Send(msg)
}

func (b *Bot) handleTempCommand(message *tgbotapi.Message) error{
	if len(message.Text) <= 5{
		text := "Введите название Города после команды /temp"
		msg := tgbotapi.NewMessage(message.Chat.ID,text)
		_, err := b.bot.Send(msg)
		return err
	}
	message.Text = message.Text[6:]
	log.Println(message.Text)
	resp, err := http.Get("https://api.openweathermap.org/data/2.5/weather?q=" + message.Text + "&APPID=4b106504c938739971ea66f159abff1d&units=metric")
	log.Println(err)
	if err != nil {
		log.Println(err)
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
		return err
	}

	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err!=nil{
		log.Fatal(err)
	}

	if result["message"] == "city not found"{
		text := "Город не найден"
		msg := tgbotapi.NewMessage(message.Chat.ID,text)
		_, err = b.bot.Send(msg)
		return nil
	}
	r := result["main"].(map[string]interface{})
	i := fmt.Sprintf("%.0f", r["temp"])
	text := "В городе " + message.Text + " " + i + " градусов по цельсию"
	msg := tgbotapi.NewMessage(message.Chat.ID,text)
	_, err = b.bot.Send(msg)
	return err
}

func (b *Bot) handleStartCommand(message *tgbotapi.Message) error{
	msg := tgbotapi.NewMessage(message.Chat.ID, "Введите /temp 'Город' для вывода температура в выбранном городе")
	b.bot.Send(msg)
	return nil
}
func (b *Bot) handleHelpCommand(message *tgbotapi.Message) error{
	msg := tgbotapi.NewMessage(message.Chat.ID, "Введите /temp 'Город' для вывода температура в выбранном городе")
	b.bot.Send(msg)
	return nil
}
func (b *Bot) handleUnknownCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Неизвестная команда, введите /help для помощи")
	_, err := b.bot.Send(msg)
	return err
}
