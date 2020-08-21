package user

import "github.com/zawawahoge/ddd-example/usecase/user/create"

type UserController struct {
	userCreateUsecase create.UserCreateUsecase
}

func NewUserController(userCreateUsecase create.UserCreateUsecase) *UserController{
	return &UserController{
		userCreateUsecase: userCreateUsecase,
	}
}

func (c *UserController) CreateUser(userName string) {
	inputData := &create.UserCreateInputData{
		UserName: userName,
	}
	c.userCreateUsecase.Handle(inputData)
}