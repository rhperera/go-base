package user

import "github.com/rhperera/go-base/domain"

type handler struct {
	userService domain.UserService
}

func NewHandler(us domain.UserService) *handler {
	return &handler{userService: us}
}