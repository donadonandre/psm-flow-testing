package account

import (
	"psm-validate/internal/domain/id_generator"
)

type Account struct {
	ReferenceId    uint32 `json:"reference_id" bson:"reference_id"`
	DocumentNumber string `json:"document_number" bson:"document_number"`
}

type Entity struct {
	Id             uint32 `json:"id" gorm:"primary_key"`
	DocumentNumber uint32 `json:"document_number"`
}

type Input struct {
	DocumentNumber string `json:"document_number" bson:"document_number"`
}

func (a Input) ToAccount() Account {
	return Account{
		ReferenceId:    id_generator.GetNextId("account"),
		DocumentNumber: a.DocumentNumber,
	}
}
