package router

import (
	"net/http"

	"github.com/dsantaguida/idle-clicker/services/gateway/internal/api/handler"
	"github.com/dsantaguida/idle-clicker/services/gateway/internal/client"
	"github.com/go-chi/chi/v5"
)

func NewRouter(client client.IdleClient) *chi.Mux{
	r := chi.NewRouter()

	r.Route("/v1", func(r chi.Router) {
		apiHandler := handler.CreateApiHandler(&client)

		r.Method(http.MethodPost, "/register", http.HandlerFunc(apiHandler.Register))
		r.Method(http.MethodPost, "/login", http.HandlerFunc(apiHandler.Login))
		r.Method(http.MethodPost, "/update", http.HandlerFunc(apiHandler.UpdateBankValue))
	})

	return r
}