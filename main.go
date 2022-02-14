package main

import (
	"Restobook/configs"
	"Restobook/delivery/controllers/auth"
	"Restobook/delivery/controllers/ratings"
	"Restobook/delivery/controllers/restaurants"
	"Restobook/delivery/controllers/topup"
	"Restobook/delivery/controllers/transactions"
	"Restobook/delivery/controllers/users"
	"Restobook/delivery/routes"
	ratingRepo "Restobook/repository/ratings"
	restaurantRepo "Restobook/repository/restaurants"
	topupRepo "Restobook/repository/topup"
	transactionRepo "Restobook/repository/transactions"
	usersRepo "Restobook/repository/users"

	"Restobook/utils"
	"fmt"

	"github.com/labstack/echo/v4"
)

func main() {

	config := configs.GetConfig()
	db := utils.InitDB(config)

	// xendit.Opt.SecretKey = common.XENDIT_SECRET_KEY

	e := echo.New()
	usersRepo := usersRepo.NewUsersRepo(db)
	usersCtrl := users.NewUsersControllers(usersRepo)
	adminCtrl := auth.NewAdminControllers(usersRepo)

	restaurantRepo := restaurantRepo.NewRestaurantsRepo(db)
	restaurantsCtrl := restaurants.NewRestaurantsControllers(restaurantRepo)

	topupRepo := topupRepo.NewTopUpRepo(db)
	topupCtrl := topup.NewTopUpControllers(topupRepo)

	transactionRepo := transactionRepo.NewTransactionRepo(db)
	transactionCtrl := transactions.NewTransactionsControllers(transactionRepo)

	ratingRepo := ratingRepo.NewRatingsRepo(db)
	ratingCtrl := ratings.NewRatingController(ratingRepo)

	routes.RegisterPath(e, adminCtrl, usersCtrl, restaurantsCtrl, transactionCtrl, topupCtrl, ratingCtrl)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", config.Port)))
}
