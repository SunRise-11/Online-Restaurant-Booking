package transactions

import (
	"Restobook/delivery/common"
	"Restobook/delivery/controllers/restaurants"
	"Restobook/delivery/controllers/users"
	"Restobook/entities"
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var jwtTokenUser, jwtTokenRestaurant, dateCreate string
var countCreate int

func TestCrateTransaction(t *testing.T) {
	e := echo.New()
	t.Run("Login User For Error", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]string{
			"email":    "junius@outlook.my",
			"password": "junius123",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/users/login")

		userCtrl := users.NewUsersControllers(mockUserRepository{})
		userCtrl.LoginAuthCtrl()(context)

		responses := users.LoginResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)
		jwtTokenUser = responses.Token
		assert.Equal(t, 200, responses.Code)
		assert.Equal(t, "Successful Operation", responses.Message)
	})
	t.Run("Error Update User Balance", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]interface{}{
			"restaurant_id": 1,
			"date_time":     "2022-02-08 10:00",
			"person":        1,
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))
		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/transaction/rejected")

		restoCtrl := NewTransactionsControllers(mockTransactionRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(restoCtrl.CreateTransactionCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}
		responses := TransactionResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)
		assert.Equal(t, 500, responses.Code)
		assert.Equal(t, "Internal Server Error", responses.Message)
	})
	t.Run("Error Binding Transaction", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]interface{}{
			"restaurant_id": "asd",
			"date_time":     "2022-02-09 10:00",
			"person":        "asd",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))
		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/transaction")

		restoCtrl := NewTransactionsControllers(mockFalseTransactionRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(restoCtrl.CreateTransactionCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}
		responses := restaurants.LoginResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)
		jwtTokenRestaurant = responses.Token

		assert.Equal(t, 400, responses.Code)
		assert.Equal(t, "Bad Request", responses.Message)
	})
	t.Run("Error Get Balance From User", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]interface{}{
			"restaurant_id": 2,
			"date_time":     "2022-02-09 10:00",
			"person":        5,
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))
		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/transaction")

		restoCtrl := NewTransactionsControllers(mockTransactionRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(restoCtrl.CreateTransactionCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}
		responses := restaurants.LoginResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)
		assert.Equal(t, 500, responses.Code)
		assert.Equal(t, "Internal Server Error", responses.Message)
	})
	t.Run("Login User", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]string{
			"email":    "ilham@outlook.my",
			"password": "ilham123",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/users/login")

		userCtrl := users.NewUsersControllers(mockUserRepository{})
		userCtrl.LoginAuthCtrl()(context)

		responses := users.LoginResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)
		jwtTokenUser = responses.Token
		assert.Equal(t, 200, responses.Code)
		assert.Equal(t, "Successful Operation", responses.Message)
	})
	t.Run("Error Get Resto Detail", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]interface{}{
			"restaurant_id": 2,
			"date_time":     "2022-02-09 10:00",
			"person":        5,
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))
		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/transaction")

		restoCtrl := NewTransactionsControllers(mockTransactionRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(restoCtrl.CreateTransactionCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}
		responses := restaurants.LoginResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)
		assert.Equal(t, 500, responses.Code)
		assert.Equal(t, "Internal Server Error", responses.Message)
	})
	t.Run("Error Restaurant Not Oepn Today", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]interface{}{
			"restaurant_id": 1,
			"date_time":     "2022-02-09 11:00",
			"person":        3,
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))
		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/transaction")

		restoCtrl := NewTransactionsControllers(mockTransactionRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(restoCtrl.CreateTransactionCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}
		responses := restaurants.LoginResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)
		jwtTokenRestaurant = responses.Token

		assert.Equal(t, 500, responses.Code)
		assert.Equal(t, "This Restaurant is Not Open Today", responses.Message)
	})
	t.Run("Error Transaction Before Restaurant Open", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]interface{}{
			"restaurant_id": 1,
			"date_time":     "2022-02-08 07:00",
			"person":        3,
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))
		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/transaction")

		restoCtrl := NewTransactionsControllers(mockTransactionRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(restoCtrl.CreateTransactionCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}
		responses := restaurants.LoginResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)
		jwtTokenRestaurant = responses.Token

		assert.Equal(t, 500, responses.Code)
		assert.Equal(t, "Sorry This Restaurant is Not Open Yet", responses.Message)
	})
	t.Run("Error Transaction After Restaurant Open", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]interface{}{
			"restaurant_id": 1,
			"date_time":     "2022-02-08 17:00",
			"person":        3,
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))
		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/transaction")

		restoCtrl := NewTransactionsControllers(mockTransactionRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(restoCtrl.CreateTransactionCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}
		responses := restaurants.LoginResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)
		jwtTokenRestaurant = responses.Token

		assert.Equal(t, 500, responses.Code)
		assert.Equal(t, "Sorry This Restaurant already closed", responses.Message)
	})
	t.Run("Error Transaction Money Not Enough", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]interface{}{
			"restaurant_id": 1,
			"date_time":     "2022-02-08 10:00",
			"person":        30,
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))
		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/transaction")

		restoCtrl := NewTransactionsControllers(mockTransactionRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(restoCtrl.CreateTransactionCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}
		responses := restaurants.LoginResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)
		jwtTokenRestaurant = responses.Token

		assert.Equal(t, 500, responses.Code)
		assert.Equal(t, "Your Money is Not Enough For Booking This Restaurant", responses.Message)
	})
	t.Run("Success Create Transaction", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]interface{}{
			"restaurant_id": 1,
			"date_time":     "2022-02-08 10:00",
			"person":        1,
		})
		dateCreate = "2022-02-08 10:00"
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))
		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/transaction")

		restoCtrl := NewTransactionsControllers(mockTransactionRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(restoCtrl.CreateTransactionCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}
		responses := restaurants.LoginResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)
		jwtTokenRestaurant = responses.Token

		assert.Equal(t, 200, responses.Code)
		assert.Equal(t, "Successful Operation", responses.Message)
	})
	t.Run("Error Create Transaction", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]interface{}{
			"restaurant_id": 1,
			"date_time":     "2022-02-08 10:00",
			"person":        1,
		})
		dateCreate = "2022-02-08 10:00"
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))
		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/transaction")

		restoCtrl := NewTransactionsControllers(mockFalseTransactionRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(restoCtrl.CreateTransactionCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}
		responses := restaurants.LoginResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)
		jwtTokenRestaurant = responses.Token

		assert.Equal(t, 500, responses.Code)
		assert.Equal(t, "Internal Server Error", responses.Message)
	})
	t.Run("Error Create Transaction at Same Hours", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]interface{}{
			"restaurant_id": 1,
			"date_time":     "2022-02-08 10:00",
			"person":        1,
		})
		dateCreate = "2022-02-08 10:00"
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))
		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/transaction")

		restoCtrl := NewTransactionsControllers(mockTransactionRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(restoCtrl.CreateTransactionCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}
		responses := restaurants.LoginResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)
		jwtTokenRestaurant = responses.Token

		assert.Equal(t, 500, responses.Code)
		assert.Equal(t, "You Already Booked at This Hour", responses.Message)
	})
	t.Run("Error Create Transaction Seat Not Available", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]interface{}{
			"restaurant_id": 1,
			"date_time":     "2022-02-08 11:00",
			"person":        6,
		})
		dateCreate = "2022-02-08 10:00"
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))
		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/transaction")

		restoCtrl := NewTransactionsControllers(mockTransactionRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(restoCtrl.CreateTransactionCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}
		responses := restaurants.LoginResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)
		jwtTokenRestaurant = responses.Token

		assert.Equal(t, 500, responses.Code)
		assert.Equal(t, "Just 4 Seats Available at This Hour", responses.Message)
	})
	t.Run("Login User For Error", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]string{
			"email":    "herlianto@outlook.my",
			"password": "herlianto123",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/users/login")

		userCtrl := users.NewUsersControllers(mockUserRepository{})
		userCtrl.LoginAuthCtrl()(context)

		responses := users.LoginResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)
		jwtTokenUser = responses.Token
		assert.Equal(t, 200, responses.Code)
		assert.Equal(t, "Successful Operation", responses.Message)
	})
	t.Run("Error Create Transaction", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]interface{}{
			"restaurant_id": 1,
			"date_time":     "2022-02-08 10:00",
			"person":        1,
		})
		dateCreate = "2022-02-08 10:00"
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))
		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/transaction")

		restoCtrl := NewTransactionsControllers(mockTransactionRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(restoCtrl.CreateTransactionCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}
		responses := restaurants.LoginResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)
		jwtTokenRestaurant = responses.Token

		assert.Equal(t, 500, responses.Code)
		assert.Equal(t, "Internal Server Error", responses.Message)
	})

}
func TestGetAllWaiting(t *testing.T) {
	e := echo.New()
	t.Run("Login User For Error", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]string{
			"email":    "junius@outlook.my",
			"password": "junius123",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/users/login")

		userCtrl := users.NewUsersControllers(mockUserRepository{})
		userCtrl.LoginAuthCtrl()(context)

		responses := users.LoginResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)
		jwtTokenUser = responses.Token
		assert.Equal(t, 200, responses.Code)
		assert.Equal(t, "Successful Operation", responses.Message)
	})
	t.Run("Error Get All Waiting For User", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))
		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/transaction/waiting")

		restoCtrl := NewTransactionsControllers(mockTransactionRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(restoCtrl.GetAllWaitingCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}
		responses := restaurants.LoginResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)
		jwtTokenRestaurant = responses.Token

		assert.Equal(t, 500, responses.Code)
		assert.Equal(t, "Internal Server Error", responses.Message)
	})
	t.Run("Login User For Error", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]string{
			"email":    "ilham@outlook.my",
			"password": "ilham123",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/users/login")

		userCtrl := users.NewUsersControllers(mockUserRepository{})
		userCtrl.LoginAuthCtrl()(context)

		responses := users.LoginResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)
		jwtTokenUser = responses.Token
		assert.Equal(t, 200, responses.Code)
		assert.Equal(t, "Successful Operation", responses.Message)
	})
	t.Run("Succes Get All Waiting For User", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))
		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/transaction/waiting")

		restoCtrl := NewTransactionsControllers(mockTransactionRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(restoCtrl.GetAllWaitingCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}
		responses := restaurants.LoginResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)
		jwtTokenRestaurant = responses.Token

		assert.Equal(t, 200, responses.Code)
		assert.Equal(t, "Successful Operation", responses.Message)
	})
}

