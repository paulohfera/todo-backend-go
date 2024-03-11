package api

import (
	"net/http"

	"github.com/paulohfera/todo-backend-go/configs"
	"github.com/paulohfera/todo-backend-go/internal/rest"
)

type Api struct {
}

func New(config *configs.Configuration, rest *rest.Rest) *Api {
	mux := http.NewServeMux()
	rest.RegisterRouters(mux)

	http.ListenAndServe(config.Http.Port, mux)
	return &Api{}
}
