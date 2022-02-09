package users

type LoginRequestFormat struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type UserRequestFormat struct {
	Email        string `json:"email"`
	Password     string `json:"password"`
	Name         string `json:"name"`
	Phone_Number string `json:"phone_number"`
}
