package telegram

import (
	"fmt"
	"log"
	"path/filepath"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/jeliot12/psychologyBot/module"
	"github.com/jeliot12/psychologyBot/module/getplaylist"
)

const commandStart = "start"

const YOUTUBE_DEFAULT_PLAYLIST_URL = "https://www.youtube.com/playlist?list="

var global int
var TitlePlayList string
var IdPlayList string

var numericKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("JavaScript"),
		tgbotapi.NewKeyboardButton("C#"),
		tgbotapi.NewKeyboardButton("Python"),
	),
)

func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Такой команды нету :(")

	switch message.Command() {
	case commandStart:
		msg.Text = `Приветствую, я помогу тебе найти полезные курсы
и roadmap-ы по нужному языку программированию.`
		msg.ReplyMarkup = numericKeyboard
		_, err := b.bot.Send(msg)
		return err
	default:
		_, err := b.bot.Send(msg)
		return err
	}
}

func (b *Bot) handleMessage(message *tgbotapi.Message) {
	items, errors := module.RetrieveTitle("UCO-JVU-S4o_25knhbPyDrgQ")
	if errors != nil {
		log.Panic(errors)
	}
	item, err := getplaylist.RetrievePlayList("UCO-JVU-S4o_25knhbPyDrgQ")
	if err != nil {
		log.Panic(err)
	}
	if len(items) < 1 {
		fmt.Println("")
	}
	if len(item) < 1 {
		fmt.Println("")
	}
	global = len(item)

	// log.Printf("[%s] %s", message.From.UserName, message.Text)
	if message.Text != "JavaScript" {
		if message.Text != "C#" {
			if message.Text != "Python" {
				msg := tgbotapi.NewMessage(message.Chat.ID, "Я не понял вашу команду :(")
				b.bot.Send(msg)
			}
		}
	}
	if message.Text == "JavaScript" {
		for i := 0; i < global; i++ {
			if items[i].Snippet.Title == "JavaScript" {
				IdPlayList = YOUTUBE_DEFAULT_PLAYLIST_URL + item[1].Id
				msg := tgbotapi.NewMessage(message.Chat.ID, IdPlayList)
				b.bot.Send(msg)
				photo_roadmap := tgbotapi.NewPhotoShare(message.Chat.ID, "https://media.tproger.ru/uploads/2022/05/JavaScript.png")
				b.bot.Send(photo_roadmap)
				fname := "file/javascript/useful.txt"
				abs_fname, err := filepath.Abs(fname)
				if err != nil {
					log.Fatal(err)
				}
				file_upl := tgbotapi.NewDocumentUpload(message.Chat.ID, abs_fname)
				b.bot.Send(file_upl)
				break
			}
		}
	}
	if message.Text == "C#" {
		for i := 0; i < global; i++ {
			if items[i].Snippet.Title == "C#" {
				IdPlayList = YOUTUBE_DEFAULT_PLAYLIST_URL + item[0].Id
				msg := tgbotapi.NewMessage(message.Chat.ID, IdPlayList)
				b.bot.Send(msg)
				fname := "file/csharp/useful.txt"
				abs_fname, err := filepath.Abs(fname)
				if err != nil {
					log.Fatal(err)
				}
				file_upload := tgbotapi.NewDocumentUpload(message.Chat.ID, abs_fname)
				b.bot.Send(file_upload)
				break
			}
		}
	}
	if message.Text == "Python" {
		msg := tgbotapi.NewMessage(message.Chat.ID, "Для этого языка roadmap и полезные ресурсы в сборе.")
		b.bot.Send(msg)
	}
}
