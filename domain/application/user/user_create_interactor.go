package user

import (
	"time"

	"github.com/zawawahoge/ddd-example/domain/domain/users"
	"github.com/zawawahoge/ddd-example/usecase/user/create"
)

type userCreateInteractor struct {
	userRepo  users.UserRepository
	presenter create.UserCreatePresenter
}

// NewUserCreateInteractor is interactor of user create.
func NewUserCreateInteractor(userRepo users.UserRepository, presenter create.UserCreatePresenter) create.UserCreateUsecase {
	return &userCreateInteractor{
		userRepo:  userRepo,
		presenter: presenter,
	}
}

func (in *userCreateInteractor) Handle(inputData *create.UserCreateInputData) error {
	user := users.NewUser(inputData.UserName)
	in.userRepo.Save(user)

	outputData := &create.UserCreateOutputData{
		UserID:  user.ID,
		Created: time.Now(),
	}
	in.presenter.Complete(outputData)
	return nil
}
