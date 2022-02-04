package routes

import (
	"Restobook/delivery/common"
	"Restobook/delivery/controllers/restaurants"
	"Restobook/delivery/controllers/topup"
	"Restobook/delivery/controllers/transactions"
	"Restobook/delivery/controllers/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterPath(e *echo.Echo, uctrl *users.UsersController, rctrl *restaurants.RestaurantsController, tctrl *transactions.TransactionsController, tpctrl *topup.TopUpController) {

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
	e.POST("/restaurants/register", rctrl.RegisterRestoCtrl())
	e.POST("/restaurants/login", rctrl.LoginRestoCtrl())
	e.GET("/restaurant", rctrl.GetRestoByIdCtrl(), middleware.JWT([]byte(common.JWT_SECRET_KEY)))
	e.PUT("/restaurant", rctrl.UpdateRestoByIdCtrl(), middleware.JWT([]byte(common.JWT_SECRET_KEY)))
	e.POST("/restaurant/detail", rctrl.CreateDetailRestoByIdCtrl(), middleware.JWT([]byte(common.JWT_SECRET_KEY)))
	e.PUT("/restaurant/detail", rctrl.UpdateDetailRestoByIdCtrl(), middleware.JWT([]byte(common.JWT_SECRET_KEY)))
	// e.DELETE("/restaurant", rctrl.DeleteUserCtrl(), middleware.JWT([]byte(common.JWT_SECRET_KEY)))

	// ---------------------------------------------------------------------
	// CRUD Transactions
	// ---------------------------------------------------------------------
	e.POST("/transaction", tctrl.CreateTransactionCtrl(), middleware.JWT(([]byte(common.JWT_SECRET_KEY))))
	e.GET("/transaction/waiting", tctrl.GetAllWaitingCtrl(), middleware.JWT(([]byte(common.JWT_SECRET_KEY))))
	e.GET("/transaction/accepted", tctrl.GetAllAcceptedCtrl(), middleware.JWT(([]byte(common.JWT_SECRET_KEY))))
	e.GET("/transaction/history", tctrl.GetHistoryCtrl(), middleware.JWT(([]byte(common.JWT_SECRET_KEY))))

	// CRUD TopUp
	// ---------------------------------------------------------------------
	e.POST("/topup", tpctrl.TopUp(), middleware.JWT(([]byte(common.JWT_SECRET_KEY))))
	e.GET("/topup/pending", tpctrl.GetAllWaiting(), middleware.JWT(([]byte(common.JWT_SECRET_KEY))))
	e.GET("/topup/history", tpctrl.GetAllPaid(), middleware.JWT(([]byte(common.JWT_SECRET_KEY))))
}
