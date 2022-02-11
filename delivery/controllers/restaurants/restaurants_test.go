package restaurants

import (
	"Restobook/configs"
	"Restobook/delivery/common"
	"Restobook/delivery/controllers/auth"
	"Restobook/entities"
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var jwtTokenRestaurant string
var jwtTokenAdmin string

func Test_Register_Admin(t *testing.T) {

	ec := echo.New()

	t.Run("Register Admin", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]string{
			"email":    "admin@outlook.my",
			"password": "admin123",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")

		context := ec.NewContext(req, res)
		context.SetPath("/admin/register")

		adminCtrl := auth.NewAdminControllers(mockUserRepository{})
		adminCtrl.RegisterAdminCtrl()(context)

		responses := RestaurantsResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, 200, responses.Code)
		assert.Equal(t, "Successful Operation", responses.Message)
	})

	t.Run("Login Admin", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]string{
			"email":    "admin@outlook.my",
			"password": "admin123",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")

		context := ec.NewContext(req, res)
		context.SetPath("/admin/login")

		adminCtrl := auth.NewAdminControllers(mockUserRepository{})
		adminCtrl.LoginAdminCtrl()(context)

		responses := LoginResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)
		jwtTokenAdmin = responses.Token

		assert.Equal(t, 200, responses.Code)
		assert.Equal(t, "Successful Operation", responses.Message)
	})
}

func Test_Register_Restaurant(t *testing.T) {

	ec := echo.New()

	t.Run("400 Register Restaurant", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]int{
			"email": 1,
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")

		context := ec.NewContext(req, res)
		context.SetPath("/restaurants/register")

		restoCtrl := NewRestaurantsControllers(mockFalseRestaurantRepository{})
		restoCtrl.RegisterRestoCtrl()(context)

		responses := RestaurantsResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, 400, responses.Code)
		assert.Equal(t, "Bad Request", responses.Message)
	})

	t.Run("500 Register Restaurant", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]string{
			"email":    "restaurant1@outlook.my",
			"password": "resto123",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")

		context := ec.NewContext(req, res)
		context.SetPath("/restaurants/register")

		restoCtrl := NewRestaurantsControllers(mockFalseRestaurantRepository{})
		restoCtrl.RegisterRestoCtrl()(context)

		responses := RestaurantsResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, 500, responses.Code)
		assert.Equal(t, "Internal Server Error", responses.Message)
	})

	t.Run("200 Register Restaurant", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]string{
			"email":    "restaurant1@outlook.my",
			"password": "resto123",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")

		context := ec.NewContext(req, res)
		context.SetPath("/restaurants/register")

		restoCtrl := NewRestaurantsControllers(mockRestaurantRepository{})
		restoCtrl.RegisterRestoCtrl()(context)

		responses := RestaurantsResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, 200, responses.Code)
		assert.Equal(t, "Successful Operation", responses.Message)
	})

}

func Test_Login_Restaurant(t *testing.T) {

	ec := echo.New()

	t.Run("400 Register Restaurant", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]int{
			"email": 1,
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")

		context := ec.NewContext(req, res)
		context.SetPath("/restaurants/login")

		restoCtrl := NewRestaurantsControllers(mockRestaurantRepository{})
		restoCtrl.LoginRestoCtrl()(context)

		responses := LoginResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, 400, responses.Code)
		assert.Equal(t, "Bad Request", responses.Message)
	})

	t.Run("404 Login Restaurant", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]string{
			"email":    "restaurant1@outlook.my",
			"password": "resto1234",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")

		context := ec.NewContext(req, res)
		context.SetPath("/restaurants/login")

		restoCtrl := NewRestaurantsControllers(mockFalseRestaurantRepository{})
		restoCtrl.LoginRestoCtrl()(context)

		responses := LoginResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)
		jwtTokenRestaurant = responses.Token

		assert.Equal(t, 404, responses.Code)
		assert.Equal(t, "Not Found", responses.Message)
	})

	t.Run("200 Login Restaurant", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]string{
			"email":    "restaurant1@outlook.my",
			"password": "resto123",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")

		context := ec.NewContext(req, res)
		context.SetPath("/restaurants/login")

		restoCtrl := NewRestaurantsControllers(mockRestaurantRepository{})
		restoCtrl.LoginRestoCtrl()(context)

		responses := LoginResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)
		jwtTokenRestaurant = responses.Token

		assert.Equal(t, 200, responses.Code)
		assert.Equal(t, "Successful Operation", responses.Message)
	})

}

func Test_UpdateMyResto_Restaurant(t *testing.T) {

	ec := echo.New()

	t.Run("400 UpdateMyResto Restaurant", func(t *testing.T) {

		reqBody, _ := json.Marshal(map[string]interface{}{
			"email": 1,
		})

		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenRestaurant))

		context := ec.NewContext(req, res)
		context.SetPath("/myrestaurant")

		restaurantCtrl := NewRestaurantsControllers(mockFalseRestaurantRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(restaurantCtrl.UpdateMyRestoCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}

		responses := RestaurantsResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, 400, responses.Code)
		assert.Equal(t, "Bad Request", responses.Message)
	})

	t.Run("404 UpdateMyResto Restaurant", func(t *testing.T) {

		reqBody, _ := json.Marshal(map[string]interface{}{
			"email": "updaterestaurant1@outlook.my",
		})

		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenRestaurant))

		context := ec.NewContext(req, res)
		context.SetPath("/myrestaurant")

		restaurantCtrl := NewRestaurantsControllers(mockFalseRestaurantRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(restaurantCtrl.UpdateMyRestoCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}

		responses := RestaurantsResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, 404, responses.Code)
		assert.Equal(t, "Not Found", responses.Message)
	})

	t.Run("200 UpdateMyResto Restaurant", func(t *testing.T) {

		reqBody, _ := json.Marshal(map[string]interface{}{
			"email":    "updaterestaurant1@outlook.my",
			"password": "restaurant1234",
		})

		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenRestaurant))

		context := ec.NewContext(req, res)
		context.SetPath("/myrestaurant")

		restaurantCtrl := NewRestaurantsControllers(mockRestaurantRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(restaurantCtrl.UpdateMyRestoCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}

		responses := RestaurantsResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, 200, responses.Code)
		assert.Equal(t, "Successful Operation", responses.Message)
	})

}

