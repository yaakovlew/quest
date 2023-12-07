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
			if len(update.CallbackQuery.Data) > 4 {
				if strings.HasPrefix(update.CallbackQuery.Data, "page") {
					str := update.CallbackQuery.Data[len("page"):]
					page, err := strconv.Atoi(str)
					if err == nil {
						h.ShowQuestPage(update, page)
						continue
					}
				}
			}
			h.AnotherMessage(update)
			continue
		}

		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() {
			h.KeyBoard(update)
			switch update.Message.Command() {
			case "start":
			case "help":
				h.Help(update)
			case "support":
				h.Support(update)
			case "catalog":
				h.ShowQuestPage(update, 1)
			default:
				if len(update.Message.Command()) > 4 {
					if update.Message.Command()[0:4] == "info" {
						h.DetailQuestInfo(update)
						continue
					}
				}
				h.AnotherMessage(update)
			}
		} else {
			switch update.Message.Text {
			case "📚Каталог":
				h.ShowQuestPage(update, 1)
			case "🛠️Поддержка":
				h.Support(update)
			case "🆘🤝Помощь":
				h.Help(update)
			default:
				h.AnotherMessage(update)
			}
		}

	}
}