func TestGetAllWaitingForResto(t *testing.T) {
	e := echo.New()
	t.Run("Login Error Restaurant", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]string{
			"email":    "restaurant2@outlook.my",
			"password": "resto123",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/restaurants/login")

		restoCtrl := restaurants.NewRestaurantsControllers(mockRestaurantRepository{})
		restoCtrl.LoginRestoCtrl()(context)

		responses := restaurants.LoginResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)
		jwtTokenRestaurant = responses.Token
		assert.Equal(t, 200, responses.Code)
		assert.Equal(t, "Successful Operation", responses.Message)
	})
	t.Run("Error Get All Waiting For Resto", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenRestaurant))
		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/transaction/restaurant/waiting")

		restoCtrl := NewTransactionsControllers(mockTransactionRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(restoCtrl.GetAllWaitingForRestoCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}
		responses := TransactionResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)
		assert.Equal(t, 500, responses.Code)
		assert.Equal(t, "Internal Server Error", responses.Message)
	})
	t.Run("Login Restaurant", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]string{
			"email":    "restaurant1@outlook.my",
			"password": "resto123",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/restaurants/login")

		restoCtrl := restaurants.NewRestaurantsControllers(mockRestaurantRepository{})
		restoCtrl.LoginRestoCtrl()(context)

		responses := restaurants.LoginResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)
		jwtTokenRestaurant = responses.Token

		assert.Equal(t, 200, responses.Code)
		assert.Equal(t, "Successful Operation", responses.Message)
	})
	t.Run("Succes Get All Waiting For Resto", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenRestaurant))
		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/transaction/restaurant/waiting")

		restoCtrl := NewTransactionsControllers(mockTransactionRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(restoCtrl.GetAllWaitingForRestoCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}
		responses := TransactionResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)
		assert.Equal(t, 200, responses.Code)
		assert.Equal(t, "Successful Operation", responses.Message)
	})
}
func TestGetAllAcceptedForResto(t *testing.T) {
	e := echo.New()
	t.Run("Login Error Restaurant", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]string{
			"email":    "restaurant2@outlook.my",
			"password": "resto123",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/restaurants/login")

		restoCtrl := restaurants.NewRestaurantsControllers(mockRestaurantRepository{})
		restoCtrl.LoginRestoCtrl()(context)

		responses := restaurants.LoginResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)
		jwtTokenRestaurant = responses.Token
		assert.Equal(t, 200, responses.Code)
		assert.Equal(t, "Successful Operation", responses.Message)
	})
	t.Run("Error Get All Accepted For Resto", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenRestaurant))
		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/transaction/restaurant/waiting")

		restoCtrl := NewTransactionsControllers(mockTransactionRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(restoCtrl.GetAllAcceptedForRestoCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}
		responses := TransactionResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)
		assert.Equal(t, 500, responses.Code)
		assert.Equal(t, "Internal Server Error", responses.Message)
	})
	t.Run("Login Restaurant", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]string{
			"email":    "restaurant1@outlook.my",
			"password": "resto123",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/restaurants/login")

		restoCtrl := restaurants.NewRestaurantsControllers(mockRestaurantRepository{})
		restoCtrl.LoginRestoCtrl()(context)

		responses := restaurants.LoginResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)
		jwtTokenRestaurant = responses.Token

		assert.Equal(t, 200, responses.Code)
		assert.Equal(t, "Successful Operation", responses.Message)
	})
	t.Run("Succes Get All Accepted For Resto", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenRestaurant))
		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/transaction/restaurant/waiting")

		restoCtrl := NewTransactionsControllers(mockTransactionRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(restoCtrl.GetAllAcceptedForRestoCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}
		responses := TransactionResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)
		assert.Equal(t, 200, responses.Code)
		assert.Equal(t, "Successful Operation", responses.Message)
	})
}
func TestGetAllHistory(t *testing.T) {
	e := echo.New()
	t.Run("Login User For Error", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]string{
			"email":    "junius@outlook.my",
			"password": "junius123",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/users/login")

		userCtrl := users.NewUsersControllers(mockUserRepository{})
		userCtrl.LoginAuthCtrl()(context)

		responses := users.LoginResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)
		jwtTokenUser = responses.Token
		assert.Equal(t, 200, responses.Code)
		assert.Equal(t, "Successful Operation", responses.Message)
	})
	t.Run("Error Get All History For User", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))
		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/transaction/waiting")

		restoCtrl := NewTransactionsControllers(mockTransactionRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(restoCtrl.GetHistoryCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}
		responses := restaurants.LoginResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)
		jwtTokenRestaurant = responses.Token

		assert.Equal(t, 500, responses.Code)
		assert.Equal(t, "Internal Server Error", responses.Message)
	})
	t.Run("Login User For Error", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]string{
			"email":    "ilham@outlook.my",
			"password": "ilham123",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/users/login")

		userCtrl := users.NewUsersControllers(mockUserRepository{})
		userCtrl.LoginAuthCtrl()(context)

		responses := users.LoginResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)
		jwtTokenUser = responses.Token
		assert.Equal(t, 200, responses.Code)
		assert.Equal(t, "Successful Operation", responses.Message)
	})
	t.Run("Succes Get All History For User", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))
		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/transaction/waiting")

		restoCtrl := NewTransactionsControllers(mockTransactionRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(restoCtrl.GetHistoryCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}
		responses := restaurants.LoginResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)
		jwtTokenRestaurant = responses.Token

		assert.Equal(t, 200, responses.Code)
		assert.Equal(t, "Successful Operation", responses.Message)
	})
}
func TestGetAllAccepted(t *testing.T) {
	e := echo.New()
	t.Run("Login User For Error", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]string{
			"email":    "junius@outlook.my",
			"password": "junius123",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/users/login")

		userCtrl := users.NewUsersControllers(mockUserRepository{})
		userCtrl.LoginAuthCtrl()(context)

		responses := users.LoginResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)
		jwtTokenUser = responses.Token
		assert.Equal(t, 200, responses.Code)
		assert.Equal(t, "Successful Operation", responses.Message)
	})
	t.Run("Error Get All Accepted For User", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))
		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/transaction/waiting")

		restoCtrl := NewTransactionsControllers(mockTransactionRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(restoCtrl.GetAllAcceptedCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}
		responses := restaurants.LoginResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)
		jwtTokenRestaurant = responses.Token

		assert.Equal(t, 500, responses.Code)
		assert.Equal(t, "Internal Server Error", responses.Message)
	})
	t.Run("Login User For Error", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]string{
			"email":    "ilham@outlook.my",
			"password": "ilham123",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/users/login")

		userCtrl := users.NewUsersControllers(mockUserRepository{})
		userCtrl.LoginAuthCtrl()(context)

		responses := users.LoginResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)
		jwtTokenUser = responses.Token
		assert.Equal(t, 200, responses.Code)
		assert.Equal(t, "Successful Operation", responses.Message)
	})
	t.Run("Succes Get All Accepted For User", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))
		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/transaction/waiting")

		restoCtrl := NewTransactionsControllers(mockTransactionRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(restoCtrl.GetAllAcceptedCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}
		responses := restaurants.LoginResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)
		jwtTokenRestaurant = responses.Token

		assert.Equal(t, 200, responses.Code)
		assert.Equal(t, "Successful Operation", responses.Message)
	})
}
func TestCancelTransaction(t *testing.T) {
	e := echo.New()
	t.Run("Login User", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]string{
			"email":    "ilham@outlook.my",
			"password": "ilham123",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/users/login")

		userCtrl := users.NewUsersControllers(mockUserRepository{})
		userCtrl.LoginAuthCtrl()(context)

		responses := users.LoginResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)
		jwtTokenUser = responses.Token
		assert.Equal(t, 200, responses.Code)
		assert.Equal(t, "Successful Operation", responses.Message)
	})
	t.Run("Error Bind Accepted Transaction", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]int{
			"id":     1,
			"status": 1,
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))
		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/transaction/cancel")

		restoCtrl := NewTransactionsControllers(mockFalseTransactionRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(restoCtrl.CancelTransactionCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}
		responses := TransactionResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)
		assert.Equal(t, 400, responses.Code)
		assert.Equal(t, "Bad Request", responses.Message)
	})
	t.Run("Sucess Cancel Transaction", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]interface{}{
			"id":     1,
			"status": "Cancel",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))
		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/transaction/cancel")

		restoCtrl := NewTransactionsControllers(mockTransactionRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(restoCtrl.CancelTransactionCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}
		responses := TransactionResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)
		assert.Equal(t, 200, responses.Code)
		assert.Equal(t, "Successful Operation", responses.Message)
	})

	t.Run("Error Get Transaction By Id", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]interface{}{
			"id":     2,
			"status": "Cancel",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))
		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/transaction/cancel")

		restoCtrl := NewTransactionsControllers(mockFalseTransactionRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(restoCtrl.CancelTransactionCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}
		responses := TransactionResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)
		assert.Equal(t, 500, responses.Code)
		assert.Equal(t, "Internal Server Error", responses.Message)
	})
	t.Run("Error Update Transaction status", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]interface{}{
			"id":     1,
			"status": "waiting for confirmation",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))
		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/transaction/cancel")

		restoCtrl := NewTransactionsControllers(mockTransactionRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(restoCtrl.CancelTransactionCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}
		responses := TransactionResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)
		assert.Equal(t, 500, responses.Code)
		assert.Equal(t, "Internal Server Error", responses.Message)
	})
	t.Run("Cover Balance < 0", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]interface{}{
			"id":     2,
			"status": "Cancel",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))
		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/transaction/cancel")

		restoCtrl := NewTransactionsControllers(mockTransactionRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(restoCtrl.CancelTransactionCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}
		responses := TransactionResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)
		assert.Equal(t, 500, responses.Code)
		assert.Equal(t, "Your Money is Not Enough For Cancel This Transaction", responses.Message)
	})
	t.Run("Login User For Error", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]string{
			"email":    "junius@outlook.my",
			"password": "junius123",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/users/login")

		userCtrl := users.NewUsersControllers(mockUserRepository{})
		userCtrl.LoginAuthCtrl()(context)

		responses := users.LoginResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)
		jwtTokenUser = responses.Token
		assert.Equal(t, 200, responses.Code)
		assert.Equal(t, "Successful Operation", responses.Message)
	})
	t.Run("Error Update User Balance", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]interface{}{
			"id":     1,
			"status": "waiting for confirmation",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))
		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/transaction/cancel")

		restoCtrl := NewTransactionsControllers(mockTransactionRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(restoCtrl.CancelTransactionCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}
		responses := TransactionResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)
		assert.Equal(t, 500, responses.Code)
		assert.Equal(t, "Internal Server Error", responses.Message)
	})
	t.Run("Login User For Error", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]string{
			"email":    "herlianto@outlook.my",
			"password": "herlianto123",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/users/login")

		userCtrl := users.NewUsersControllers(mockUserRepository{})
		userCtrl.LoginAuthCtrl()(context)

		responses := users.LoginResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)
		jwtTokenUser = responses.Token
		assert.Equal(t, 200, responses.Code)
		assert.Equal(t, "Successful Operation", responses.Message)
	})
	t.Run("Error Update User Reputation", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]interface{}{
			"id":     1,
			"status": "waiting for confirmation",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))
		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/transaction/sucess")

		restoCtrl := NewTransactionsControllers(mockTransactionRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(restoCtrl.CancelTransactionCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}
		responses := TransactionResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)
		assert.Equal(t, 500, responses.Code)
		assert.Equal(t, "Internal Server Error", responses.Message)
	})
	t.Run("Login User For Error", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]string{
			"email":    "andrew@outlook.my",
			"password": "andrew123",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/users/login")

		userCtrl := users.NewUsersControllers(mockUserRepository{})
		userCtrl.LoginAuthCtrl()(context)

		responses := users.LoginResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)
		jwtTokenUser = responses.Token
		assert.Equal(t, 200, responses.Code)
		assert.Equal(t, "Successful Operation", responses.Message)
	})
	t.Run("Error Update Compare", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]interface{}{
			"id":     1,
			"status": "waiting for confirmation",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))
		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/transaction/cancel")

		restoCtrl := NewTransactionsControllers(mockTransactionRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(restoCtrl.CancelTransactionCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}
		responses := TransactionResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)
		assert.Equal(t, 500, responses.Code)
		assert.Equal(t, "You Cant Cancel This Transaction at This Time", responses.Message)
	})
}
func TestAcceptTransaction(t *testing.T) {
	e := echo.New()
	t.Run("Login Restaurant", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]string{
			"email":    "restaurant1@outlook.my",
			"password": "resto123",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/restaurants/login")

		restoCtrl := restaurants.NewRestaurantsControllers(mockRestaurantRepository{})
		restoCtrl.LoginRestoCtrl()(context)

		responses := restaurants.LoginResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)
		jwtTokenRestaurant = responses.Token

		assert.Equal(t, 200, responses.Code)
		assert.Equal(t, "Successful Operation", responses.Message)
	})
	t.Run("Success Accepted Transaction", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]interface{}{
			"id":     1,
			"status": "Accepted",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenRestaurant))
		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/transaction")

		restoCtrl := NewTransactionsControllers(mockTransactionRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(restoCtrl.AcceptTransactionCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}
		responses := TransactionResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)
		assert.Equal(t, 200, responses.Code)
		assert.Equal(t, "Successful Operation", responses.Message)
	})
	t.Run("Error Bind Accepted Transaction", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]int{
			"id":     1,
			"status": 1,
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenRestaurant))
		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/transaction")

		restoCtrl := NewTransactionsControllers(mockFalseTransactionRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(restoCtrl.AcceptTransactionCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}
		responses := TransactionResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)
		assert.Equal(t, 400, responses.Code)
		assert.Equal(t, "Bad Request", responses.Message)
	})
	t.Run("Error Get Transaction By Status Transaction", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]interface{}{
			"id":     2,
			"status": "Accepted",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenRestaurant))
		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/transaction/accepted")

		restoCtrl := NewTransactionsControllers(mockTransactionRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(restoCtrl.AcceptTransactionCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}
		responses := TransactionResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)
		assert.Equal(t, 500, responses.Code)
		assert.Equal(t, "Internal Server Error", responses.Message)
	})
	t.Run("Error Update Transaction By Status", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]interface{}{
			"id":     1,
			"status": "waiting for confirmation",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenRestaurant))
		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/transaction/accepted")

		restoCtrl := NewTransactionsControllers(mockTransactionRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(restoCtrl.AcceptTransactionCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}
		responses := TransactionResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)
		assert.Equal(t, 500, responses.Code)
		assert.Equal(t, "Internal Server Error", responses.Message)
	})
}
func TestRejectTransaction(t *testing.T) {
	e := echo.New()
	t.Run("Login Restaurant", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]string{
			"email":    "restaurant1@outlook.my",
			"password": "resto123",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/restaurants/login")

		restoCtrl := restaurants.NewRestaurantsControllers(mockRestaurantRepository{})
		restoCtrl.LoginRestoCtrl()(context)

		responses := restaurants.LoginResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)
		jwtTokenRestaurant = responses.Token

		assert.Equal(t, 200, responses.Code)
		assert.Equal(t, "Successful Operation", responses.Message)
	})
	t.Run("Success Rejected Transaction", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]interface{}{
			"id":     1,
			"status": "Rejected",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenRestaurant))
		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/transaction/rejected")

		restoCtrl := NewTransactionsControllers(mockTransactionRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(restoCtrl.RejectTransactionCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}
		responses := TransactionResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)
		assert.Equal(t, 200, responses.Code)
		assert.Equal(t, "Successful Operation", responses.Message)
	})
	t.Run("Error Binding Rejected Transaction", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]int{
			"id":     1,
			"status": 1,
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenRestaurant))
		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/transaction/rejected")

		restoCtrl := NewTransactionsControllers(mockFalseTransactionRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(restoCtrl.RejectTransactionCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}
		responses := TransactionResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)
		assert.Equal(t, 400, responses.Code)
		assert.Equal(t, "Bad Request", responses.Message)
	})
	t.Run("Error Get Transaction By Status", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]interface{}{
			"id":     2,
			"status": "waiting for confirmation",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenRestaurant))
		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/transaction/rejected")

		restoCtrl := NewTransactionsControllers(mockTransactionRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(restoCtrl.RejectTransactionCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}
		responses := TransactionResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)
		assert.Equal(t, 500, responses.Code)
		assert.Equal(t, "Internal Server Error", responses.Message)
	})
	t.Run("Error Update User Balance", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]interface{}{
			"id":     3,
			"status": "waiting for confirmation",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenRestaurant))
		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/transaction/rejected")

		restoCtrl := NewTransactionsControllers(mockTransactionRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(restoCtrl.RejectTransactionCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}
		responses := TransactionResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)
		assert.Equal(t, 500, responses.Code)
		assert.Equal(t, "Internal Server Error", responses.Message)
	})
	t.Run("Error Update Transaction status", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]interface{}{
			"id":     1,
			"status": "waiting for confirmation",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenRestaurant))
		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/transaction/rejected")

		restoCtrl := NewTransactionsControllers(mockTransactionRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(restoCtrl.RejectTransactionCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}
		responses := TransactionResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)
		assert.Equal(t, 500, responses.Code)
		assert.Equal(t, "Internal Server Error", responses.Message)
	})
	t.Run("Error Get Transaction By Status Transaction", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]interface{}{
			"id":     2,
			"status": "Fail",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenRestaurant))
		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/transaction/accepted")

		restoCtrl := NewTransactionsControllers(mockTransactionRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(restoCtrl.FailTransactionCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}
		responses := TransactionResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)
		assert.Equal(t, 500, responses.Code)
		assert.Equal(t, "Internal Server Error", responses.Message)
	})

}
func TestSucessTransaction(t *testing.T) {
	e := echo.New()
	t.Run("Login Restaurant", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]string{
			"email":    "restaurant1@outlook.my",
			"password": "resto123",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/restaurants/login")

		restoCtrl := restaurants.NewRestaurantsControllers(mockRestaurantRepository{})
		restoCtrl.LoginRestoCtrl()(context)

		responses := restaurants.LoginResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)
		jwtTokenRestaurant = responses.Token

		assert.Equal(t, 200, responses.Code)
		assert.Equal(t, "Successful Operation", responses.Message)
	})
	t.Run("Error Binding Success Transaction Status", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]int{
			"id":     1,
			"status": 1,
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenRestaurant))
		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/transaction/sucess")

		restoCtrl := NewTransactionsControllers(mockFalseTransactionRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(restoCtrl.SuccessTransactionCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}
		responses := TransactionResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)
		assert.Equal(t, 400, responses.Code)
		assert.Equal(t, "Bad Request", responses.Message)
	})
	t.Run("Success Transaction Status", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]interface{}{
			"id":      1,
			"user_id": 1,
			"status":  "Success",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenRestaurant))
		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/transaction/sucess")

		restoCtrl := NewTransactionsControllers(mockTransactionRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(restoCtrl.SuccessTransactionCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}
		responses := TransactionResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)
		assert.Equal(t, 200, responses.Code)
		assert.Equal(t, "Successful Operation", responses.Message)
	})
	t.Run("Error Update Transaction Status", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]interface{}{
			"id":      1,
			"user_id": 1,
			"status":  "waiting for confirmation",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenRestaurant))
		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/transaction/sucess")

		restoCtrl := NewTransactionsControllers(mockTransactionRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(restoCtrl.SuccessTransactionCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}
		responses := TransactionResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)
		assert.Equal(t, 500, responses.Code)
		assert.Equal(t, "Internal Server Error", responses.Message)
	})
	t.Run("Error Get Transaction By Status", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]interface{}{
			"id":     2,
			"status": "Success",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenRestaurant))
		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/transaction/sucess")

		restoCtrl := NewTransactionsControllers(mockTransactionRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(restoCtrl.SuccessTransactionCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}
		responses := TransactionResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)
		assert.Equal(t, 500, responses.Code)
		assert.Equal(t, "Internal Server Error", responses.Message)
	})
	t.Run("Error Update User Reputation", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]interface{}{
			"id":      1,
			"user_id": 2,
			"status":  "waiting for confirmation",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenRestaurant))
		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/transaction/sucess")

		restoCtrl := NewTransactionsControllers(mockTransactionRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(restoCtrl.SuccessTransactionCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}
		responses := TransactionResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)
		assert.Equal(t, 500, responses.Code)
		assert.Equal(t, "Internal Server Error", responses.Message)
	})
	t.Run("Error Update Transaction Status", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]interface{}{
			"id":      1,
			"user_id": 1,
			"status":  "waiting for confirmation",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenRestaurant))
		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/transaction/sucess")

		restoCtrl := NewTransactionsControllers(mockTransactionRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(restoCtrl.FailTransactionCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}
		responses := TransactionResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)
		assert.Equal(t, 500, responses.Code)
		assert.Equal(t, "Internal Server Error", responses.Message)
	})
}
func TestFailTransaction(t *testing.T) {
	e := echo.New()
	t.Run("Login Restaurant", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]string{
			"email":    "restaurant1@outlook.my",
			"password": "resto123",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/restaurants/login")

		restoCtrl := restaurants.NewRestaurantsControllers(mockRestaurantRepository{})
		restoCtrl.LoginRestoCtrl()(context)

		responses := restaurants.LoginResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)
		jwtTokenRestaurant = responses.Token

		assert.Equal(t, 200, responses.Code)
		assert.Equal(t, "Successful Operation", responses.Message)
	})
	t.Run("Error Binding Fail Transaction Status", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]int{
			"id":     1,
			"status": 1,
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenRestaurant))
		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/transaction/sucess")

		restoCtrl := NewTransactionsControllers(mockFalseTransactionRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(restoCtrl.FailTransactionCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}
		responses := TransactionResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)
		assert.Equal(t, 400, responses.Code)
		assert.Equal(t, "Bad Request", responses.Message)
	})
	t.Run("Sucess Update Status Fail Transaction", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]interface{}{
			"id":      1,
			"user_id": 1,
			"status":  "Fail",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenRestaurant))
		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/transaction/sucess")

		restoCtrl := NewTransactionsControllers(mockTransactionRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(restoCtrl.FailTransactionCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}
		responses := TransactionResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)
		assert.Equal(t, 200, responses.Code)
		assert.Equal(t, "Successful Operation", responses.Message)
	})
	t.Run("Sucess Update Status Fail Transaction", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]interface{}{
			"id":      3,
			"user_id": 1,
			"status":  "Fail",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenRestaurant))
		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/transaction/sucess")

		restoCtrl := NewTransactionsControllers(mockTransactionRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(restoCtrl.FailTransactionCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}
		responses := TransactionResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)
		assert.Equal(t, 200, responses.Code)
		assert.Equal(t, "Successful Operation", responses.Message)
	})
	t.Run("Error Update User Reputation", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]interface{}{
			"id":      1,
			"user_id": 2,
			"status":  "waiting for confirmation",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenRestaurant))
		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/transaction/sucess")

		restoCtrl := NewTransactionsControllers(mockTransactionRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(restoCtrl.FailTransactionCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}
		responses := TransactionResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)
		assert.Equal(t, 500, responses.Code)
		assert.Equal(t, "Internal Server Error", responses.Message)
	})
}

