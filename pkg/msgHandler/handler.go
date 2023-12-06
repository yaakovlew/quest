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

func (h *MSGHandler) Help(message *tgbotapi.Message) error {
	msgText := "/support - команда для перехода в поддержку\n/catalog - команда для отображения каталога квестов"
	msg := tgbotapi.NewMessage(message.Chat.ID, msgText)
	if _, err := h.bot.Send(msg); err != nil {
		return err
	}

	return nil
}

func (h *MSGHandler) AnotherMessage(message *tgbotapi.Message) error {
	msgText := "/support - команда для переходу в поддержку"
	msg := tgbotapi.NewMessage(message.Chat.ID, msgText)
	if _, err := h.bot.Send(msg); err != nil {
		return err
	}

	return nil
}

func (h *MSGHandler) KeyBoard(message *tgbotapi.Message) error {
	btnCatalog1 := tgbotapi.NewKeyboardButton("📚Каталог")
	btnCatalog2 := tgbotapi.NewKeyboardButton("🛠️Поддержка")
	btnCatalog3 := tgbotapi.NewKeyboardButton("🆘🤝Помощь")
	msg := tgbotapi.NewMessage(message.Chat.ID, "Введите имя, телефон и возраст через запятую")
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

func (h *MSGHandler) Support(message *tgbotapi.Message) error {
	msgText := os.Getenv("TG_SUPPORT_LYNC")
	msg := tgbotapi.NewMessage(message.Chat.ID, msgText)
	if _, err := h.bot.Send(msg); err != nil {
		return err
	}

	return nil
}

func (h *MSGHandler) DetailQuestInfo(message *tgbotapi.Message) error {
	commandText := strings.TrimPrefix(message.Text, "/info-")
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

	msg := tgbotapi.NewMessage(message.Chat.ID, msgText)
	if _, err := h.bot.Send(msg); err != nil {
		return err
	}

	return nil
}

func (h *MSGHandler) ShowQuestPage(message *tgbotapi.Message, currentPage int) error {
	page, _ := h.service.GetPageAmount()

	inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("1", "button1"),
			tgbotapi.NewInlineKeyboardButtonData("Окно 2", "button2"),
			tgbotapi.NewInlineKeyboardButtonData("Окно 3", "button3"),
		),
	)
	msg := tgbotapi.NewMessage(message.Chat.ID, strconv.Itoa(page))
	msg.ReplyMarkup = inlineKeyboard
	if _, err := h.bot.Send(msg); err != nil {
		return err
	}

	return nil
}
