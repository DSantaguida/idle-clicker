package model

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type BankUpdate struct {
	Token string `json:"Token"`
	Value int64  `json:"Value"`
}

type LoginResponse struct {
	Token string `json:"Token"`
	Value int64  `json:"Value"`
}