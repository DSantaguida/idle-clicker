package models

type Bank struct {
	Id    string
	Value int
}

func CreateBank(id string, value int) *Bank {
	return &Bank{Id: id, Value: value}
}