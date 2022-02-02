package users

type LoginRequestFormat struct {
	Email    string `json:"email" form:"password"`
	Password string `json:"password" form:"password"`
}
type UserRequestFormat struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Name     string `json:"name" form:"name"`
}
