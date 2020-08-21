package users

import "github.com/zawawahoge/ddd-example/domain/domain/users"

type memoryUserRepository struct {
	users []*users.User
}

// NewMemoryUserRepository returns memory user repository implementation.
func NewMemoryUserRepository() users.UserRepository {
	return &memoryUserRepository{
		users: make([]*users.User, 0),
	}
}

func (r *memoryUserRepository) Save(user *users.User) error {
	r.users = append(r.users, user)
	return nil
}

func (r *memoryUserRepository) FindAll() ([]*users.User, error) {
	return r.users, nil
}
