package telegram

import (
	"github.com/SadovovAlex/botrecept/internal/config"

	"github.com/mymmrac/telego"
)

func CreateBot() (*telego.Bot, error) {
	var err error
	bot, err := telego.NewBot(config.TelegramToken)
	if config.BotAPIURL != "" {
		bot, err = telego.NewBot(config.TelegramToken, telego.WithAPIServer(config.BotAPIURL))
	}
	return bot, err
}
