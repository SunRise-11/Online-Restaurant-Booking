package routes

import (
	"Restobook/delivery/common"
	"Restobook/delivery/controllers/restaurants"
	"Restobook/delivery/controllers/transactions"
	"Restobook/delivery/controllers/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterPath(e *echo.Echo, uctrl *users.UsersController, rctrl *restaurants.RestaurantsController, tctrl *transactions.TransactionsController) {
	// ---------------------------------------------------------------------
	// CRUD Users
	// ---------------------------------------------------------------------
	e.POST("/users/register", uctrl.RegisterUserCtrl())
	e.POST("/users/login", uctrl.LoginAuthCtrl())
	e.GET("/user", uctrl.GetUserByIdCtrl(), middleware.JWT([]byte(common.JWT_SECRET_KEY)))
	e.PUT("/user", uctrl.UpdateUserCtrl(), middleware.JWT([]byte(common.JWT_SECRET_KEY)))
	e.DELETE("/user", uctrl.DeleteUserCtrl(), middleware.JWT([]byte(common.JWT_SECRET_KEY)))

	// ---------------------------------------------------------------------
	// CRUD Restaurants
	// ---------------------------------------------------------------------
	e.POST("/restaurants/register", rctrl.RegisterUserCtrl())
	e.POST("/restaurants/login", rctrl.LoginAuthCtrl())

	// ---------------------------------------------------------------------
	// CRUD Transactions
	// ---------------------------------------------------------------------
	e.POST("/transaction", tctrl.CreateTransactionCtrl(), middleware.JWT(([]byte(common.JWT_SECRET_KEY))))
	e.GET("/transaction/waiting", tctrl.GetAllWaitingCtrl(), middleware.JWT(([]byte(common.JWT_SECRET_KEY))))
	e.GET("/transaction/accepted", tctrl.GetAllWaitingCtrl(), middleware.JWT(([]byte(common.JWT_SECRET_KEY))))
	e.GET("/transaction/history", tctrl.GetAllWaitingCtrl(), middleware.JWT(([]byte(common.JWT_SECRET_KEY))))
}