type mockTransactionRepository struct{}

func (m mockTransactionRepository) Create(newTransaction entities.Transaction) (entities.Transaction, error) {
	countCreate++
	if newTransaction.UserID == 2 || newTransaction.UserID == 3 {
		return entities.Transaction{}, errors.New("")
	}
	return entities.Transaction{
		Model:        gorm.Model{},
		ID:           newTransaction.ID,
		UserID:       newTransaction.UserID,
		RestaurantID: newTransaction.UserID,
		DateTime:     newTransaction.DateTime,
		Persons:      newTransaction.Persons,
		Total:        newTransaction.Total,
		Status:       newTransaction.Status}, nil
}
func (m mockTransactionRepository) GetAllWaiting(userId uint) ([]entities.Transaction, error) {
	if userId == 1 {
		return []entities.Transaction{{
			ID:           1,
			UserID:       userId,
			RestaurantID: 1,
			DateTime:     time.Time{},
			Persons:      1,
			Total:        10000,
			Status:       "waiting for confirmation",
		}}, nil
	} else {
		return []entities.Transaction{{
			ID:           1,
			UserID:       userId,
			RestaurantID: 1,
			DateTime:     time.Time{},
			Persons:      1,
			Total:        10000,
			Status:       "waiting for confirmation",
		}}, errors.New("")
	}

}
func (m mockTransactionRepository) GetAllWaitingForResto(restaurantId uint) ([]entities.Transaction, error) {
	if restaurantId == 1 {
		return []entities.Transaction{{
			ID:           1,
			UserID:       1,
			RestaurantID: restaurantId,
			DateTime:     time.Time{},
			Persons:      1,
			Total:        10000,
			Status:       "waiting for confirmation",
		}}, nil
	} else {
		return []entities.Transaction{{
			ID:           1,
			UserID:       1,
			RestaurantID: restaurantId,
			DateTime:     time.Time{},
			Persons:      1,
			Total:        10000,
			Status:       "waiting for confirmation",
		}}, errors.New("")
	}
}
func (m mockTransactionRepository) GetAllAcceptedForResto(restaurantId uint) ([]entities.Transaction, error) {
	if restaurantId == 1 {
		return []entities.Transaction{{
			ID:           1,
			UserID:       1,
			RestaurantID: restaurantId,
			DateTime:     time.Time{},
			Persons:      1,
			Total:        10000,
			Status:       "Accepted",
		}}, nil
	} else {
		return []entities.Transaction{{
			ID:           1,
			UserID:       1,
			RestaurantID: restaurantId,
			DateTime:     time.Time{},
			Persons:      1,
			Total:        10000,
			Status:       "Accepted",
		}}, errors.New("")
	}
}
func (m mockTransactionRepository) GetHistory(userId uint) ([]entities.Transaction, error) {
	if userId == 1 {
		return []entities.Transaction{{
			ID:           1,
			UserID:       userId,
			RestaurantID: 1,
			DateTime:     time.Time{},
			Persons:      1,
			Total:        10000,
			Status:       "Success",
		}}, nil
	} else {
		return []entities.Transaction{{
			ID:           1,
			UserID:       userId,
			RestaurantID: 1,
			DateTime:     time.Time{},
			Persons:      1,
			Total:        10000,
			Status:       "Success",
		}}, errors.New("")
	}
}
func (m mockTransactionRepository) GetAllAppointed(userId uint) ([]entities.Transaction, error) {
	if userId == 1 {
		return []entities.Transaction{{
			ID:           1,
			UserID:       userId,
			RestaurantID: 1,
			DateTime:     time.Time{},
			Persons:      1,
			Total:        10000,
			Status:       "Accepted",
		}}, nil
	} else {
		return []entities.Transaction{{
			ID:           1,
			UserID:       userId,
			RestaurantID: 1,
			DateTime:     time.Time{},
			Persons:      1,
			Total:        10000,
			Status:       "Accepted",
		}}, errors.New("")
	}
}
func (m mockTransactionRepository) GetBalance(userId uint) (entities.User, error) {
	if userId == 1 {
		return entities.User{ID: 1, Balance: 60000}, nil
	} else if userId == 2 {
		return entities.User{ID: 1, Balance: 60000}, nil
	} else if userId == 3 {
		return entities.User{ID: 1, Balance: 60000}, nil
	} else {
		return entities.User{ID: 0, Balance: 0}, errors.New("")
	}

}
func (m mockTransactionRepository) GetRestoDetail(restaurantId uint) (entities.RestaurantDetail, error) {
	if restaurantId == 1 {
		return entities.RestaurantDetail{ID: 1, Name: "Resto 1", Open_Hour: "08:00", Close_Hour: "13:00", Open: "0,1", Price: 10000, Seats: 5, Status: "OPEN"}, nil
	} else {
		return entities.RestaurantDetail{ID: 2, Status: "CLOSE"}, errors.New("")
	}

}
func (m mockTransactionRepository) UpdateUserBalance(userId uint, balance int) (entities.User, error) {
	if userId == 1 {
		return entities.User{}, nil
	} else if userId == 3 {
		return entities.User{}, nil
	} else {
		return entities.User{}, errors.New("")
	}

}
func (m mockTransactionRepository) UpdateUserReputation(userId uint, reputation int) (entities.User, error) {
	if userId == 1 {
		return entities.User{}, nil
	} else {
		return entities.User{}, errors.New("")
	}

}
func (m mockTransactionRepository) UpdateTransactionStatus(newTransaction entities.Transaction) (entities.Transaction, error) {
	if newTransaction.Status == "Accepted" || newTransaction.Status == "Rejected" || newTransaction.Status == "Fail" || newTransaction.Status == "Cancel" || newTransaction.Status == "Success" {
		return entities.Transaction{ID: 1}, nil
	} else {
		return entities.Transaction{ID: 1}, errors.New("")
	}

}
func (m mockTransactionRepository) GetTransactionById(id, userId uint) (entities.Transaction, error) {
	if id == 2 {
		return entities.Transaction{
			ID:           id,
			UserID:       userId,
			RestaurantID: id,
			Status:       "Accepted",
			User: entities.User{
				ID:         2,
				Reputation: 0,
				Balance:    0,
			},
		}, nil
	} else if userId == 1 {
		return entities.Transaction{
			ID:           id,
			UserID:       userId,
			RestaurantID: id,
			Status:       "Accepted",
			User: entities.User{
				ID:         1,
				Reputation: 0,
				Balance:    100000,
			},
		}, nil
	} else if userId == 2 {
		return entities.Transaction{
			ID:           id,
			UserID:       userId,
			RestaurantID: id,
			Status:       "waiting for confirmation",
			User: entities.User{
				ID:         2,
				Reputation: 0,
				Balance:    100000,
			},
		}, nil
	} else if userId == 3 {
		return entities.Transaction{
			ID:           id,
			UserID:       userId,
			RestaurantID: id,
			Status:       "Accepted",
			User: entities.User{
				ID:         2,
				Reputation: 0,
				Balance:    100000,
			},
		}, nil
	} else if userId == 4 {
		return entities.Transaction{
			ID:           id,
			UserID:       userId,
			RestaurantID: id,
			Status:       "Cancel",
			User: entities.User{
				ID:         2,
				Reputation: 0,
				Balance:    100000,
			},
		}, nil
	} else {
		return entities.Transaction{}, errors.New("")
	}

}
func (m mockTransactionRepository) GetTotalSeat(restaurantId uint, dateTime string) (int, error) {

	return 1, nil
}
func (m mockTransactionRepository) CheckSameHour(restaurantId, userId uint, dateTime string) (bool, error) {
	if countCreate > 0 && userId == 1 && dateTime == dateCreate {
		return true, nil
	}
	return false, nil
}
func (m mockTransactionRepository) GetReputationUser(userId uint) (entities.User, error) {
	return entities.User{}, nil
}
func (m mockTransactionRepository) GetTransactionUserByStatus(id, restaurant_id uint, status string) (entities.Transaction, error) {
	if id == 1 && restaurant_id == 1 {
		return entities.Transaction{
			Model:        gorm.Model{},
			ID:           1,
			UserID:       1,
			RestaurantID: restaurant_id,
			DateTime:     time.Time{},
			Persons:      0,
			Total:        0,
			Status:       "waiting for confirmation",
			User: entities.User{
				Model:      gorm.Model{},
				ID:         1,
				Reputation: 96,
				Balance:    0,
			},
		}, nil
	} else if id == 3 && restaurant_id == 1 {
		return entities.Transaction{ID: 2, UserID: 2, Status: "waiting for confirmation", User: entities.User{
			Reputation: 0,
		}}, nil
	} else {
		return entities.Transaction{}, errors.New("")
	}

}

