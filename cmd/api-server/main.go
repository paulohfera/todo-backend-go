package main

import (
	"github.com/paulohfera/todo-backend-go/internal"
	"github.com/paulohfera/todo-backend-go/pkg/api"
)

func main() {
	api := internal.RegisterServices()
	shutdown(api)
}

func shutdown(api *api.Api) {
}
