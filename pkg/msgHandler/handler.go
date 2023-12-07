package msgHandler

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"os"
	"quest/pkg/models"
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
	var chatId int64
	if update.CallbackQuery != nil {
		chatId = update.CallbackQuery.Message.Chat.ID
	} else if update.Message != nil {
		chatId = update.Message.Chat.ID
	} else {
		return fmt.Errorf("not found user")
	}

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
	if quests == nil {
		msgText = "Квестов пока нет"
		msg := tgbotapi.NewMessage(chatId, msgText)
		if _, err := h.bot.Send(msg); err != nil {
			return err
		}

		return nil
	}

	inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(),
	)
	inlineKeyboard.InlineKeyboard[0] = h.leafThroughPage(currentPage)

	msg := tgbotapi.NewMessage(chatId, msgText)
	msg.ReplyMarkup = inlineKeyboard
	if _, err := h.bot.Send(msg); err != nil {
		return err
	}

	return nil
}

func (h *MSGHandler) PersonRegistrate(update tgbotapi.Update) error {
	msgText := "Введите имя, телефон и возраст через запятую"
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, msgText)
	if _, err := h.bot.Send(msg); err != nil {
		return err
	}

	return nil
}

func (h *MSGHandler) ParseToRegistrate(update tgbotapi.Update) error {
	user, err := h.validateUser(update)
	if err != nil {
		go h.PersonRegistrate(update)
		return err
	}

	if _, err := h.service.CreateUser(user); err != nil {
		go h.PersonRegistrate(update)
		return err
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Успешно!")
	if _, err := h.bot.Send(msg); err != nil {
		go h.PersonRegistrate(update)
		return err
	}

	return nil
}

func (h *MSGHandler) validateUser(update tgbotapi.Update) (models.User, error) {
	output := strings.ReplaceAll(update.Message.Text, " ", "")
	str := strings.Split(output, ",")
	if len(str) != 3 {
		return models.User{}, fmt.Errorf("not valid message")
	}

	age, err := strconv.Atoi(str[2])
	if err != nil {
		return models.User{}, fmt.Errorf("not valid message: %s", err.Error())
	}

	if age < 1 || age > 125 {
		return models.User{}, fmt.Errorf("not valid message")
	}

	if len(str[0]) < 2 || len(str[0]) > 255 {
		return models.User{}, fmt.Errorf("not valid message")
	}

	if len(str[1]) != 12 {
		return models.User{}, fmt.Errorf("not valid message")
	}

	if str[1][:2] != "+7" {
		return models.User{}, fmt.Errorf("not valid message")
	}

	_, err = strconv.Atoi(str[1][2:])
	if err != nil {
		return models.User{}, fmt.Errorf("not valid message: %s", err.Error())
	}

	user := models.User{
		TgUserId: int(update.Message.Chat.ID),
		Name:     str[0],
		Age:      age,
		Phone:    str[1],
	}

	return user, nil
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
