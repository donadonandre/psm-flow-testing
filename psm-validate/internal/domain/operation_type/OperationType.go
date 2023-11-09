package operation_type

type OperationType struct {
	Id          uint32 `json:"id"`
	Description string `json:"description"`
}

func validOperationTypes() []OperationType {
	validList := make([]OperationType, 0)
	validList = append(validList, OperationType{Id: 1, Description: "COMPRA A VISTA"})
	validList = append(validList, OperationType{Id: 2, Description: "COMPRA PARCELADA"})
	validList = append(validList, OperationType{Id: 3, Description: "SAQUE"})
	validList = append(validList, OperationType{Id: 4, Description: "PAGAMENTO"})
	return validList
}

func ExistsOperation(id uint32) bool {
	for _, operation := range validOperationTypes() {
		if operation.Id == id {
			return true
		}
	}
	return false
}
