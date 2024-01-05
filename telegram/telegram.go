package telegram

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Message struct {
	Text    string `json:"text"`
	GroupId int64  `json:"group_id"`
}

type ErrorMessage struct {
	Error string `json:"error"`
}

func SendTelegram(w http.ResponseWriter, r *http.Request) {
	token := os.Getenv("TELEGRAM_BOT_TOKEN")
	var errorMessage ErrorMessage
	message := Message{}
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		errorMessage.Error = "Erro ao decodificar o JSON"
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorMessage)
		return
	}
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Printf("Erro ao criar o bot: %v", err)
		errorMessage.Error = "Erro ao criar o bot"
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorMessage)
		return
	}
	alertText := tgbotapi.NewMessage(message.GroupId, message.Text)
	retorno, err := bot.Send(alertText)
	if err != nil {
		log.Printf("Erro ao enviar mensagem: %v", err)
		errorMessage.Error = "Erro ao enviar mensagem"
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorMessage)
		return
	}
	log.Printf("Mensagem envidada com sucesso: %d", retorno.MessageID)

}
