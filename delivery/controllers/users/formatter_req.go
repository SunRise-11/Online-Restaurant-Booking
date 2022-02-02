package users

type LoginRequestFormat struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}
type RegisterRequestFormat struct {
	Email        string `json:"email" form:"email"`
	Password     string `json:"password" form:"password"`
	Name         string `json:"name" form:"name"`
	Phone_Number int    `json:"phone_number" form:"phone_number"`
}
