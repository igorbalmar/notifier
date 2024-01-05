package main

import (
	"log"
	"net/http"
	telegram "notifier/telegram"
	"os"

	"github.com/gorilla/mux"
)

func main() {

	/*f, err := os.OpenFile("alertmanager.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0744)
	if err != nil {
		log.Fatalf("Erro ao abrir arquivo de log!\n%s", err)
	}

	defer f.Close()
	*/
	log.Println("Starting alertmanager...")

	if os.Getenv("TELEGRAM_BOT_TOKEN") == "" {
		log.Fatal("Telegram token não foi definido!\nVerifique se a varíavel TELEGRAM_BOT_TOKEN foi exportada/definida corretamente!")
	}

	router := mux.NewRouter()
	router.HandleFunc("/telegram", telegram.SendTelegram).Methods("POST")
	//router.HandleFunc("/sendmail", email.SendMail).Methods("POST")
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
	//var recipients []string
	//for _, mail := range os.Args {
	//	recipients = append(recipients, mail)
	//}
	//sendmail.SendMail(recipients, "Alerta de servidor caído!", "mobuss.com.br", "HTTTP Get failed", time.Now().Format(time.UnixDate))
	//slack.SlackNotify("This is a test alert!")
}
