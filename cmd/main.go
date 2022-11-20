package main

import (
	"log"

	"github.com/sabmile/todo-app-by-zhashkevych"
	"github.com/sabmile/todo-app-by-zhashkevych/pkg/handler"
	"github.com/sabmile/todo-app-by-zhashkevych/pkg/repository"
	"github.com/sabmile/todo-app-by-zhashkevych/pkg/service"
)

func main() {
	// handlers := new(handler.Handler)
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occurred while running http server: %s", err.Error())
	}
}
