package msgHandler

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"quest/pkg/models"
)

func (h *MSGHandler) IsUserRegistrate(update tgbotapi.Update) bool {
	var userId int64
	if update.Message != nil {
		userId = update.Message.Chat.ID
	}

	if update.CallbackQuery != nil {
		userId = update.CallbackQuery.Message.Chat.ID
	}

	user, err := h.service.FindUser(int(userId))
	if err != nil || user == (models.User{}) {
		return false
	}

	return true
}
