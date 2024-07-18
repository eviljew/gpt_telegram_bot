package master

import (
	"log"
	"telegrambot_zerocoder/gpt"
	"telegrambot_zerocoder/lib/e"
	"telegrambot_zerocoder/telegram"
)

type Master struct {
	tg  *telegram.Telegram
	gpt *gpt.Gpt
}

func New(tg *telegram.Telegram, gpt *gpt.Gpt) *Master {
	return &Master{
		tg:  tg,
		gpt: gpt,
	}
}

func (m *Master) Updater() {
	//Получаем обновления от бота
	updates := m.tg.StartUpdater()

	for update := range updates {
		if update.Message == nil {
			continue
		}

		fileId := m.tg.GetFileIDIfExists(update)
		filePath, err := m.tg.GetFileIfExists(fileId)
		if err != nil {
			log.Println(e.Wrap("get photo file id", err))
		}

		answer, err := m.gpt.Send(update.Message.Text, filePath)
		if err != nil {
			log.Println(e.Wrap("send message to gpt", err))
		}

		if _, err := m.tg.Send(update.Message.Chat.ID, answer); err != nil {
			log.Println(e.Wrap("send message to user", err))
		}
	}
}