type mockFalseTransactionRepository struct{}

func (m mockFalseTransactionRepository) Create(newTransaction entities.Transaction) (entities.Transaction, error) {
	return entities.Transaction{
		Model:        gorm.Model{},
		ID:           newTransaction.ID,
		UserID:       newTransaction.UserID,
		RestaurantID: newTransaction.UserID,
		DateTime:     newTransaction.DateTime,
		Persons:      newTransaction.Persons,
		Total:        newTransaction.Total,
		Status:       newTransaction.Status}, errors.New("")
}
func (m mockFalseTransactionRepository) GetAllWaiting(userId uint) ([]entities.Transaction, error) {
	return []entities.Transaction{}, errors.New("")
}
func (m mockFalseTransactionRepository) GetAllWaitingForResto(restaurantId uint) ([]entities.Transaction, error) {
	return []entities.Transaction{}, errors.New("")
}
func (m mockFalseTransactionRepository) GetAllAcceptedForResto(restaurantId uint) ([]entities.Transaction, error) {
	return []entities.Transaction{}, errors.New("")
}
func (m mockFalseTransactionRepository) GetHistory(userId uint) ([]entities.Transaction, error) {
	return []entities.Transaction{}, errors.New("")
}
func (m mockFalseTransactionRepository) GetAllAppointed(userId uint) ([]entities.Transaction, error) {
	return []entities.Transaction{}, errors.New("")
}
func (m mockFalseTransactionRepository) GetBalance(userId uint) (entities.User, error) {
	return entities.User{ID: 1, Balance: 1000000}, errors.New("")
}
func (m mockFalseTransactionRepository) GetRestoDetail(restaurantId uint) (entities.RestaurantDetail, error) {
	return entities.RestaurantDetail{}, errors.New("")
}
func (m mockFalseTransactionRepository) UpdateUserBalance(userId uint, balance int) (entities.User, error) {
	return entities.User{}, errors.New("")
}
func (m mockFalseTransactionRepository) UpdateUserReputation(userId uint, reputation int) (entities.User, error) {
	return entities.User{}, errors.New("")
}
func (m mockFalseTransactionRepository) UpdateTransactionStatus(newTransaction entities.Transaction) (entities.Transaction, error) {
	return entities.Transaction{}, errors.New("")
}

