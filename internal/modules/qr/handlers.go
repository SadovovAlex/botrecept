package qr_test

import (
	"fmt"

	"github.com/SadovovAlex/botrecept/internal/utils/helpers"
	"github.com/mymmrac/telego"
	"github.com/mymmrac/telego/telegohandler"
	"github.com/mymmrac/telego/telegoutil"
)

func scanQR(bot *telego.Bot, update telego.Update) {
	message := update.Message

	// Формирование ответа с выбранным рецептом
	response := fmt.Sprintf("<b>%s</b>\nИнгредиенты:\n%s\nИнструкции:\n%s\n",
		"zzzz\n")
	bot.SendMessage(&telego.SendMessageParams{
		ChatID:    telegoutil.ID(message.Chat.ID),
		Text:      response,
		ParseMode: "HTML",
	})
}

func Load(bh *telegohandler.BotHandler, bot *telego.Bot) {
	helpers.Store("qr")
	bh.Handle(scanQR, telegohandler.Or(
		telegohandler.CommandEqual("recept"),
		telegohandler.CallbackDataPrefix("recept"),
	))
}
