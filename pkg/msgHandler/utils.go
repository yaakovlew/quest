package msgHandler

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
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
		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() {
			h.KeyBoard(update.Message)
			switch update.Message.Command() {
			case "start":
			case "help":
				h.Help(update.Message)
			case "support":
				h.Support(update.Message)
			case "catalog":
				h.ShowQuestPage(update.Message, 0)
			case "info":
				h.DetailQuestInfo(update.Message)
			default:
				h.AnotherMessage(update.Message)
			}
		} else {
			switch update.Message.Text {
			case "📚Каталог":
				h.ShowQuestPage(update.Message, 0)
			case "🛠️Поддержка":
				h.Support(update.Message)
			case "🆘🤝Помощь":
				h.Help(update.Message)
			default:
				h.AnotherMessage(update.Message)
			}
		}

	}
}
