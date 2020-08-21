package create

// UserCreateUsecase is usecase of user create.
type UserCreateUsecase interface {
	Handle(inputData *UserCreateInputData) error
}
