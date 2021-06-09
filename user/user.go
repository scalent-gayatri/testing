package user

import (
	"fmt"

	"github.com/scalent-gayatri/testing/doer"
)

type User struct {
	Doer doer.Doer
}

func (u *User) Use() error {
	fmt.Println("Inside Use() ******")
	return u.Doer.DoSomething(123, "Hello GoMock")
}
