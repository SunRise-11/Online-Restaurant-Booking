package routes

import (
	"Restobook/delivery/controllers/users"

	"github.com/labstack/echo/v4"
)

func RegisterPath(e *echo.Echo, uctrl *users.UsersController) {
	// ---------------------------------------------------------------------
	// CRUD Users
	// ---------------------------------------------------------------------
	e.POST("/users/register", uctrl.RegisterUserCtrl())
	e.POST("/users/login", uctrl.LoginAuthCtrl())
}
