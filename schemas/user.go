package schemas

import "strings"

// TODO: validate string length
// TODO: validate must not have whitespace
type Register struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (r *Register) Prepare() {
	r.Username = strings.TrimSpace(r.Username)
	r.Email = strings.ToLower(strings.TrimSpace(r.Email))
	r.Password = strings.TrimSpace(r.Password)
}
