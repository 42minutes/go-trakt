package trakt

import "fmt"

type AuthMethod interface {
	fmt.Stringer
}

type TokenAuth struct {
	AccessToken string
}

func (t TokenAuth) String() string {
	return fmt.Sprintf("Bearer %s", t.AccessToken)
}
