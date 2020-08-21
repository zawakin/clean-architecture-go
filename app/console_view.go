package main

import (
	"fmt"

	"github.com/zawawahoge/ddd-example/app/user/create"
)

// ConsoleView is console application.
type ConsoleView struct {
	subject *create.UserCreateSubject
}

// NewConsoleView returns new console app.
func NewConsoleView(subject *create.UserCreateSubject) *ConsoleView {
	return &ConsoleView{
		subject: subject,
	}
}

// Update updates view of console app.
func (a *ConsoleView) Update(viewModel *create.UserCreateViewModel) {
	fmt.Printf("id = %s\tcreated = %s\n", viewModel.UserID, viewModel.CreatedDate)
}