func Test_Get_Restaurant(t *testing.T) {

	ec := echo.New()

	t.Run("404 Get Restaurant", func(t *testing.T) {

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenRestaurant))

		context := ec.NewContext(req, res)
		context.SetPath("/myrestaurant")

		restaurantCtrl := NewRestaurantsControllers(mockFalseRestaurantRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(restaurantCtrl.GetMyRestoCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}

		responses := RestaurantsResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, 404, responses.Code)
		assert.Equal(t, "Not Found", responses.Message)
	})

	t.Run("200 Get Restaurant", func(t *testing.T) {

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenRestaurant))

		context := ec.NewContext(req, res)
		context.SetPath("/myrestaurant")

		restaurantCtrl := NewRestaurantsControllers(mockRestaurantRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(restaurantCtrl.GetMyRestoCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}

		responses := RestaurantsResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, 200, responses.Code)
		assert.Equal(t, "Successful Operation", responses.Message)
	})

}

func Test_CreateDetail_Restaurant(t *testing.T) {

	ec := echo.New()

	t.Run("400 CreateDetail Restaurant", func(t *testing.T) {

		reqBody, _ := json.Marshal(map[string]int{
			"name": 1,
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenRestaurant))

		context := ec.NewContext(req, res)
		context.SetPath("/myrestaurant/detail")

		restaurantCtrl := NewRestaurantsControllers(mockFalseRestaurantRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(restaurantCtrl.CreateDetailRestoCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}

		responses := RestaurantsResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, 400, responses.Code)
		assert.Equal(t, "Bad Request", responses.Message)
	})

	t.Run("404 CreateDetail Restaurant", func(t *testing.T) {

		reqBody, _ := json.Marshal(map[string]interface{}{
			"name": "Restaurant 1",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenRestaurant))

		context := ec.NewContext(req, res)
		context.SetPath("/myrestaurant/detail")

		restaurantCtrl := NewRestaurantsControllers(mockFalseRestaurantRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(restaurantCtrl.CreateDetailRestoCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}

		responses := RestaurantsResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, 404, responses.Code)
		assert.Equal(t, "Not Found", responses.Message)
	})

	t.Run("200 CreateDetail Restaurant", func(t *testing.T) {

		reqBody, _ := json.Marshal(map[string]interface{}{
			"name": "Restaurant 1",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenRestaurant))

		context := ec.NewContext(req, res)
		context.SetPath("/myrestaurant/detail")

		restaurantCtrl := NewRestaurantsControllers(mockRestaurantRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(restaurantCtrl.CreateDetailRestoCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}

		responses := RestaurantsResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, 200, responses.Code)
		assert.Equal(t, "Successful Operation", responses.Message)
	})

}

func Test_UpdateDetail_Restaurant(t *testing.T) {

	ec := echo.New()

	t.Run("400 UpdateDetail Restaurant", func(t *testing.T) {

		reqBody, _ := json.Marshal(map[string]int{
			"open_hour": 1,
		})

		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenRestaurant))

		context := ec.NewContext(req, res)
		context.SetPath("/myrestaurant/detail")

		restaurantCtrl := NewRestaurantsControllers(mockFalseRestaurantRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(restaurantCtrl.UpdateDetailRestoCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}

		responses := RestaurantsResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, 400, responses.Code)
		assert.Equal(t, "Bad Request", responses.Message)
	})

	t.Run("404 UpdateDetail Restaurant", func(t *testing.T) {

		reqBody, _ := json.Marshal(map[string]interface{}{
			"open_hour": "11:30",
		})

		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenRestaurant))

		context := ec.NewContext(req, res)
		context.SetPath("/myrestaurant/detail")

		restaurantCtrl := NewRestaurantsControllers(mockFalseRestaurantRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(restaurantCtrl.UpdateDetailRestoCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}

		responses := RestaurantsResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, 404, responses.Code)
		assert.Equal(t, "Not Found", responses.Message)
	})

	t.Run("200 UpdateDetail Restaurant", func(t *testing.T) {

		reqBody, _ := json.Marshal(map[string]interface{}{
			"Open_hour": "11:30",
		})

		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenRestaurant))

		context := ec.NewContext(req, res)
		context.SetPath("/myrestaurant/detail")

		restaurantCtrl := NewRestaurantsControllers(mockRestaurantRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(restaurantCtrl.UpdateDetailRestoCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}

		responses := RestaurantsResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, 200, responses.Code)
		assert.Equal(t, "Successful Operation", responses.Message)
	})

}
func Test_GestWaiting_Restaurant(t *testing.T) {

	ec := echo.New()

	t.Run("404 GetsWaiting by admin", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenAdmin))

		context := ec.NewContext(req, res)
		context.SetPath("/admin/waiting")

		restaurantCtrl := NewRestaurantsControllers(mockFalseRestaurantRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(restaurantCtrl.GetsWaiting())(context); err != nil {
			log.Fatal(err)
			return
		}

		responses := RestaurantsResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, 404, responses.Code)
		assert.Equal(t, "Not Found", responses.Message)
	})

	t.Run("406 GetsWaiting by admin", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenRestaurant))

		context := ec.NewContext(req, res)
		context.SetPath("/admin/waiting")

		restaurantCtrl := NewRestaurantsControllers(mockFalseRestaurantRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(restaurantCtrl.GetsWaiting())(context); err != nil {
			log.Fatal(err)
			return
		}

		responses := RestaurantsResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, 406, responses.Code)
		assert.Equal(t, "Not Accepted", responses.Message)
	})

	t.Run("200 GetsWaiting by admin", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenAdmin))

		context := ec.NewContext(req, res)
		context.SetPath("/admin/waiting")

		restaurantCtrl := NewRestaurantsControllers(mockRestaurantRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(restaurantCtrl.GetsWaiting())(context); err != nil {
			log.Fatal(err)
			return
		}

		responses := RestaurantsResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, 200, responses.Code)
		assert.Equal(t, "Successful Operation", responses.Message)
	})

}