func (m mockFalseTransactionRepository) GetTransactionById(id, userId uint) (entities.Transaction, error) {
	return entities.Transaction{}, errors.New("")
}
func (m mockFalseTransactionRepository) GetTotalSeat(restaurantId uint, dateTime string) (int, error) {
	return 1, errors.New("")
}
func (m mockFalseTransactionRepository) CheckSameHour(restaurantId, userId uint, dateTime string) (bool, error) {
	return false, errors.New("")
}
func (m mockFalseTransactionRepository) GetReputationUser(userId uint) (entities.User, error) {
	return entities.User{}, errors.New("")
}
func (m mockFalseTransactionRepository) GetTransactionUserByStatus(id, restaurant_id uint, status string) (entities.Transaction, error) {
	return entities.Transaction{}, errors.New("")
}

type mockUserRepository struct{}

func (m mockUserRepository) RegisterAdmin(newUser entities.User) (entities.User, error) {
	return entities.User{ID: 1, Name: "admin"}, nil
}

func (m mockUserRepository) Register(newUser entities.User) (entities.User, error) {
	hash := sha256.Sum256([]byte("ilham123"))
	passwordS := fmt.Sprintf("%x", hash[:])
	return entities.User{ID: 1, Name: "ilham", Password: passwordS, Email: "ilham@outlook.my", Reputation: 80, Balance: 1000000}, nil
}

