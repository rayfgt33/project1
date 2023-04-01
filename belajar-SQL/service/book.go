package service

import (
	"belajar-SQL/model"
)

type BookService interface {
	CreateBook(in model.Book) (res model.Book, err error)
	GetBookByID(id int64) (res model.Book, err error)
	UpdateBook(in model.Book) (res model.Book, err error)
	DeleteBook(id int64) (err error)
}

func (s *Service) CreateBook(in model.Book) (res model.Book, err error) {
	// call repo
	res, err = s.repo.CreateBook(in)
	if err != nil {
		return res, err
	}

	return res, nil

}

func (s *Service) GetBookByID(id int64) (res model.Book, err error) {
	return s.repo.GetBookByID(id)
}

func (s *Service) UpdateBook(in model.Book) (res model.Book, err error) {
	return s.repo.UpdateBook(in)
}

func (s *Service) DeleteBook(id int64) (err error) {
	return s.repo.DeleteBook(id)
}
