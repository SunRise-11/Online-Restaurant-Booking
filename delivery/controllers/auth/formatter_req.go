package auth

type LoginRequestFormat struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type RegisterRequestFormat struct {
	Name         string `json:"name"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	Phone_Number string `json:"phone_number"`
}