func (m mockUserRepository) LoginUser(email, password string) (entities.User, error) {
	hash := sha256.Sum256([]byte(password))
	passwordS := fmt.Sprintf("%x", hash[:])
	if email == "ilham@outlook.my" {
		return entities.User{ID: 1, Name: "ilham", Password: passwordS, Email: email}, nil
	} else if email == "junius@outlook.my" {
		return entities.User{ID: 2, Name: "junius", Password: passwordS, Email: email}, nil
	} else if email == "herlianto@outlook.my" {
		return entities.User{ID: 3, Name: "herlianto", Password: passwordS, Email: email}, nil
	} else {
		return entities.User{ID: 4, Name: "andrew", Password: passwordS, Email: email}, nil
	}

}

func (m mockUserRepository) Get(userID uint) (entities.User, error) {
	return entities.User{ID: 1, Name: "ilham"}, nil
}

func (m mockUserRepository) Update(userID uint, updateUser entities.User) (entities.User, error) {
	return entities.User{ID: 1, Name: "junius"}, nil
}

func (m mockUserRepository) Delete(userID uint) (entities.User, error) {
	return entities.User{ID: 1}, nil
}

type mockRestaurantRepository struct{}

func (m mockRestaurantRepository) Register(newUser entities.Restaurant) (entities.Restaurant, error) {
	return entities.Restaurant{ID: 1, Email: "restaurant1@outlook.my"}, nil
}

