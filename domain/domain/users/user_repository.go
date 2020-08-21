package users

// UserRepository is repository of user.
type UserRepository interface {
	Save(user *User) error
	FindAll() ([]*User, error)
}
