package create

import (
	"time"

	"github.com/zawawahoge/ddd-example/domain/domain/users"
)

// UserCreateOutputData is output data of user create.
type UserCreateOutputData struct {
	UserID  users.UserID
	Created time.Time
}
