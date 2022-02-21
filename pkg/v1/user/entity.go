package user

import "strings"

type Entity struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`

	Email string `json:"email"`
}

func (u Entity) Name() string {
	return strings.Join([]string{u.Firstname, u.Lastname}, " ")
}
