package msgHandler

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"os"
	"quest/pkg/service"
	"strconv"
	"strings"
)

type MSGHandler struct {
	service *service.Service
	bot     *tgbotapi.BotAPI
}

func NewMSGHandler(service *service.Service, bot *tgbotapi.BotAPI) *MSGHandler {
	return &MSGHandler{
		service: service,
		bot:     bot,
	}
}

func (h *MSGHandler) Help(update tgbotapi.Update) error {
	msgText := "/support - команда для перехода в поддержку\n/catalog - команда для отображения каталога квестов"
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, msgText)
	if _, err := h.bot.Send(msg); err != nil {
		return err
	}

	return nil
}

func (h *MSGHandler) AnotherMessage(update tgbotapi.Update) error {
	msgText := "/support - команда для переходу в поддержку"
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, msgText)
	if _, err := h.bot.Send(msg); err != nil {
		return err
	}

	return nil
}

func (h *MSGHandler) KeyBoard(update tgbotapi.Update) error {
	btnCatalog1 := tgbotapi.NewKeyboardButton("📚Каталог")
	btnCatalog2 := tgbotapi.NewKeyboardButton("🛠️Поддержка")
	btnCatalog3 := tgbotapi.NewKeyboardButton("🆘🤝Помощь")
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Меню:")
	msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			btnCatalog1,
			btnCatalog2,
			btnCatalog3,
		),
	)

	if _, err := h.bot.Send(msg); err != nil {
		return err
	}

	return nil
}

func (h *MSGHandler) Support(update tgbotapi.Update) error {
	msgText := os.Getenv("TG_SUPPORT_LYNC")
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, msgText)
	if _, err := h.bot.Send(msg); err != nil {
		return err
	}

	return nil
}

func (h *MSGHandler) DetailQuestInfo(update tgbotapi.Update) error {
	commandText := strings.TrimPrefix(update.Message.Text, "/info")
	questId, err := strconv.Atoi(commandText)
	if err != nil {
		return err
	}

	quest, err := h.service.GetQuest(questId)
	if err != nil {
		return err
	}

	msgText := fmt.Sprintf(`
Название: %s
Применение %s
Цели: %s
Описание: %s
Возрастное ограничение: %d лет
Сложность: %s
Длительность: %d минут 
Примерная локация: %s
Организатор: %s`,
		quest.Name, quest.AuthorComment, quest.Point, quest.Description, quest.AgeLevel,
		quest.Difficult, quest.Duration, quest.Location, quest.Organizer,
	)

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, msgText)
	if _, err := h.bot.Send(msg); err != nil {
		return err
	}

	return nil
}

func (h *MSGHandler) ShowQuestPage(update tgbotapi.Update, currentPage int) error {
	inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(),
	)
	inlineKeyboard.InlineKeyboard[0] = h.leafThroughPage(currentPage)

	quests, err := h.service.GetQuestsByPage(currentPage)
	if err != nil {
		return err
	}

	msgText := ""
	for i, quest := range quests {
		msgText = msgText + fmt.Sprintf("Название: %s\nОписание: %s\nСсылка: /info%d",
			quest.Name, quest.Description, quest.Id)
		if i != len(quests) {
			msgText = msgText + "\n\n"
		}
	}

	var chatId int64
	if update.CallbackQuery != nil {
		chatId = update.CallbackQuery.Message.Chat.ID
	} else if update.Message != nil {
		chatId = update.Message.Chat.ID
	} else {
		return fmt.Errorf("not found user")
	}

	msg := tgbotapi.NewMessage(chatId, msgText)
	msg.ReplyMarkup = inlineKeyboard
	if _, err := h.bot.Send(msg); err != nil {
		return err
	}

	return nil
}

func (h *MSGHandler) leafThroughPage(currentPage int) []tgbotapi.InlineKeyboardButton {
	page, err := h.service.GetPageAmount()
	if err != nil {
		return nil
	}
	var buttons []tgbotapi.InlineKeyboardButton
	if page != 1 && currentPage != 1 {
		button := "page1"
		buttons = append(
			buttons,
			tgbotapi.InlineKeyboardButton{
				Text:         "1",
				CallbackData: &button,
			},
		)
	}

	if currentPage != 1 {
		button := fmt.Sprintf("page%d", currentPage-1)
		buttons = append(
			buttons,
			tgbotapi.InlineKeyboardButton{
				Text:         "<",
				CallbackData: &button,
			},
		)
	}

	if currentPage < page {
		button := fmt.Sprintf("page%d", currentPage+1)
		buttons = append(
			buttons,
			tgbotapi.InlineKeyboardButton{
				Text:         ">",
				CallbackData: &button,
			},
		)
	}

	if page != 1 && currentPage != page {
		button := fmt.Sprintf("page%d", page)
		buttons = append(
			buttons,
			tgbotapi.InlineKeyboardButton{
				Text:         strconv.Itoa(page),
				CallbackData: &button,
			},
		)
	}

	return buttons
}
