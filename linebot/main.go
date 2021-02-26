package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/line/line-bot-sdk-go/linebot"
)

func saveImg(content io.ReadCloser) {
	img, name, err := image.Decode(content)
	if err != nil {
		fmt.Println(err)
	}
	f, err := os.Create("img/" + time.Now().String() + "." + name)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	jpeg.Encode(f, img, nil)
}

func callback(w http.ResponseWriter, r *http.Request) {
	bot, err := linebot.New(
		os.Getenv("CHANNEL_SECRET"),
		os.Getenv("CHANNEL_ACCESS_TOKEN"),
	)
	if err != nil {
		log.Fatal(err)
	}
	events, err := bot.ParseRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			action := linebot.NewDatetimePickerAction(
				"Select date",
				"storeId=12345",
				"datetime",
				"2017-12-25t00:00",
				"2018-01-24t23:59",
				"2017-12-25t00:00",
			)
			switch msg := event.Message.(type) {
			case *linebot.TextMessage:
				if msg.Text == "Hi" {
					template := linebot.NewConfirmTemplate("Hello World", action, action)
					message := linebot.NewTemplateMessage("Sorry :(, please update your app.", template)
					_, err := bot.ReplyMessage(event.ReplyToken, message).Do()
					if err != nil {
						fmt.Println(err)
					}
				} else {
					_, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(msg.Text)).Do()
					if err != nil {
						w.WriteHeader(http.StatusBadRequest)
						fmt.Println(err)
					}
					w.WriteHeader(http.StatusOK)
				}

			case *linebot.ImageMessage:
				content, _ := bot.GetMessageContent(msg.ID).Do()
				defer content.Content.Close()
				fmt.Println(content.ContentType)
				saveImg(content.Content)

			default:
				_, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("hi")).Do()
				if err != nil {
					w.WriteHeader(http.StatusBadRequest)
					fmt.Println(err)
				}
				w.WriteHeader(http.StatusOK)
			}
		} else if event.Type == linebot.EventTypePostback {
			j, _ := event.MarshalJSON()
			fmt.Println(string(j))

			// _, err := bot.
			// if err != nil {
			// 	fmt.Println(err)
			// }
		} else {
			fmt.Println(event.Type)
		}
	}
}

func main() {
	http.HandleFunc("/callback", callback)
	err := http.ListenAndServe(":7000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
