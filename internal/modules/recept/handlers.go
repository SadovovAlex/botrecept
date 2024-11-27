package recept

import (
	"github.com/mymmrac/telego"
	"github.com/mymmrac/telego/telegohandler"
	"github.com/mymmrac/telego/telegoutil"
)

func getrecept(bot *telego.Bot, update telego.Update) {

	message := update.Message

	bot.SendMessage(&telego.SendMessageParams{
		ChatID:    telegoutil.ID(message.Chat.ID),
		Text:      "dsds",
		ParseMode: "HTML",
		LinkPreviewOptions: &telego.LinkPreviewOptions{
			PreferLargeMedia: true,
			ShowAboveText:    true,
		},
		ReplyParameters: &telego.ReplyParameters{
			MessageID: message.MessageID,
		},
	})

}

func Load(bh *telegohandler.BotHandler, bot *telego.Bot) {
	bh.Handle(getrecept, telegohandler.Or(
		telegohandler.CommandEqual("recept"),
		telegohandler.CallbackDataPrefix("recept"),
	))
}
