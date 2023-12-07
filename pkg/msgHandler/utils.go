package msgHandler

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strconv"
	"strings"
)

func (h *MSGHandler) ValidateMessage() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := h.bot.GetUpdatesChan(u)
	if err != nil {
		log.Fatalf("error initialize message channel: %s", err.Error())
		return
	}

	for update := range updates {
		if update.CallbackQuery != nil {
			if !h.IsUserRegistrate(update) {
				go h.PersonRegistrate(update)
				continue
			}
			if len(update.CallbackQuery.Data) > 4 {
				if strings.HasPrefix(update.CallbackQuery.Data, "page") {
					str := update.CallbackQuery.Data[len("page"):]
					page, err := strconv.Atoi(str)
					if err == nil {
						go h.ShowQuestPage(update, page)
						continue
					}
				}
			}
			go h.AnotherMessage(update)
			continue
		}

		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() {
			if !h.IsUserRegistrate(update) {
				go h.PersonRegistrate(update)
				continue
			}
			go h.KeyBoard(update)
			switch update.Message.Command() {
			case "start":
			case "help":
				go h.Help(update)
			case "support":
				go h.Support(update)
			case "catalog":
				go h.ShowQuestPage(update, 1)
			default:
				if len(update.Message.Command()) > 4 {
					if update.Message.Command()[0:4] == "info" {
						go h.DetailQuestInfo(update)
						continue
					}
				}
				go h.AnotherMessage(update)
			}
		} else {
			switch update.Message.Text {
			case "📚Каталог":
				if !h.IsUserRegistrate(update) {
					go h.PersonRegistrate(update)
					continue
				}
				go h.ShowQuestPage(update, 1)
			case "🛠️Поддержка":
				if !h.IsUserRegistrate(update) {
					go h.PersonRegistrate(update)
					continue
				}
				go h.Support(update)
			case "🆘🤝Помощь":
				if !h.IsUserRegistrate(update) {
					go h.PersonRegistrate(update)
					continue
				}
				go h.Help(update)
			default:
				if !h.IsUserRegistrate(update) {
					go h.ParseToRegistrate(update)
					continue
				}
				go h.AnotherMessage(update)
			}
		}

	}
}
