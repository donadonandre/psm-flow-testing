package router

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/rabbitmq/amqp091-go"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"net/http"
	"psm-validate/internal/domain/account"
	"psm-validate/internal/domain/operation_type"
	"psm-validate/internal/domain/transaction"
	"psm-validate/internal/infrastructure/database"
	"psm-validate/internal/infrastructure/messagebroker"
	"time"
)

func checkAccount(referenceId uint32) bool {
	var accountAux account.Account
	_, collection = database.MongoClientGetCollection("account")
	filter := bson.M{"reference_id": referenceId}
	return collection.FindOne(context.Background(), filter).Decode(&accountAux) != nil
}

func SaveTransaction(w http.ResponseWriter, r *http.Request) {
	var transactionInput transaction.Input

	err := json.NewDecoder(r.Body).Decode(&transactionInput)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if checkAccount(transactionInput.AccountId) {
		http.Error(w, "Account not found", http.StatusUnprocessableEntity)
		return
	}

	if !operation_type.ExistsOperation(transactionInput.OperationType) {
		http.Error(w, "Invalid Operation Type", http.StatusUnprocessableEntity)
		return
	}

	transactionDocument := toTransaction(transactionInput)

	_, err = collection.InsertOne(context.Background(), transactionDocument)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	go sendTransactionToRabbitMQ(transactionInput)

	w.WriteHeader(http.StatusCreated)
	response, _ := json.Marshal(transactionDocument)
	_, err = w.Write(response)
	if err != nil {
		return
	}
}

func toTransaction(pi transaction.Input) transaction.Transaction {
	return transaction.Transaction{
		CreatedAt:     time.Now(),
		AccountId:     pi.AccountId,
		OperationType: pi.OperationType,
		Amount:        pi.Amount,
	}
}

func sendTransactionToRabbitMQ(transaction transaction.Input) {
	var body = []byte(fmt.Sprintf(`{ 
		"account_id" : %d, 
		"operation_type" : %d,
		"amount" : %.2f }`, transaction.AccountId, transaction.OperationType, transaction.Amount))

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
		"transaction", // Nome da fila
		true,          // Durável
		false,         // Auto-deletável
		false,         // Exclusiva
		false,         // Sem espera
		nil,           // Argumentos extras
	)
	if err != nil {
		log.Println("Error declaring queue: ", err)
	}

	err = rabbitChannel.QueueBind(
		"transaction",
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
}
