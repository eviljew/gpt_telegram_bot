package main

import (
	"flag"
	"log"
	"telegrambot_zerocoder/gpt"
	"telegrambot_zerocoder/master"
	"telegrambot_zerocoder/telegram"
)

var (
	tgToken  = flag.String("tg-token", "", "Token for telegram bot api")
	gptToken = flag.String("gpt-token", "", "Token for gpt api")
)

func init() {
	flag.Parse()

	if *tgToken == "" {
		log.Fatal("telegram token is not specified")
	}

	if *gptToken == "" {
		log.Fatal("gpt token is not specified")
	}
}

func main() {
	// получение telegram токена из входных параметров запуска и запуск клиента Telegram бота
	tg := telegram.New(*tgToken)

	// получение telegram токена из входных параметров запуска и запуск клиента Telegram бота
	g := gpt.New(*gptToken)

	// Создание образа мастера программы, он руководит работой бота
	m := master.New(tg, g)

	// Запуск сервера веб-хука Telegram бота
	m.Updater()
}
