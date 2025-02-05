package qr

import (
	"fmt" // Не забудьте импортировать пакет image
	"image"
	"log"
	"net/http"

	"github.com/SadovovAlex/botrecept/internal/utils/helpers"
	"github.com/liyue201/goqr"
	"github.com/mymmrac/telego"
	"github.com/mymmrac/telego/telegohandler"
	"github.com/mymmrac/telego/telegoutil"
)

const telegramAPIURL = "https://api.telegram.org/file/bot%s/%s"

func scanQR(bot *telego.Bot, update telego.Update) {
	message := update.Message

	// Логируем все содержимое сообщения
	log.Printf("Received message: %+v\n\n", message)

	// Проверяем, есть ли фото в сообщении
	if message.Photo != nil && len(message.Photo) > 0 {
		// Получаем ID самого большого фото
		fileID := (message.Photo)[len(message.Photo)-1].FileID
		log.Printf("fileID: %+v\n\n", fileID)

		// Получаем файл
		file, err := bot.GetFile(&telego.GetFileParams{FileID: fileID})
		if err != nil {
			log.Printf("получение err: %+v\n\n", err)
			bot.SendMessage(&telego.SendMessageParams{
				ChatID: telegoutil.ID(message.Chat.ID),
				Text:   "Ошибка при получении fileID.",
			})
			return
		}

		// Скачиваем файл
		// Формируем полный URL для скачивания файла
		filePath := file.FilePath
		fullURL := fmt.Sprintf(telegramAPIURL, bot.Token(), filePath)

		log.Printf("fileID: %+v\n\n", fullURL)
		resp, err := http.Get(fullURL)
		if err != nil {
			log.Printf("fileID: %+v\n\n", err)
			bot.SendMessage(&telego.SendMessageParams{
				ChatID: telegoutil.ID(message.Chat.ID),
				Text:   "Ошибка при скачивании файла.",
			})
			return
		}
		defer resp.Body.Close()

		//Декодируем изображение
		img, _, err := image.Decode(resp.Body)
		if err != nil { // Исправлено: добавлен if для проверки ошибки
			bot.SendMessage(&telego.SendMessageParams{
				ChatID: telegoutil.ID(message.Chat.ID),
				Text:   "Ошибка при декодировании изображения.",
			})
			return
		}

		// Распознаем QR-код
		//qrCode, err := qrcode.Decode(resp.Body)
		qrCode, err := goqr.Recognize(img)
		if err != nil {
			log.Printf("fileID: %+v\n\n", err)
			bot.SendMessage(&telego.SendMessageParams{
				ChatID: telegoutil.ID(message.Chat.ID),
				Text:   "Ошибка при распознавании QR-кода.",
			})
			return
		}

		// Формируем ответ с данными из QR-кода
		//response := fmt.Sprintf("<b>Содержимое QR-кода:</b>\n%s", qrCode.Content)
		var qrPayload string

		for _, code := range qrCode { // Предполагается, что qrCodes - это срез QR-кодов
			qrPayload += string(code.Payload) // Собираем все Payload в одну строку с переносами
		}
		fmt.Printf("qrCode text: %s\n", string(qrPayload))               // Преобразуем []uint8 в строку
		response := fmt.Sprintf("<b>QR-код:</b>\n%s", string(qrPayload)) //

		bot.SendMessage(&telego.SendMessageParams{
			ChatID:    telegoutil.ID(message.Chat.ID),
			Text:      response,
			ParseMode: "HTML",
		})
	} else {
		bot.SendMessage(&telego.SendMessageParams{
			ChatID: telegoutil.ID(message.Chat.ID),
			Text:   "Пожалуйста, отправьте фото с QR-кодом.",
		})
	}
}

func Load(bh *telegohandler.BotHandler, bot *telego.Bot) {
	helpers.Store("qr")
	bh.Handle(scanQR, telegohandler.Or(
		//telegohandler.CommandEqual("recept"),
		//telegohandler.CallbackDataPrefix("recept"),
		telegohandler.AnyMessage(),
	))
}