func Test_Approve_Restaurant(t *testing.T) {

	ec := echo.New()

	t.Run("400 Approve Restaurant", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]interface{}{
			"resto_id": "1",
			"status":   "OPEN",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenAdmin))

		context := ec.NewContext(req, res)
		context.SetPath("/admin/approve")

		restaurantCtrl := NewRestaurantsControllers(mockRestaurantRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(restaurantCtrl.Approve())(context); err != nil {
			log.Fatal(err)
			return
		}

		responses := RestaurantsResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, 400, responses.Code)
		assert.Equal(t, "Bad Request", responses.Message)
	})

	t.Run("406 Approve Restaurant", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]interface{}{
			"resto_id": 1,
			"status":   "OPEN",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenRestaurant))

		context := ec.NewContext(req, res)
		context.SetPath("/admin/approve")

		restaurantCtrl := NewRestaurantsControllers(mockRestaurantRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(restaurantCtrl.Approve())(context); err != nil {
			log.Fatal(err)
			return
		}

		responses := RestaurantsResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, 406, responses.Code)
		assert.Equal(t, "Not Accepted", responses.Message)
	})

	t.Run("404 Approve Restaurant", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]interface{}{
			"resto_id": 1,
			"status":   "OPEN",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenAdmin))

		context := ec.NewContext(req, res)
		context.SetPath("/admin/approve")

		restaurantCtrl := NewRestaurantsControllers(mockFalseRestaurantRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(restaurantCtrl.Approve())(context); err != nil {
			log.Fatal(err)
			return
		}

		responses := RestaurantsResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, 404, responses.Code)
		assert.Equal(t, "Not Found", responses.Message)
	})

	t.Run("200 Approve Restaurant", func(t *testing.T) {

		reqBody, _ := json.Marshal(map[string]interface{}{
			"resto_id": 1,
			"status":   "OPEN",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenAdmin))

		context := ec.NewContext(req, res)
		context.SetPath("/admin/approve")

		restaurantCtrl := NewRestaurantsControllers(mockRestaurantRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(restaurantCtrl.Approve())(context); err != nil {
			log.Fatal(err)
			return
		}

		responses := RestaurantsResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, 200, responses.Code)
		assert.Equal(t, "Successful Operation", responses.Message)
	})

}

func Test_Gets_Restaurant(t *testing.T) {

	ec := echo.New()

	t.Run("404 Gets Restaurant", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")

		context := ec.NewContext(req, res)
		context.SetPath("/restaurants")

		restaurantCtrl := NewRestaurantsControllers(mockFalseRestaurantRepository{})
		restaurantCtrl.Gets()(context)

		responses := RestaurantsResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, 404, responses.Code)
		assert.Equal(t, "Not Found", responses.Message)
	})

	t.Run("200 Gets Restaurant", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")

		context := ec.NewContext(req, res)
		context.SetPath("/restaurants")

		restaurantCtrl := NewRestaurantsControllers(mockRestaurantRepository{})
		restaurantCtrl.Gets()(context)

		responses := RestaurantsResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, 200, responses.Code)
		assert.Equal(t, "Successful Operation", responses.Message)
	})

}

