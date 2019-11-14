package main

const (
	TELEGRAM_API_URL = "https://api.telegram.org/"
)

type telegramBot struct {
	baseUrl string
}

func NewTelegramBot(apiKey string) *telegramBot {
	baseUrl := TELEGRAM_API_URL + "bot" + apiKey + "/"
	t := telegramBot{
		baseUrl: baseUrl,
	}
	return &t
}

