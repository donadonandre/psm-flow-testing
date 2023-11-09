package router

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/rabbitmq/amqp091-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
	"psm-validate/internal/domain/account"
	"psm-validate/internal/infrastructure/database"
	"psm-validate/internal/infrastructure/messagebroker"
)

var _, collection = database.MongoClientGetCollection("account")

func FindByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	db := database.GetConnection()
	accountEntity := account.Entity{}
	db.Table("account").Where("id = ?", id).First(&accountEntity)

	w.Header().Set("Content-Type", "application/json")
	if accountEntity.Id == 0 {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(accountEntity)
	}

}

func GetAccountByDocumentNumber(document string) *mongo.SingleResult {
	filter := bson.M{"document_number": document}
	return collection.FindOne(context.TODO(), filter)
}

func CreateAccount(w http.ResponseWriter, r *http.Request) {
	//	mongoInit()
	// defer database.CloseMongo()

	var accountInput account.Input
	err := json.NewDecoder(r.Body).Decode(&accountInput)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = GetAccountByDocumentNumber(accountInput.DocumentNumber).Decode(&accountInput)
	if err == nil {
		http.Error(w, "Account already exists", http.StatusConflict)
		return
	}

	toAccount := accountInput.ToAccount()
	_, err = collection.InsertOne(context.Background(), toAccount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	go sendAccountToRabbitMQ(toAccount)

	w.WriteHeader(http.StatusCreated)
	response, _ := json.Marshal(toAccount)
	w.Write(response)
}

func sendAccountToRabbitMQ(account account.Account) {
	body := []byte(fmt.Sprintf(`{ "reference_id" : %d, "document_number" : "%s"}`, account.ReferenceId, account.DocumentNumber))

	rabbitChannel := messagebroker.InitRabbitMQ()

	err := rabbitChannel.ExchangeDeclare(
		"creation",
		"topic",
		true,
		false,
		false,
		false,
		nil)
	if err != nil {
		log.Println("Error declaring Exchange: ", err)
	}

	_, err = rabbitChannel.QueueDeclare(
		"account", // Nome da fila
		true,      // Durável
		false,     // Auto-deletável
		false,     // Exclusiva
		false,     // Sem espera
		nil,       // Argumentos extras
	)
	if err != nil {
		log.Println("Error declaring queue: ", err)
	}

	err = rabbitChannel.QueueBind(
		"account",
		"",
		"creation",
		false,
		nil,
	)
	if err != nil {
		log.Println("Error binding queue: ", err)
	}

	err = rabbitChannel.PublishWithContext(
		context.Background(),
		"creation",
		"",
		false,
		false,
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		log.Println("Error sending to RabbitMQ: ", err)
	}

	messagebroker.CloseRabbitMQ()
}