func Test_GetsByOpen_Restaurant(t *testing.T) {

	ec := echo.New()

	t.Run("404 GetsByOpen Restaurant", func(t *testing.T) {
		query := make(url.Values)
		query.Set("date_time", "2022-03-07 08:00:00")
		req := httptest.NewRequest(http.MethodGet, "/?"+query.Encode(), nil)
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")

		context := ec.NewContext(req, res)
		context.SetPath("/restaurants/open")

		restaurantCtrl := NewRestaurantsControllers(mockFalseRestaurantRepository{})
		restaurantCtrl.GetsByOpen()(context)

		responses := RestaurantsResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, 404, responses.Code)
		assert.Equal(t, "Not Found", responses.Message)
	})

	t.Run("404 GetsByOpen Restaurant", func(t *testing.T) {
		query := make(url.Values)
		query.Set("date_time", "2022-03-10 08:00:00")
		req := httptest.NewRequest(http.MethodGet, "/?"+query.Encode(), nil)
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")

		context := ec.NewContext(req, res)
		context.SetPath("/restaurants/open")

		restaurantCtrl := NewRestaurantsControllers(mockRestaurantRepository{})
		restaurantCtrl.GetsByOpen()(context)

		responses := RestaurantsResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, 404, responses.Code)
		assert.Equal(t, "Not Found", responses.Message)
	})

	t.Run("200 GetsByOpen Restaurant", func(t *testing.T) {
		query := make(url.Values)
		query.Set("date_time", "2022-03-08 08:00:00")
		req := httptest.NewRequest(http.MethodGet, "/?"+query.Encode(), nil)
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")

		context := ec.NewContext(req, res)
		context.SetPath("/restaurants/open")

		restaurantCtrl := NewRestaurantsControllers(mockRestaurantRepository{})
		restaurantCtrl.GetsByOpen()(context)

		responses := RestaurantsResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, 200, responses.Code)
		assert.Equal(t, "Successful Operation", responses.Message)
	})

	t.Run("200 GetsByOpen Restaurant", func(t *testing.T) {
		query := make(url.Values)
		query.Set("date_time", "2022-03-09 08:00:00")
		req := httptest.NewRequest(http.MethodGet, "/?"+query.Encode(), nil)
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")

		context := ec.NewContext(req, res)
		context.SetPath("/restaurants/open")

		restaurantCtrl := NewRestaurantsControllers(mockRestaurantRepository{})
		restaurantCtrl.GetsByOpen()(context)

		responses := RestaurantsResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, 200, responses.Code)
		assert.Equal(t, "Successful Operation", responses.Message)
	})

	t.Run("200 GetsByOpen Restaurant", func(t *testing.T) {
		query := make(url.Values)
		query.Set("date_time", "2022-03-07 11:30:00")
		req := httptest.NewRequest(http.MethodGet, "/?"+query.Encode(), nil)
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")

		context := ec.NewContext(req, res)
		context.SetPath("/restaurants/open")

		restaurantCtrl := NewRestaurantsControllers(mockRestaurantRepository{})
		restaurantCtrl.GetsByOpen()(context)

		responses := RestaurantsResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, 200, responses.Code)
		assert.Equal(t, "Successful Operation", responses.Message)
	})

}
func Test_Delete_Restaurant(t *testing.T) {

	ec := echo.New()

	t.Run("400 Delete Restaurant", func(t *testing.T) {

		req := httptest.NewRequest(http.MethodDelete, "/", nil)
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenRestaurant))

		context := ec.NewContext(req, res)
		context.SetPath("/restaurant")

		restaurantCtrl := NewRestaurantsControllers(mockFalseRestaurantRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(restaurantCtrl.DeleteRestoCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}

		responses := RestaurantsResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, 406, responses.Code)
		assert.Equal(t, "Not Accepted", responses.Message)
	})

	t.Run("400 Delete Restaurant", func(t *testing.T) {

		reqBody, _ := json.Marshal(map[string]interface{}{
			"resto_id": "2",
		})

		req := httptest.NewRequest(http.MethodDelete, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenAdmin))

		context := ec.NewContext(req, res)
		context.SetPath("/restaurant")

		restaurantCtrl := NewRestaurantsControllers(mockFalseRestaurantRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(restaurantCtrl.DeleteRestoCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}

		responses := RestaurantsResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, 400, responses.Code)
		assert.Equal(t, "Bad Request", responses.Message)
	})

	t.Run("404 Delete Restaurant", func(t *testing.T) {

		reqBody, _ := json.Marshal(map[string]interface{}{
			"resto_id": 1,
		})

		req := httptest.NewRequest(http.MethodDelete, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenAdmin))

		context := ec.NewContext(req, res)
		context.SetPath("/restaurant")

		restaurantCtrl := NewRestaurantsControllers(mockFalseRestaurantRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(restaurantCtrl.DeleteRestoCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}

		responses := RestaurantsResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, 404, responses.Code)
		assert.Equal(t, "Not Found", responses.Message)
	})

	t.Run("200 Delete Restaurant", func(t *testing.T) {

		reqBody, _ := json.Marshal(map[string]interface{}{
			"resto_id": 1,
		})

		req := httptest.NewRequest(http.MethodDelete, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenAdmin))

		context := ec.NewContext(req, res)
		context.SetPath("/restaurant")

		restaurantCtrl := NewRestaurantsControllers(mockRestaurantRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(restaurantCtrl.DeleteRestoCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}

		responses := RestaurantsResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, 200, responses.Code)
		assert.Equal(t, "Successful Operation", responses.Message)
	})

}

func Test_Export_Restaurant(t *testing.T) {

	config := configs.GetConfig()
	fmt.Println(config)

	ec := echo.New()

	t.Run("200 Login Restaurant", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]string{
			"email":    "restaurant1@outlook.my",
			"password": "resto123",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")

		context := ec.NewContext(req, res)
		context.SetPath("/restaurants/login")

		restoCtrl := NewRestaurantsControllers(mockRestaurantRepository{})
		restoCtrl.LoginRestoCtrl()(context)

		responses := LoginResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)
		jwtTokenRestaurant = responses.Token

		assert.Equal(t, 200, responses.Code)
		assert.Equal(t, "Successful Operation", responses.Message)
	})

	t.Run("500 Export Restaurant", func(t *testing.T) {
		query := make(url.Values)
		query.Set("date_time", "2022-03-07 00:00:00")

		req := httptest.NewRequest(http.MethodPost, "/?"+query.Encode(), nil)
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenRestaurant))

		context := ec.NewContext(req, res)
		context.SetPath("/restaurant/report")

		restaurantCtrl := NewRestaurantsControllers(mockFalseRestaurantRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(restaurantCtrl.ExportPDF())(context); err != nil {
			log.Fatal(err)
			return
		}

		responses := RestaurantsResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, 500, responses.Code)
		assert.Equal(t, "Internal Server Error", responses.Message)
	})

	t.Run("200 Export Restaurant", func(t *testing.T) {
		// fmt.Println("====> TOKEN", jwtTokenRestaurant)
		query := make(url.Values)
		query.Set("date_time", "2022-03-07 00:00:00")

		req := httptest.NewRequest(http.MethodPost, "/?"+query.Encode(), nil)
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenRestaurant))

		context := ec.NewContext(req, res)
		context.SetPath("/restaurant/report")

		// fmt.Println("====>REQ", req)
		// fmt.Println("====>RES", res)

		restaurantCtrl := NewRestaurantsControllers(mockRestaurantRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(restaurantCtrl.ExportPDF())(context); err != nil {
			log.Fatal(err)
			return
		}

		responses := RestaurantsResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, 200, responses.Code)
		assert.Equal(t, "Successful Operation", responses.Message)
	})

}

type mockUserRepository struct{}

func (m mockUserRepository) RegisterAdmin(newUser entities.User) (entities.User, error) {
	return entities.User{
		Model:       gorm.Model{},
		ID:          1,
		Email:       "admin@outlook.my",
		Password:    "admin123",
		Name:        "admin",
		PhoneNumber: "0877",
		Reputation:  999,
		Balance:     100,
	}, nil
}

func (m mockUserRepository) Register(newUser entities.User) (entities.User, error) {
	return entities.User{
		Model:       gorm.Model{},
		ID:          2,
		Email:       "herlianto@outlook.my",
		Password:    "herlianto123",
		Name:        "herlianto",
		PhoneNumber: "0877",
		Reputation:  999,
		Balance:     80,
	}, nil
}

