package create

// UserCreatePresenter is presenter of user create.
type UserCreatePresenter interface {
	Complete(outputData *UserCreateOutputData)
}
