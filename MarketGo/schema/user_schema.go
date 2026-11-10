package schema

type UserSchema struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Username string `json:"username,omitempty"`
}