func (m mockUserRepository) LoginUser(email, password string) (entities.User, error) {
	hash := sha256.Sum256([]byte("admin123"))
	passwordS := fmt.Sprintf("%x", hash[:])
	return entities.User{
		Model:       gorm.Model{},
		ID:          1,
		Email:       "admin@outlook.my",
		Password:    passwordS,
		Name:        "admin",
		PhoneNumber: "0877",
		Reputation:  999,
		Balance:     100,
	}, nil
}

func (m mockUserRepository) Get(userID uint) (entities.User, error) {
	return entities.User{
		Model:       gorm.Model{},
		ID:          1,
		Email:       "admin@outlook.my",
		Password:    "admin123",
		Name:        "admin",
		PhoneNumber: "0877",
		Reputation:  999,
		Balance:     100,
	}, nil
}

func (m mockUserRepository) Update(userID uint, updateUser entities.User) (entities.User, error) {
	return entities.User{
		Model:       gorm.Model{},
		ID:          1,
		Email:       "admin@outlook.my",
		Password:    "admin123",
		Name:        "admin",
		PhoneNumber: "0877",
		Reputation:  999,
		Balance:     100,
	}, nil
}

func (m mockUserRepository) Delete(userID uint) (entities.User, error) {
	return entities.User{
		Model:       gorm.Model{},
		ID:          1,
		Email:       "admin@outlook.my",
		Password:    "admin123",
		Name:        "admin",
		PhoneNumber: "0877",
		Reputation:  999,
		Balance:     100,
	}, nil
}

type mockRestaurantRepository struct{}

func (m mockRestaurantRepository) Register(newUser entities.Restaurant) (entities.Restaurant, error) {
	hash := sha256.Sum256([]byte("resto123"))
	passwordS := fmt.Sprintf("%x", hash[:])
	return entities.Restaurant{
		Model:              gorm.Model{},
		ID:                 1,
		Email:              "restaurant1@outlook.my",
		Password:           passwordS,
		RestaurantDetailID: 1,
		RestaurantDetail:   entities.RestaurantDetail{},
	}, nil
}

func (m mockRestaurantRepository) Login(email, password string) (entities.Restaurant, error) {
	hash := sha256.Sum256([]byte("resto123"))
	passwordS := fmt.Sprintf("%x", hash[:])
	return entities.Restaurant{
		Model:              gorm.Model{},
		ID:                 1,
		Email:              email,
		Password:           passwordS,
		RestaurantDetailID: 1,
		RestaurantDetail:   entities.RestaurantDetail{},
	}, nil
}

func (m mockRestaurantRepository) Update(restaurantID uint, updateUser entities.Restaurant) (entities.Restaurant, error) {
	hash := sha256.Sum256([]byte("resto1234"))
	passwordS := fmt.Sprintf("%x", hash[:])
	return entities.Restaurant{
		Model:              gorm.Model{},
		ID:                 restaurantID,
		Email:              "updaterestaurant1@outlook.my",
		Password:           passwordS,
		RestaurantDetailID: restaurantID,
		RestaurantDetail:   entities.RestaurantDetail{},
	}, nil
}

func (m mockRestaurantRepository) Get(restaurantID uint) (entities.Restaurant, entities.RestaurantDetail, error) {
	return entities.Restaurant{
			Model: gorm.Model{},
			ID:    restaurantID,
			Email: "restaurant1@outlook.my",
		}, entities.RestaurantDetail{
			Model:       gorm.Model{},
			ID:          restaurantID,
			Name:        "Restaurant Name",
			PhoneNumber: "",
			Status:      "DISABLED",
		}, nil
}

func (m mockRestaurantRepository) CreateDetail(restaurantID uint, updateUser entities.RestaurantDetail) (entities.RestaurantDetail, error) {
	return entities.RestaurantDetail{
		Model:          gorm.Model{},
		ID:             restaurantID,
		Name:           "Restaurant Nasi Padang Jago",
		Open_Hour:      "11:00",
		Close_Hour:     "15:00",
		Open:           "Monday,Tuesday",
		Close:          "Wednesday,Thursday,Friday,Saturday,Sunday",
		Price:          10000,
		Latitude:       0,
		Longitude:      0,
		City:           "Jakarta",
		Address:        "JL.Taman Daan Mogot No.2",
		PhoneNumber:    "0877",
		ProfilePicture: "https://",
		Seats:          10,
		Description:    "Resto Nasi Padang",
	}, nil
}

func (m mockRestaurantRepository) UpdateDetail(restaurantID uint, updateUser entities.RestaurantDetail) (entities.RestaurantDetail, error) {
	return entities.RestaurantDetail{
		Model:          gorm.Model{},
		ID:             restaurantID,
		Name:           "Restaurant Nasi Padang Jago",
		Open_Hour:      "11:30",
		Close_Hour:     "15:00",
		Open:           "Monday,Tuesday,Friday",
		Close:          "Wednesday,Thursday,Saturday,Sunday",
		Price:          10000,
		Latitude:       0,
		Longitude:      0,
		City:           "Jakarta",
		Address:        "JL.Taman Daan Mogot No.2",
		PhoneNumber:    "0877",
		ProfilePicture: "https://",
		Seats:          100,
		Status:         "Waiting For Approval",
		Description:    "Resto Nasi Padang",
	}, nil
}

func (m mockRestaurantRepository) GetsWaiting() ([]entities.RestaurantDetail, error) {
	return []entities.RestaurantDetail{{
		Model:          gorm.Model{},
		ID:             0,
		Name:           "Restaurant Nasi Padang Jago",
		Open_Hour:      "11:30",
		Close_Hour:     "15:00",
		Open:           "Monday,Tuesday,Friday",
		Close:          "Wednesday,Thursday,Saturday,Sunday",
		Price:          10000,
		Latitude:       0,
		Longitude:      0,
		City:           "Jakarta",
		Address:        "JL.Taman Daan Mogot No.2",
		PhoneNumber:    "0877",
		ProfilePicture: "https://",
		Seats:          100,
		Status:         "Waiting For Approval",
		Description:    "Resto Nasi Padang",
		Rating:         []entities.Rating{},
		Restaurant:     []entities.Restaurant{},
	}}, nil
}

