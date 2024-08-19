package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/dsantaguida/idle-clicker/services/gateway/internal/api/model"
	"github.com/dsantaguida/idle-clicker/services/gateway/internal/client"
)

type ApiHandler struct {
	client client.IdleClient
}

func CreateApiHandler(client *client.IdleClient) *ApiHandler {
	return &ApiHandler{client: *client}
}

func (handler *ApiHandler) Register(writer http.ResponseWriter, request *http.Request) {
	log.Print("Register request")
	ctx := request.Context()
	user := &model.User{}
	if err := json.NewDecoder(request.Body).Decode(user); err != nil {
		http.Error(writer, err.Error(), 500)
		return
	}

	err := handler.client.Register(ctx, user.Username, user.Password)
	if err != nil {
		http.Error(writer, err.Error(), 500)
		return
	}

	writer.WriteHeader(http.StatusOK)
}

func (handler *ApiHandler) Login(writer http.ResponseWriter, request *http.Request) {
	log.Print("Login request")
	ctx := request.Context()
	user := &model.User{}
	if err := json.NewDecoder(request.Body).Decode(user); err != nil {
		http.Error(writer, err.Error(), 500)
		return
	}

	token, value, err := handler.client.Login(ctx, user.Username, user.Password)
	if err != nil {
		http.Error(writer, err.Error(), 500)
		writer.Write([]byte(err.Error()))
		return
	}

	writer.WriteHeader(http.StatusOK)

	resp := &model.LoginResponse{
		Token: token,
		Value: int64(value),
	}
	err = json.NewEncoder(writer).Encode(resp)
	if err != nil {
		http.Error(writer, err.Error(), 500)
	}
}

func (handler *ApiHandler) UpdateBankValue(writer http.ResponseWriter, request *http.Request) {
	log.Print("Update bank request")
	ctx := request.Context()
	bankUpdate := &model.BankUpdate{}
	if err := json.NewDecoder(request.Body).Decode(bankUpdate); err != nil {
		http.Error(writer, err.Error(), 500)
		return
	}

	err := handler.client.UpdateBankValue(ctx, bankUpdate.Token, bankUpdate.Value)
	if err != nil {
		http.Error(writer, err.Error(), 500)
		writer.Write([]byte(err.Error()))
		return
	}

	writer.WriteHeader(http.StatusOK)
}