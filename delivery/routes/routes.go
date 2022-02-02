package routes

import (
	"Restobook/delivery/controllers/users"

	"github.com/labstack/echo/v4"
)

func RegisterPath(e *echo.Echo, uctrl *users.UsersController) {
	// ---------------------------------------------------------------------
	// CRUD Users
	// ---------------------------------------------------------------------
	e.POST("/register", uctrl.RegisterUserCtrl())
}
