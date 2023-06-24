package service

import "gin_sample/db"

type UserService interface {
	GetUser(name string) (string, bool)
	UpdateUser(name string, value string)
}

func NewUserService() UserService {
	return &UserServiceImpl{
		db: db.DB,
	}
}

type UserServiceImpl struct {
	db map[string]string
}

func (s UserServiceImpl) GetUser(name string) (string, bool) {
	value, ok := s.db[name]
	return value, ok
}

func (s UserServiceImpl) UpdateUser(name string, value string) {
	s.db[name] = value
}