func (m mockRestaurantRepository) Login(email, password string) (entities.Restaurant, error) {
	hash := sha256.Sum256([]byte(password))
	passwordS := fmt.Sprintf("%x", hash[:])
	if email == "restaurant1@outlook.my" {
		return entities.Restaurant{
			ID:                 1,
			Email:              email,
			Password:           passwordS,
			RestaurantDetailID: 1}, nil
	} else {
		return entities.Restaurant{
			ID:                 2,
			Email:              email,
			Password:           passwordS,
			RestaurantDetailID: 1}, nil
	}

}

func (m mockRestaurantRepository) GetsWaiting() ([]entities.RestaurantDetail, error) {
	return []entities.RestaurantDetail{{ID: 1}}, nil
}

func (m mockRestaurantRepository) Approve(restaurantID uint, status string) (entities.RestaurantDetail, error) {
	return entities.RestaurantDetail{ID: 1}, nil
}

func (m mockRestaurantRepository) Get(restaurantID uint) (entities.Restaurant, entities.RestaurantDetail, error) {
	return entities.Restaurant{ID: 1}, entities.RestaurantDetail{ID: 1}, nil
}

func (m mockRestaurantRepository) GetsByOpen(open int) ([]entities.RestaurantDetail, error) {
	return []entities.RestaurantDetail{{}}, nil
}