func (m mockRestaurantRepository) Approve(restaurantID uint, status string) (entities.RestaurantDetail, error) {
	return entities.RestaurantDetail{
		Model:       gorm.Model{},
		ID:          1,
		Name:        "Restauran Nasi Padang Jago",
		PhoneNumber: "0877",
		Status:      status,
		Rating:      []entities.Rating{},
		Restaurant: []entities.Restaurant{{
			Model:              gorm.Model{},
			ID:                 restaurantID,
			Email:              "restaurant1@outlook.my",
			RestaurantDetailID: restaurantID,
			RestaurantDetail:   entities.RestaurantDetail{},
		}},
	}, nil
}

func (m mockRestaurantRepository) Gets() ([]entities.RestaurantDetail, error) {
	return []entities.RestaurantDetail{{
		Model:          gorm.Model{},
		ID:             1,
		Name:           "Restaurant Nasi Padang Jago",
		Open_Hour:      "11:30",
		Close_Hour:     "15:00",
		Open:           "Monday",
		Close:          "Tuesday,Wednesday,Thursday,Friday,Saturday,Sunday",
		Price:          10000,
		Latitude:       0,
		Longitude:      0,
		City:           "Jakarta",
		Address:        "JL.Taman Daan Mogot No.2",
		PhoneNumber:    "0877",
		ProfilePicture: "https://",
		Seats:          100,
		Status:         "OPEN",
		Description:    "Resto Nasi Padang",
		Rating:         []entities.Rating{{}},
		Restaurant:     []entities.Restaurant{},
	}, {
		Model:          gorm.Model{},
		ID:             2,
		Name:           "Restaurant Bubur Padang Jago",
		Open_Hour:      "07:00",
		Close_Hour:     "11:00",
		Open:           "Tuesday",
		Close:          "Wednesday,Thursday,Friday,Saturday,Sunday,Monday",
		Price:          5000,
		Latitude:       0,
		Longitude:      0,
		City:           "Jakarta",
		Address:        "JL.Taman Daan Mogot No.3",
		PhoneNumber:    "0877",
		ProfilePicture: "https://",
		Seats:          100,
		Status:         "OPEN",
		Description:    "Resto Bubur Padang",
		Rating:         []entities.Rating{},
		Restaurant:     []entities.Restaurant{},
	}}, nil
}

func (m mockRestaurantRepository) GetsByOpen(open int) ([]entities.RestaurantDetail, error) {
	if open != 3 {
		return []entities.RestaurantDetail{{
			Model:          gorm.Model{},
			ID:             1,
			Name:           "Restaurant Nasi Padang Jago",
			Open_Hour:      "11:30",
			Close_Hour:     "15:00",
			Open:           "Monday",
			Close:          "Tuesday,Wednesday,Thursday,Friday,Saturday,Sunday",
			Price:          10000,
			Latitude:       0,
			Longitude:      0,
			City:           "Jakarta",
			Address:        "JL.Taman Daan Mogot No.2",
			PhoneNumber:    "0877",
			ProfilePicture: "https://",
			Seats:          100,
			Status:         "OPEN",
			Description:    "Resto Nasi Padang",
			Rating:         []entities.Rating{},
			Restaurant: []entities.Restaurant{{
				Model:              gorm.Model{},
				ID:                 1,
				Email:              "restaurant1@outlook.my",
				RestaurantDetailID: 1,
				RestaurantDetail:   entities.RestaurantDetail{},
			}},
		}, {
			Model:          gorm.Model{},
			ID:             2,
			Name:           "Restaurant Bubur Padang Jago",
			Open_Hour:      "07:00",
			Close_Hour:     "11:00",
			Open:           "Tuesday",
			Close:          "Wednesday,Thursday,Friday,Saturday,Sunday,Monday",
			Price:          5000,
			Latitude:       0,
			Longitude:      0,
			City:           "Jakarta",
			Address:        "JL.Taman Daan Mogot No.3",
			PhoneNumber:    "0877",
			ProfilePicture: "https://",
			Seats:          100,
			Status:         "OPEN",
			Description:    "Resto Bubu Padang",
			Rating:         []entities.Rating{},
			Restaurant: []entities.Restaurant{{
				Model:              gorm.Model{},
				ID:                 2,
				Email:              "restaurant2@outlook.my",
				RestaurantDetailID: 2,
				RestaurantDetail:   entities.RestaurantDetail{},
			}},
		}, {
			Model:          gorm.Model{},
			ID:             3,
			Name:           "Retaurant Failed",
			Open_Hour:      "07:00",
			Close_Hour:     "11:00",
			Open:           "Thursday",
			Close:          "Friday,Saturday,Sunday,Monday,Tuesday,Wednesday",
			Price:          5000,
			Latitude:       0,
			Longitude:      0,
			City:           "Jakarta",
			Address:        "JL.Taman Daan Mogot No.3",
			PhoneNumber:    "0877",
			ProfilePicture: "https://",
			Seats:          100,
			Status:         "OPEN",
			Description:    "Resto Bubu Padang",
			Rating:         []entities.Rating{},
			Restaurant: []entities.Restaurant{{
				Model:              gorm.Model{},
				ID:                 3,
				Email:              "restaurant3@outlook.my",
				RestaurantDetailID: 3,
				RestaurantDetail:   entities.RestaurantDetail{},
			}},
		}}, nil
	} else {
		return []entities.RestaurantDetail{{
			Model:          gorm.Model{},
			ID:             0,
			Name:           "",
			Open_Hour:      "",
			Close_Hour:     "",
			Open:           "",
			Close:          "",
			Price:          0,
			Latitude:       0,
			Longitude:      0,
			City:           "",
			Address:        "",
			PhoneNumber:    "",
			ProfilePicture: "",
			Seats:          0,
			Status:         "",
			Description:    "",
			Rating:         []entities.Rating{},
			Restaurant: []entities.Restaurant{{
				Model:              gorm.Model{},
				ID:                 0,
				Email:              "",
				RestaurantDetailID: 0,
				RestaurantDetail:   entities.RestaurantDetail{},
			}},
		}}, errors.New("FAILED GETS BY OPEN")
	}
}

