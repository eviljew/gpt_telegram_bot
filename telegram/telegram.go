package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

type Telegram struct {
	token  string
	client *tgbotapi.BotAPI
}

func New(token string) *Telegram {
	client, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal(err)
	}

	return &Telegram{
		token:  token,
		client: client,
	}
}

func (tg *Telegram) GetFileIfExists(fileId string) (string, error) {
	if fileId == "" {
		return "", nil
	}

	fileUrl, err := tg.client.GetFileDirectURL(fileId)
	if err != nil {
		return "", err
	}

	return fileUrl, nil
}

func (tg *Telegram) StartUpdater() tgbotapi.UpdatesChannel {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	return tg.client.GetUpdatesChan(u)
}

// Отправка текстового сообщения ботом пользователю
func (tg *Telegram) Send(chatID int64, msg string) (string, error) {
	send, err := tg.client.Send(tgbotapi.NewMessage(chatID, msg))
	if err != nil {
		return "", err
	}

	return send.Text, nil
}

// Получение id файла из сообщения пользователя, если таковой имеется.
// Используется для получения ссылки на файл через Telegram api
func (tg *Telegram) GetFileIDIfExists(update tgbotapi.Update) string {
	if update.Message.Photo == nil {
		return ""
	}

	return update.Message.Photo[0].FileID
}
