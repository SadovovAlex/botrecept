package modules

import (
	"context"
	"fmt"
	"strings"
	"sync"

	//"github.com/SadovovAlex/botrecept/internal/database"

	"github.com/SadovovAlex/botrecept/internal/database"
	"github.com/SadovovAlex/botrecept/internal/modules/menu"
	"github.com/SadovovAlex/botrecept/internal/modules/misc"
	"github.com/SadovovAlex/botrecept/internal/modules/recept"
	"github.com/SadovovAlex/botrecept/internal/modules/stickers"
	"github.com/SadovovAlex/botrecept/internal/utils/helpers"

	"github.com/mymmrac/telego"
	"github.com/mymmrac/telego/telegohandler"
)

var (
	packageLoadersMutex sync.Mutex
	packageLoaders      = map[string]func(*telegohandler.BotHandler, *telego.Bot){
		//  "afk":    afk.Load,
		// "config": config.Load,
		//	"lastfm":   lastfm.Load,
		//	"medias":   medias.Load,
		"menu":     menu.Load,
		"misc":     misc.Load,
		"stickers": stickers.Load,
		//	"sudoers":  sudoers.Load,
		"recept": recept.Load,
	}
)

func BotHandler(ctx context.Context, bot *telego.Bot, updates <-chan telego.Update) (*telegohandler.BotHandler, error) {
	bh, err := telegohandler.NewBotHandler(bot, updates)
	if err != nil {
		return nil, fmt.Errorf("create bot handler: %w", err)
	}

	return bh, nil
}

type Handler struct {
	bot *telego.Bot
	bh  *telegohandler.BotHandler
}

func NewHandler(bot *telego.Bot, bh *telegohandler.BotHandler) *Handler {
	return &Handler{
		bot: bot,
		bh:  bh,
	}
}

func (h *Handler) RegisterHandlers() {
	h.bh.Use(database.SaveUsers)
	h.bh.Use(helpers.CheckDisabled)

	var wg sync.WaitGroup
	done := make(chan struct{}, len(packageLoaders))
	moduleNames := make([]string, 0, len(packageLoaders))

	for name, loadFunc := range packageLoaders {
		wg.Add(1)

		go func(name string, loadFunc func(*telegohandler.BotHandler, *telego.Bot)) {
			defer wg.Done()

			packageLoadersMutex.Lock()
			defer packageLoadersMutex.Unlock()

			loadFunc(h.bh, h.bot)

			done <- struct{}{}
			moduleNames = append(moduleNames, name)
		}(name, loadFunc)
	}

	go func() {
		wg.Wait()
		close(done)
	}()

	for range done {
	}

	joinedModuleNames := strings.Join(moduleNames, ", ")

	fmt.Printf("\033[0;35mModules Loaded:\033[0m %s\n", joinedModuleNames)
}