func (m mockRestaurantRepository) GetExistSeat(restauranId uint, date_time string) ([]entities.Transaction, int, error) {

	if restauranId != 3 {
		date_find1 := "2022-03-07 12:00:00"
		date_parse1, _ := time.Parse("2006-01-02 15:04:05", date_find1)

		date_find2 := "2022-03-08 17:30:00"
		date_parse2, _ := time.Parse("2006-01-02 15:04:05", date_find2)

		return []entities.Transaction{{
			Model:        gorm.Model{},
			ID:           1,
			UserID:       1,
			RestaurantID: 1,
			DateTime:     date_parse1,
			Persons:      1,
			Total:        10000,
			Status:       "SUCCESS",
			User:         entities.User{},
			Restaurant:   entities.Restaurant{},
		}, {
			Model:        gorm.Model{},
			ID:           2,
			UserID:       1,
			RestaurantID: 2,
			DateTime:     date_parse2,
			Persons:      1,
			Total:        5000,
			Status:       "SUCCESS",
			User:         entities.User{},
			Restaurant:   entities.Restaurant{},
		}}, 1, nil
	} else {
		return []entities.Transaction{{
			Model:        gorm.Model{},
			ID:           0,
			UserID:       0,
			RestaurantID: 0,
			DateTime:     time.Time{},
			Persons:      0,
			Total:        0,
			Status:       "",
			User:         entities.User{},
			Restaurant:   entities.Restaurant{},
		}}, 0, errors.New("FAILED GETS EXIST SEAT")
	}
}

func (m mockRestaurantRepository) Delete(restaurantID uint) (entities.Restaurant, error) {
	return entities.Restaurant{
		Model:              gorm.Model{},
		ID:                 restaurantID,
		Email:              "updaterestaurant1@outlook.my",
		RestaurantDetailID: restaurantID,
	}, nil
}

func (m mockRestaurantRepository) Export(restaurantId uint, date string) ([]entities.Transaction, error) {
	return []entities.Transaction{{
		Model:        gorm.Model{},
		ID:           restaurantId,
		UserID:       1,
		RestaurantID: restaurantId,
		DateTime:     time.Time{},
		Persons:      1,
		Total:        10000,
		Status:       "Success",
		User: entities.User{
			Model:       gorm.Model{},
			ID:          restaurantId,
			Email:       "user1@outlook.my",
			Password:    "user123",
			Name:        "user1",
			PhoneNumber: "0877",
			Reputation:  80,
			Balance:     999,
		},
		Restaurant: entities.Restaurant{
			Model:              gorm.Model{},
			ID:                 restaurantId,
			Email:              "restaurant1@outlook.my",
			Password:           "resto123",
			RestaurantDetailID: restaurantId,
			RestaurantDetail: entities.RestaurantDetail{
				Model:          gorm.Model{},
				ID:             restaurantId,
				Name:           "Restaurant Nasi Padang Jago",
				Open_Hour:      "11:30",
				Close_Hour:     "15:00",
				Open:           "Monday",
				Close:          "Tuesday,Wednesday,Thursday,Friday,Saturday,Sunday",
				Price:          10000,
				Latitude:       0,
				Longitude:      0,
				City:           "Jakarta",
				Address:        "JL.Taman Daan Mogot No.2",
				PhoneNumber:    "0877",
				ProfilePicture: "https://",
				Seats:          100,
				Status:         "OPEN",
				Description:    "Resto Nasi Padang",
				Rating:         []entities.Rating{},
				Restaurant:     []entities.Restaurant{},
			},
		},
	}, {
		Model:        gorm.Model{},
		ID:           restaurantId,
		UserID:       1,
		RestaurantID: restaurantId,
		DateTime:     time.Time{},
		Persons:      1,
		Total:        1,
		Status:       "Fail",
		User:         entities.User{},
		Restaurant:   entities.Restaurant{},
	}, {
		Model:        gorm.Model{},
		ID:           restaurantId,
		UserID:       1,
		RestaurantID: restaurantId,
		DateTime:     time.Time{},
		Persons:      1,
		Total:        1,
		Status:       "Cancel",
		User:         entities.User{},
		Restaurant:   entities.Restaurant{},
	}, {
		Model:        gorm.Model{},
		ID:           restaurantId,
		UserID:       1,
		RestaurantID: restaurantId,
		DateTime:     time.Time{},
		Persons:      1,
		Total:        1,
		Status:       "Rejected",
		User:         entities.User{},
		Restaurant:   entities.Restaurant{},
	}}, nil
}

type mockFalseRestaurantRepository struct{}

func (m mockFalseRestaurantRepository) Register(newUser entities.Restaurant) (entities.Restaurant, error) {
	hash := sha256.Sum256([]byte("resto123"))
	passwordS := fmt.Sprintf("%x", hash[:])
	return entities.Restaurant{
		Model:              gorm.Model{},
		ID:                 0,
		Email:              newUser.Email,
		Password:           passwordS,
		RestaurantDetailID: 0,
		RestaurantDetail:   entities.RestaurantDetail{},
	}, errors.New("FAILED REGISTER")
}

func (m mockFalseRestaurantRepository) Login(email, password string) (entities.Restaurant, error) {
	hash := sha256.Sum256([]byte("resto123"))
	passwordS := fmt.Sprintf("%x", hash[:])
	return entities.Restaurant{
		Model:              gorm.Model{},
		ID:                 0,
		Email:              email,
		Password:           passwordS,
		RestaurantDetailID: 0,
		RestaurantDetail:   entities.RestaurantDetail{},
	}, errors.New("FAILED LOGIN")
}

func (m mockFalseRestaurantRepository) Update(restaurantID uint, updateUser entities.Restaurant) (entities.Restaurant, error) {
	return entities.Restaurant{
		Model:              gorm.Model{},
		ID:                 0,
		Email:              "",
		Password:           "",
		RestaurantDetailID: 0,
		RestaurantDetail:   entities.RestaurantDetail{},
	}, errors.New("FAILED UPDATE")
}

