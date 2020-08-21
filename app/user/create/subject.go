package create

// NewUserCreateSubject is contructor of subject of user create.
func NewUserCreateSubject() *UserCreateSubject {
	return &UserCreateSubject{
		ViewModel: make(chan *UserCreateViewModel),
	}
}

// UserCreateSubject is event emitter of user create.
type UserCreateSubject struct {
	ViewModel chan *UserCreateViewModel
}

// Update updates view model of event emitter.
func (s *UserCreateSubject) Update(viewModel *UserCreateViewModel) {
	go func() {
		s.ViewModel <- viewModel
	}()
}
