package routes

import (
	"Restobook/delivery/common"
	"Restobook/delivery/controllers/auth"
	"Restobook/delivery/controllers/ratings"
	"Restobook/delivery/controllers/restaurants"
	"Restobook/delivery/controllers/topup"
	"Restobook/delivery/controllers/transactions"
	"Restobook/delivery/controllers/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterPath(e *echo.Echo, adctrl *auth.AdminController, uctrl *users.UsersController, rctrl *restaurants.RestaurantsController, tctrl *transactions.TransactionsController, tpctrl *topup.TopUpController, rtctrl *ratings.RatingsController) {

	// ---------------------------------------------------------------------
	// CRUD Admin
	// ---------------------------------------------------------------------
	e.POST("/admin/register", adctrl.RegisterAdminCtrl())
	e.POST("/admin/login", adctrl.LoginAdminCtrl())
	e.GET("/admin/waiting", rctrl.GetsWaiting(), middleware.JWT([]byte(common.JWT_SECRET_KEY)))
	e.POST("/admin/approve", rctrl.Approve(), middleware.JWT([]byte(common.JWT_SECRET_KEY)))

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

	e.GET("/myrestaurant", rctrl.GetMyRestoCtrl(), middleware.JWT([]byte(common.JWT_SECRET_KEY)))
	e.PUT("/myrestaurant", rctrl.UpdateMyRestoCtrl(), middleware.JWT([]byte(common.JWT_SECRET_KEY)))

	e.POST("/myrestaurant/detail", rctrl.CreateDetailRestoCtrl(), middleware.JWT([]byte(common.JWT_SECRET_KEY)))
	e.PUT("/myrestaurant/detail", rctrl.UpdateDetailRestoCtrl(), middleware.JWT([]byte(common.JWT_SECRET_KEY)))

	e.DELETE("/myrestaurant", rctrl.DeleteRestoCtrl(), middleware.JWT([]byte(common.JWT_SECRET_KEY)))

	e.GET("/restaurants", rctrl.Gets())
	e.GET("/restaurants/open", rctrl.GetsByOpen())

	// ---------------------------------------------------------------------
	// CRUD Transactions
	// ---------------------------------------------------------------------
	e.POST("/transaction", tctrl.CreateTransactionCtrl(), middleware.JWT(([]byte(common.JWT_SECRET_KEY))))
	e.GET("/transaction/waiting", tctrl.GetAllWaitingCtrl(), middleware.JWT(([]byte(common.JWT_SECRET_KEY))))
	e.GET("/transaction/restaurant/waiting", tctrl.GetAllWaitingForRestoCtrl(), middleware.JWT(([]byte(common.JWT_SECRET_KEY))))
	e.GET("/transaction/restaurant/accepted", tctrl.GetAllAcceptedForRestoCtrl(), middleware.JWT(([]byte(common.JWT_SECRET_KEY))))
	e.GET("/transaction/accepted", tctrl.GetAllAcceptedCtrl(), middleware.JWT(([]byte(common.JWT_SECRET_KEY))))
	e.GET("/transaction/history", tctrl.GetHistoryCtrl(), middleware.JWT(([]byte(common.JWT_SECRET_KEY))))
	e.PUT("/transaction/accepted", tctrl.AcceptTransactionCtrl(), middleware.JWT(([]byte(common.JWT_SECRET_KEY))))
	e.PUT("/transaction/rejected", tctrl.RejectTransactionCtrl(), middleware.JWT(([]byte(common.JWT_SECRET_KEY))))
	e.PUT("/transaction/success", tctrl.SuccessTransactionCtrl(), middleware.JWT(([]byte(common.JWT_SECRET_KEY))))
	e.PUT("/transaction/fail", tctrl.FailTransactionCtrl(), middleware.JWT(([]byte(common.JWT_SECRET_KEY))))
	e.PUT("/transaction/cancel", tctrl.CancelTransactionCtrl(), middleware.JWT(([]byte(common.JWT_SECRET_KEY))))

	// ---------------------------------------------------------------------
	// CRUD TopUp
	// ---------------------------------------------------------------------
	e.POST("/topup", tpctrl.TopUp(), middleware.JWT(([]byte(common.JWT_SECRET_KEY))))
	e.GET("/topup/pending", tpctrl.GetAllWaiting(), middleware.JWT(([]byte(common.JWT_SECRET_KEY))))
	e.GET("/topup/history", tpctrl.GetAllPaid(), middleware.JWT(([]byte(common.JWT_SECRET_KEY))))
	e.POST("/topup/callback", tpctrl.Callback())

	// ---------------------------------------------------------------------
	// CRUD Rating
	// ---------------------------------------------------------------------
	e.POST("/ratings", rtctrl.Create(), middleware.JWT(([]byte(common.JWT_SECRET_KEY))))
	e.PUT("/ratings/:restaurantId", rtctrl.Update(), middleware.JWT(([]byte(common.JWT_SECRET_KEY))))
	e.DELETE("/ratings/:restaurantId", rtctrl.Delete(), middleware.JWT(([]byte(common.JWT_SECRET_KEY))))

}
