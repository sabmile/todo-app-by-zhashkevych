package service

import "github.com/sabmile/todo-app-by-zhashkevych/pkg/repository"

type Authorization interface {
}

type TodoList interface {
}

type TodoItem interface {
}

type Serivce struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Serivce {
	return &Serivce{}
}