func (m mockFalseRestaurantRepository) Get(restaurantID uint) (entities.Restaurant, entities.RestaurantDetail, error) {
	return entities.Restaurant{
			Model: gorm.Model{},
			ID:    0,
			Email: "",
		}, entities.RestaurantDetail{
			Model:       gorm.Model{},
			ID:          0,
			Name:        "",
			PhoneNumber: "",
			Status:      "",
		}, errors.New("FAILED GET")
}

func (m mockFalseRestaurantRepository) CreateDetail(restaurantID uint, updateUser entities.RestaurantDetail) (entities.RestaurantDetail, error) {
	return entities.RestaurantDetail{
		Model:          gorm.Model{},
		ID:             0,
		Name:           "",
		Open_Hour:      "",
		Close_Hour:     "",
		Open:           "",
		Close:          "y",
		Price:          0,
		Latitude:       0,
		Longitude:      0,
		City:           "",
		Address:        "",
		PhoneNumber:    "",
		ProfilePicture: "",
		Seats:          0,
		Description:    "",
	}, errors.New("FAILED CREATE DETAIL")
}

func (m mockFalseRestaurantRepository) UpdateDetail(restaurantID uint, updateUser entities.RestaurantDetail) (entities.RestaurantDetail, error) {
	return entities.RestaurantDetail{
		Model:          gorm.Model{},
		ID:             0,
		Name:           "",
		Open_Hour:      "",
		Close_Hour:     "",
		Open:           "",
		Close:          "",
		Price:          0,
		Latitude:       0,
		Longitude:      0,
		City:           "",
		Address:        "",
		PhoneNumber:    "",
		ProfilePicture: "",
		Seats:          0,
		Status:         "",
		Description:    "",
	}, errors.New("FAILED UPDATE DETAIL")
}

func (m mockFalseRestaurantRepository) GetsWaiting() ([]entities.RestaurantDetail, error) {
	return []entities.RestaurantDetail{{
		Model:          gorm.Model{},
		ID:             0,
		Name:           "",
		Open_Hour:      "",
		Close_Hour:     "",
		Open:           "",
		Close:          "",
		Price:          0,
		Latitude:       0,
		Longitude:      0,
		City:           "",
		Address:        "",
		PhoneNumber:    "",
		ProfilePicture: "",
		Seats:          0,
		Status:         "",
		Description:    "",
		Rating:         []entities.Rating{},
		Restaurant:     []entities.Restaurant{},
	}}, errors.New("FAILED GETS WAITING")
}

func (m mockFalseRestaurantRepository) Approve(restaurantID uint, status string) (entities.RestaurantDetail, error) {
	return entities.RestaurantDetail{
		Model:       gorm.Model{},
		ID:          0,
		Name:        "",
		PhoneNumber: "",
		Status:      "",
	}, errors.New("FAILED GET APPROVE")
}

func (m mockFalseRestaurantRepository) Gets() ([]entities.RestaurantDetail, error) {
	return []entities.RestaurantDetail{{
		Model:          gorm.Model{},
		ID:             0,
		Name:           "",
		Open_Hour:      "",
		Close_Hour:     "",
		Open:           "",
		Close:          "",
		Price:          0,
		Latitude:       0,
		Longitude:      0,
		City:           "",
		Address:        "",
		PhoneNumber:    "",
		ProfilePicture: "",
		Seats:          0,
		Status:         "",
		Description:    "",
		Rating:         []entities.Rating{},
		Restaurant: []entities.Restaurant{{
			Model:              gorm.Model{},
			ID:                 0,
			Email:              "",
			Password:           "",
			RestaurantDetailID: 0,
			RestaurantDetail:   entities.RestaurantDetail{},
		}},
	}}, errors.New("FAILED GETS")
}

func (m mockFalseRestaurantRepository) GetsByOpen(open int) ([]entities.RestaurantDetail, error) {
	return []entities.RestaurantDetail{{
		Model:          gorm.Model{},
		ID:             0,
		Name:           "",
		Open_Hour:      "",
		Close_Hour:     "",
		Open:           "",
		Close:          "",
		Price:          0,
		Latitude:       0,
		Longitude:      0,
		City:           "",
		Address:        "",
		PhoneNumber:    "",
		ProfilePicture: "",
		Seats:          0,
		Status:         "",
		Description:    "",
		Rating:         []entities.Rating{},
		Restaurant:     []entities.Restaurant{},
	}}, errors.New("FAILED GET BY OPEN")
}

func (m mockFalseRestaurantRepository) GetExistSeat(restauranId uint, date_time string) ([]entities.Transaction, int, error) {

	return []entities.Transaction{{
		Model:        gorm.Model{},
		ID:           0,
		UserID:       0,
		RestaurantID: 0,
		DateTime:     time.Time{},
		Persons:      0,
		Total:        0,
		Status:       "",
		User:         entities.User{},
		Restaurant:   entities.Restaurant{},
	}}, 1, errors.New("FAILED GET EXIST SEAT")
}

func (m mockFalseRestaurantRepository) Delete(restaurantID uint) (entities.Restaurant, error) {
	return entities.Restaurant{
		Model:              gorm.Model{},
		ID:                 restaurantID,
		Email:              "updaterestaurant1@outlook.my",
		RestaurantDetailID: restaurantID,
	}, errors.New("FAILED DELETE")
}

func (m mockFalseRestaurantRepository) Export(restaurantId uint, date string) ([]entities.Transaction, error) {
	return []entities.Transaction{{
		Model:        gorm.Model{},
		ID:           restaurantId,
		UserID:       0,
		RestaurantID: restaurantId,
		DateTime:     time.Time{},
		Persons:      0,
		Total:        0,
		Status:       "",
		User:         entities.User{},
		Restaurant:   entities.Restaurant{},
	}}, errors.New("FAILED EXPORT PDF")
}
