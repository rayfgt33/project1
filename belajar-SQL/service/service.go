package service

import "belajar-SQL/repository"

type Service struct {
	repo repository.RepoInterface
}

type ServiceInterface interface {
	BookService
}

func NewService(repo repository.RepoInterface) ServiceInterface {
	return &Service{repo: repo}
}
