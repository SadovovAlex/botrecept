package sudoers

import (
	"fmt"
	"log"
	"strings"

	"github.com/SadovovAlex/botrecept/internal/config"
	"github.com/SadovovAlex/botrecept/internal/database"
	"github.com/SadovovAlex/botrecept/internal/localization"
	"github.com/SadovovAlex/botrecept/internal/utils"
	"github.com/mymmrac/telego"
	"github.com/mymmrac/telego/telegohandler"
	"github.com/mymmrac/telego/telegoutil"
)

var announceMessageText string

func announce(bot *telego.Bot, update telego.Update) {
	var lang string
	message := update.Message

	if message == nil {
		message = update.CallbackQuery.Message.(*telego.Message)
		lang = strings.ReplaceAll(update.CallbackQuery.Data, "announce ", "")
	}

	if (message == nil || message.From.ID != config.OwnerID) &&
		(update.CallbackQuery == nil || update.CallbackQuery.From.ID != config.OwnerID) {
		return
	}

	if lang == "" {
		buttons := make([][]telego.InlineKeyboardButton, 0, len(database.AvailableLocales))
		for _, lang := range database.AvailableLocales {
			loaded, ok := localization.LangBundles[lang]
			if !ok {
				log.Fatalf("Language '%s' not found in the cache.", lang)
			}
			languageFlag, _, _ := loaded.FormatMessage("language-flag")
			languageName, _, _ := loaded.FormatMessage("language-name")

			buttons = append(buttons, []telego.InlineKeyboardButton{{
				Text: languageFlag +
					languageName,
				CallbackData: fmt.Sprintf("announce %s", lang),
			}})
		}

		bot.SendMessage(&telego.SendMessageParams{
			ChatID:      telegoutil.ID(config.OwnerID),
			Text:        "Choose a language:",
			ParseMode:   "HTML",
			ReplyMarkup: telegoutil.InlineKeyboard(buttons...),
		})
		announceMessageText = utils.FormatText(message.Text, message.Entities)
		return
	}

	messageFields := strings.Fields(announceMessageText)
	if len(messageFields) < 2 {
		return
	}

	announceType := messageFields[1]
	announceMessageText = strings.Replace(announceMessageText, messageFields[0], "", 1)
	var query string

	switch announceType {
	case "groups":
		announceMessageText = strings.Replace(announceMessageText, announceType, "", 1)
		query = fmt.Sprintf("SELECT id FROM groups WHERE language = '%s';", lang)
	case "users":
		announceMessageText = strings.Replace(announceMessageText, announceType, "", 1)
		query = fmt.Sprintf("SELECT id FROM users WHERE language = '%s';", lang)
	default:
		query = fmt.Sprintf("SELECT id FROM users WHERE language = '%s' UNION ALL SELECT id FROM groups WHERE language = '%s';", lang, lang)
	}

	rows, err := database.DB.Query(query)
	if err != nil {
		return
	}
	defer rows.Close()

	var successCount, errorCount int

	for rows.Next() {
		var chatID int64
		if err := rows.Scan(&chatID); err != nil {
			continue
		}
		_, err := bot.SendMessage(&telego.SendMessageParams{
			ChatID:    telegoutil.ID(chatID),
			Text:      announceMessageText,
			ParseMode: "HTML",
		})
		if err != nil {
			errorCount++
			continue
		}

		successCount++
	}

	if err := rows.Err(); err != nil {
		return
	}

	bot.EditMessageText(&telego.EditMessageTextParams{
		ChatID:    telegoutil.ID(config.OwnerID),
		MessageID: update.CallbackQuery.Message.GetMessageID(),
		Text:      fmt.Sprintf("<b>Messages sent successfully:</b> <code>%d</code>\n<b>Messages unsent:</b> <code>%d</code>", successCount, errorCount),
		ParseMode: "HTML",
	})
	announceMessageText = ""
}

func Load(bh *telegohandler.BotHandler, bot *telego.Bot) {
	bh.Handle(announce, telegohandler.Or(
		telegohandler.CommandEqual("announce"),
		telegohandler.CallbackDataPrefix("announce"),
	))
}
