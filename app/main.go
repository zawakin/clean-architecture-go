package main

import (
	"fmt"

	"github.com/zawawahoge/ddd-example/app/user"
	"github.com/zawawahoge/ddd-example/app/user/create"
	userin "github.com/zawawahoge/ddd-example/domain/application/user"
	"github.com/zawawahoge/ddd-example/memoryinfrastructure/users"
)

func main() {
	userCreateSubject := create.NewUserCreateSubject()

	userRepository := users.NewMemoryUserRepository()
	userCreatePresenter := create.NewUserCreatePresenter(userCreateSubject)
	userCreateInteractor := userin.NewUserCreateInteractor(userRepository, userCreatePresenter)
	userController := user.NewUserController(userCreateInteractor)

	view := NewConsoleView(userCreateSubject)

	for {
		var userName string
		fmt.Print("user name: ")
		fmt.Scanln(&userName)
		if userName == "" {
			break
		}
		userController.CreateUser(userName)
		select {
		case vm := <-userCreateSubject.ViewModel:
			view.Update(vm)
		}
	}

	fmt.Println("press any key to exit.")
	fmt.Scanln()
}
