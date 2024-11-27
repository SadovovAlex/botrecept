package recept

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"

	"github.com/mymmrac/telego"
	"github.com/mymmrac/telego/telegohandler"
	"github.com/mymmrac/telego/telegoutil"
)

type Recipe struct {
	Name         string   `json:"name"`
	Ingredients  []string `json:"ingredients"`
	Instructions []string `json:"instructions"`
}

func getrecept(bot *telego.Bot, update telego.Update) {
	message := update.Message

	// Загрузка рецептов из JSON-файла
	recipes, err := loadRecipes("recept.json")
	if err != nil {
		fmt.Println("Ошибка при загрузке рецептов:", err)
		bot.SendMessage(&telego.SendMessageParams{
			ChatID:    telegoutil.ID(message.Chat.ID),
			Text:      "Произошла ошибка при загрузке рецептов.",
			ParseMode: "HTML",
		})
		return
	}

	// Выбор случайного рецепта
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(recipes))
	randomRecipe := recipes[randomIndex]

	// Формирование ответа с выбранным рецептом
	response := fmt.Sprintf("<b>%s</b>\nИнгредиенты:\n%s\nИнструкции:\n%s\n",
		randomRecipe.Name, strings.Join(randomRecipe.Ingredients, ", "), strings.Join(randomRecipe.Instructions, "\n"))
	bot.SendMessage(&telego.SendMessageParams{
		ChatID:    telegoutil.ID(message.Chat.ID),
		Text:      response,
		ParseMode: "HTML",
	})
}

func loadRecipes(filename string) ([]Recipe, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var recipes []Recipe
	err = json.Unmarshal(data, &recipes)
	if err != nil {
		return nil, err
	}

	return recipes, nil
}

func Load(bh *telegohandler.BotHandler, bot *telego.Bot) {
	bh.Handle(getrecept, telegohandler.Or(
		telegohandler.CommandEqual("recept"),
		telegohandler.CallbackDataPrefix("recept"),
	))
}
