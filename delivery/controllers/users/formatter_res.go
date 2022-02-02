package users

import "Restobook/entities"

type GetUsersResponseFormat struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
type LoginResponseFormat struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Token   string `json:"token"`
}
type UserResponseFormat struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
type TestingUserResponseFormat struct {
	Code    int           `json:"code"`
	Message string        `json:"message"`
	Data    entities.User `json:"data"`
}
type UserResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
