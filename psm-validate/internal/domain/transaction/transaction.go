package transaction

import (
	"time"
)

type Transaction struct {
	AccountId     uint32    `json:"account_id" bson:"account_id"`
	OperationType uint32    `json:"operation_type" bson:"operation_type"`
	Amount        float64   `json:"amount" bson:"amount"`
	CreatedAt     time.Time `json:"created_at" bson:"created_at"`
}

type Input struct {
	AccountId     uint32  `json:"account_id" bson:"account_id"`
	OperationType uint32  `json:"operation_type" bson:"operation_type"`
	Amount        float64 `json:"amount" bson:"amount"`
}