func (m mockRestaurantRepository) GetExistSeat(restauranId uint, date_time string) ([]entities.Transaction, int, error) {
	return []entities.Transaction{}, 1, nil
}

func (m mockRestaurantRepository) Gets() ([]entities.RestaurantDetail, error) {
	return []entities.RestaurantDetail{{ID: 1}}, nil
}

func (m mockRestaurantRepository) Update(restaurantID uint, updateUser entities.Restaurant) (entities.Restaurant, error) {
	return entities.Restaurant{ID: 1, Email: "restaurant1Update@outlook.my"}, nil
}

func (m mockRestaurantRepository) UpdateDetail(restaurantID uint, updateUser entities.RestaurantDetail) (entities.RestaurantDetail, error) {
	return entities.RestaurantDetail{ID: 1}, nil
}

func (m mockRestaurantRepository) Delete(restaurantID uint) (entities.Restaurant, error) {
	return entities.Restaurant{ID: 1}, nil
}
func (m mockRestaurantRepository) CreateDetail(restaurantId uint, updateRestaurantD entities.RestaurantDetail) (entities.RestaurantDetail, error) {
	return entities.RestaurantDetail{ID: 1}, nil
}
func (m mockRestaurantRepository) Export(restaurantId uint, date string) ([]entities.Transaction, error) {
	return []entities.Transaction{{ID: 1}}, nil
}
