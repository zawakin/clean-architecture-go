package users

// UserID is user id.
type UserID string

// User is a domain model of an user.
type User struct {
	ID   UserID
	Name string
}

// NewUser returns new user.
func NewUser(name string) *User {
	id := "generated-id-" + name
	return &User{
		ID:   UserID(id),
		Name: name,
	}
}
