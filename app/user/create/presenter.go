package create

import "github.com/zawawahoge/ddd-example/usecase/user/create"

type userCreatePresenter struct {
	subject *UserCreateSubject
}

// NewUserCreatePresenter is constructor of user create presenter.
func NewUserCreatePresenter(subject *UserCreateSubject) create.UserCreatePresenter {
	return &userCreatePresenter{
		subject: subject,
	}
}

func (p *userCreatePresenter) Complete(outputData *create.UserCreateOutputData) {
	viewModel := &UserCreateViewModel{
		UserID:      string(outputData.UserID),
		CreatedDate: outputData.Created.Local().String(),
	}
	p.subject.Update(viewModel)
}
