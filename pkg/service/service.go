package service

import "github.com/xrodazxx/App-ToDO/pkg/repository"

type Authorization interface {
}
type TodoList interface {
}
type TodoItem interface {
}
type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: repos.Authorization, // внедрение зависимости
		TodoList:      repos.TodoList,      // внедрение зависимости
		TodoItem:      repos.TodoItem,      // внедрение зависимости
	}
}
