package user

import "github.com/rhperera/go-base/domain"

type service struct {
	userRepo domain.UserRepo
}

func (s *service) GetByID(id int64) *domain.User {
	panic("implement me")
}

func NewService(uR domain.UserRepo) domain.UserService {
	return &service{userRepo: uR}
}