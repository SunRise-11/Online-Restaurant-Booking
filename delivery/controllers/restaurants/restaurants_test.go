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
	config := configs.GetConfig()
	fmt.Println(config)

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

		responses := RestaurantResponseFormat{}
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
	config := configs.GetConfig()
	fmt.Println(config)

	ec := echo.New()

	t.Run("Bad Request Register Restaurant", func(t *testing.T) {
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

		responses := RestaurantResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, 400, responses.Code)
		assert.Equal(t, "Bad Request", responses.Message)
	})

	t.Run("ISE Register Restaurant", func(t *testing.T) {
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

		responses := RestaurantResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, 500, responses.Code)
		assert.Equal(t, "Internal Server Error", responses.Message)
	})

	t.Run("Success Register Restaurant", func(t *testing.T) {
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

		responses := RestaurantResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, 200, responses.Code)
		assert.Equal(t, "Successful Operation", responses.Message)
	})

}

func Test_Login_Restaurant(t *testing.T) {
	config := configs.GetConfig()
	fmt.Println(config)

	ec := echo.New()

	t.Run("Bad Request Register Restaurant", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]int{
			"email": 1,
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

		assert.Equal(t, 400, responses.Code)
		assert.Equal(t, "Bad Request", responses.Message)
	})

	t.Run("Not Found Login Restaurant", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]string{
			"email":    "restaurant1@outlook.my",
			"password": "resto123",
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

	t.Run("Success Login Restaurant", func(t *testing.T) {
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

func Test_GetWaiting_Restaurant(t *testing.T) {
	config := configs.GetConfig()
	fmt.Println(config)

	ec := echo.New()

	t.Run("Not Found Waiting for approval by admin", func(t *testing.T) {
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

		responses := RestaurantResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, 404, responses.Code)
		assert.Equal(t, "Not Found", responses.Message)
	})

	t.Run("Token NOT Admin", func(t *testing.T) {
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

	t.Run("Not Accepted Waiting for approval by admin", func(t *testing.T) {
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

		responses := RestaurantResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, 406, responses.Code)
		assert.Equal(t, "Not Accepted", responses.Message)
	})

	t.Run("Success Get Waiting for approval by admin", func(t *testing.T) {
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

		responses := RestaurantResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, 200, responses.Code)
		assert.Equal(t, "Successful Operation", responses.Message)
	})

}
func Test_Approve_Restaurant(t *testing.T) {
	config := configs.GetConfig()
	fmt.Println(config)

	ec := echo.New()

	t.Run("Bad Request Restaurant", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]interface{}{
			"resto_id": "1",
			"status":   "OPEN",
		})

		req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(reqBody))
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

		responses := RestaurantResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, 400, responses.Code)
		assert.Equal(t, "Bad Request", responses.Message)
	})

	t.Run("Not Accepted Approve Restaurant", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]interface{}{
			"resto_id": 1,
			"status":   "OPEN",
		})

		req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(reqBody))
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

		responses := RestaurantResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, 406, responses.Code)
		assert.Equal(t, "Not Accepted", responses.Message)
	})

	t.Run("Not Found Approve Restaurant", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]interface{}{
			"resto_id": 1,
			"status":   "OPEN",
		})

		req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(reqBody))
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

		responses := RestaurantResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, 404, responses.Code)
		assert.Equal(t, "Not Found", responses.Message)
	})

	t.Run("Success Approve Restaurant", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]interface{}{
			"resto_id": 1,
			"status":   "OPEN",
		})

		req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(reqBody))
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

		responses := RestaurantResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, 200, responses.Code)
		assert.Equal(t, "Successful Operation", responses.Message)
	})

}
func Test_GETS_Restaurant(t *testing.T) {
	config := configs.GetConfig()
	fmt.Println(config)

	ec := echo.New()

	t.Run("Not Found Get Restaurants", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")

		context := ec.NewContext(req, res)
		context.SetPath("/restaurants")

		restaurantCtrl := NewRestaurantsControllers(mockFalseRestaurantRepository{})
		restaurantCtrl.Gets()(context)

		responses := RestaurantResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, 404, responses.Code)
		assert.Equal(t, "Not Found", responses.Message)
	})

	t.Run("Success Get Restaurants", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")

		context := ec.NewContext(req, res)
		context.SetPath("/restaurants")

		restaurantCtrl := NewRestaurantsControllers(mockRestaurantRepository{})
		restaurantCtrl.Gets()(context)

		responses := RestaurantResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, 200, responses.Code)
		assert.Equal(t, "Successful Operation", responses.Message)
	})
}

func Test_GetByOpen_Restaurant(t *testing.T) {
	config := configs.GetConfig()
	fmt.Println(config)

	ec := echo.New()

	// t.Run("Not Found GetByOpen Restaurants", func(t *testing.T) {
	// 	req := httptest.NewRequest(http.MethodGet, "/", nil)
	// 	res := httptest.NewRecorder()

	// 	req.Header.Set("Content-Type", "application/json")

	// 	context := ec.NewContext(req, res)
	// 	context.SetPath("/restaurants/open")

	// 	restaurantCtrl := NewRestaurantsControllers(mockFalseRestaurantRepository{})
	// 	restaurantCtrl.GetsByOpen()(context)

	// 	responses := RestaurantResponseFormat{}
	// 	json.Unmarshal([]byte(res.Body.Bytes()), &responses)

	// 	assert.Equal(t, 404, responses.Code)
	// 	assert.Equal(t, "Not Found", responses.Message)
	// })

	t.Run("Not Found 222 GetByOpen Restaurants", func(t *testing.T) {

		query := make(url.Values)
		query.Set("date_time", "2022-03-05 08:00:00")

		req := httptest.NewRequest(http.MethodGet, "/?"+query.Encode(), nil)
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")

		context := ec.NewContext(req, res)
		context.SetPath("/restaurants/open")

		restaurantCtrl := NewRestaurantsControllers(mockFalseRestaurantRepository{})
		restaurantCtrl.GetsByOpen()(context)

		responses := RestaurantResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, 404, responses.Code)
		assert.Equal(t, "Not Found", responses.Message)
	})

	t.Run("Success GetByOpen Restaurants", func(t *testing.T) {

		query := make(url.Values)
		query.Set("date_time", "2022-02-04 11:00:00")

		req := httptest.NewRequest(http.MethodGet, "/?"+query.Encode(), nil)
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")

		context := ec.NewContext(req, res)
		context.SetPath("/restaurants/open")

		restaurantCtrl := NewRestaurantsControllers(mockRestaurantRepository{})
		restaurantCtrl.GetsByOpen()(context)

		responses := RestaurantResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, 200, responses.Code)
		assert.Equal(t, "Successful Operation", responses.Message)
	})

}

func Test_GetRestoByID_Restaurant(t *testing.T) {
	config := configs.GetConfig()
	fmt.Println(config)

	ec := echo.New()

	t.Run("Not Found GetRestoByID Restaurants", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenRestaurant))

		context := ec.NewContext(req, res)
		context.SetPath("/restaurant")

		restaurantCtrl := NewRestaurantsControllers(mockFalseRestaurantRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(restaurantCtrl.GetRestoByIdCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}

		responses := RestaurantResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, 404, responses.Code)
		assert.Equal(t, "Not Found", responses.Message)
	})

	t.Run("Success GetRestoByID Restaurants", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenRestaurant))

		context := ec.NewContext(req, res)
		context.SetPath("/restaurant")

		restaurantCtrl := NewRestaurantsControllers(mockRestaurantRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(restaurantCtrl.GetRestoByIdCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}

		responses := RestaurantResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, 200, responses.Code)
		assert.Equal(t, "Successful Operation", responses.Message)
	})

}

func Test_UpdateRestoByID_Restaurant(t *testing.T) {
	config := configs.GetConfig()
	fmt.Println(config)

	ec := echo.New()

	t.Run("Bad Request UpdateRestoByID Restaurants", func(t *testing.T) {

		reqBody, _ := json.Marshal(map[string]interface{}{
			"email": 1,
		})

		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenRestaurant))

		context := ec.NewContext(req, res)
		context.SetPath("/restaurant")

		restaurantCtrl := NewRestaurantsControllers(mockFalseRestaurantRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(restaurantCtrl.UpdateRestoByIdCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}

		responses := RestaurantResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, 400, responses.Code)
		assert.Equal(t, "Bad Request", responses.Message)
	})

	t.Run("Not Found UpdateRestoByID Restaurants", func(t *testing.T) {

		reqBody, _ := json.Marshal(map[string]interface{}{
			"email": "adminupdate@outlook.my",
		})

		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenRestaurant))

		context := ec.NewContext(req, res)
		context.SetPath("/restaurant")

		restaurantCtrl := NewRestaurantsControllers(mockFalseRestaurantRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(restaurantCtrl.UpdateRestoByIdCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}

		responses := RestaurantResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, 404, responses.Code)
		assert.Equal(t, "Not Found", responses.Message)
	})

	t.Run("Success UpdateRestoByID Restaurants", func(t *testing.T) {

		reqBody, _ := json.Marshal(map[string]interface{}{
			"email":    "adminupdate@outlook.my",
			"password": "admin123",
		})

		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenRestaurant))

		context := ec.NewContext(req, res)
		context.SetPath("/restaurant")

		restaurantCtrl := NewRestaurantsControllers(mockRestaurantRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(restaurantCtrl.UpdateRestoByIdCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}

		responses := RestaurantResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, 200, responses.Code)
		assert.Equal(t, "Successful Operation", responses.Message)
	})

}

func Test_CreateDetailRestoByID_Restaurant(t *testing.T) {
	config := configs.GetConfig()
	fmt.Println(config)

	ec := echo.New()

	t.Run("Bad Request CreateDetailRestoByID Restaurants", func(t *testing.T) {

		reqBody, _ := json.Marshal(map[string]int{
			"name": 1,
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenRestaurant))

		context := ec.NewContext(req, res)
		context.SetPath("/restaurant/detail")

		restaurantCtrl := NewRestaurantsControllers(mockFalseRestaurantRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(restaurantCtrl.CreateDetailRestoByIdCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}

		responses := RestaurantResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, 400, responses.Code)
		assert.Equal(t, "Bad Request", responses.Message)
	})

	t.Run("Not Found CreateDetailRestoByID Restaurants", func(t *testing.T) {

		reqBody, _ := json.Marshal(map[string]interface{}{
			"name": "Restaurant 1",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenRestaurant))

		context := ec.NewContext(req, res)
		context.SetPath("/restaurant/detail")

		restaurantCtrl := NewRestaurantsControllers(mockFalseRestaurantRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(restaurantCtrl.CreateDetailRestoByIdCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}

		responses := RestaurantResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, 404, responses.Code)
		assert.Equal(t, "Not Found", responses.Message)
	})

	t.Run("Success UpdateRestoByID Restaurants", func(t *testing.T) {

		reqBody, _ := json.Marshal(map[string]interface{}{
			"name": "Restaurant 1",
		})

		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenRestaurant))

		context := ec.NewContext(req, res)
		context.SetPath("/restaurant/detail")

		restaurantCtrl := NewRestaurantsControllers(mockRestaurantRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(restaurantCtrl.CreateDetailRestoByIdCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}

		responses := RestaurantResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, 200, responses.Code)
		assert.Equal(t, "Successful Operation", responses.Message)
	})

}

func Test_UpdateDetailRestoByID_Restaurant(t *testing.T) {
	config := configs.GetConfig()
	fmt.Println(config)

	ec := echo.New()

	t.Run("Bad Request UpdateDetailRestoByID Restaurants", func(t *testing.T) {

		reqBody, _ := json.Marshal(map[string]int{
			"open_hour": 1,
		})

		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenRestaurant))

		context := ec.NewContext(req, res)
		context.SetPath("/restaurant/detail")

		restaurantCtrl := NewRestaurantsControllers(mockFalseRestaurantRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(restaurantCtrl.UpdateDetailRestoByIdCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}

		responses := RestaurantResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, 400, responses.Code)
		assert.Equal(t, "Bad Request", responses.Message)
	})

	t.Run("Not Found UpdateDetailRestoByID Restaurants", func(t *testing.T) {

		reqBody, _ := json.Marshal(map[string]interface{}{
			"open_hour": "11:30",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenRestaurant))

		context := ec.NewContext(req, res)
		context.SetPath("/restaurant/detail")

		restaurantCtrl := NewRestaurantsControllers(mockFalseRestaurantRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(restaurantCtrl.UpdateDetailRestoByIdCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}

		responses := RestaurantResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, 404, responses.Code)
		assert.Equal(t, "Not Found", responses.Message)
	})

	t.Run("Success UpdateDetailRestoByID Restaurants", func(t *testing.T) {

		reqBody, _ := json.Marshal(map[string]interface{}{
			"Open_hour": "12:00",
		})

		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenRestaurant))

		context := ec.NewContext(req, res)
		context.SetPath("/restaurant/detail")

		restaurantCtrl := NewRestaurantsControllers(mockRestaurantRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(restaurantCtrl.UpdateDetailRestoByIdCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}

		responses := RestaurantResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, 200, responses.Code)
		assert.Equal(t, "Successful Operation", responses.Message)
	})

}

func Test_DeleteDetailRestoByID_Restaurant(t *testing.T) {
	config := configs.GetConfig()
	fmt.Println(config)

	ec := echo.New()

	t.Run("Bad Request DeleteDetailRestoByID Restaurants", func(t *testing.T) {

		req := httptest.NewRequest(http.MethodDelete, "/", nil)
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenRestaurant))

		context := ec.NewContext(req, res)
		context.SetPath("/restaurant")

		restaurantCtrl := NewRestaurantsControllers(mockFalseRestaurantRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(restaurantCtrl.DeleteRestaurantCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}

		responses := RestaurantResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, 406, responses.Code)
		assert.Equal(t, "Not Accepted", responses.Message)
	})

	t.Run("Bad Request DeleteDetailRestoByID Restaurants", func(t *testing.T) {

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
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(restaurantCtrl.DeleteRestaurantCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}

		responses := RestaurantResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, 400, responses.Code)
		assert.Equal(t, "Bad Request", responses.Message)
	})

	t.Run("Not Found DeleteDetailRestoByID Restaurants", func(t *testing.T) {

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
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(restaurantCtrl.DeleteRestaurantCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}

		responses := RestaurantResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, 404, responses.Code)
		assert.Equal(t, "Not Found", responses.Message)
	})

	t.Run("Success DeleteDetailRestoByID Restaurants", func(t *testing.T) {

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
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(restaurantCtrl.DeleteRestaurantCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}

		responses := RestaurantResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, 200, responses.Code)
		assert.Equal(t, "Successful Operation", responses.Message)
	})

}

type mockUserRepository struct{}

func (m mockUserRepository) RegisterAdmin(newUser entities.User) (entities.User, error) {
	return entities.User{
		Model: gorm.Model{
			ID:        1,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: gorm.DeletedAt{},
		},
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
		Model: gorm.Model{
			ID:        1,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: gorm.DeletedAt{},
		},
		ID:          1,
		Email:       "admin@outlook.my",
		Password:    "admin123",
		Name:        "admin",
		PhoneNumber: "0877",
		Reputation:  999,
		Balance:     100,
	}, nil
}

func (m mockUserRepository) LoginUser(email, password string) (entities.User, error) {
	hash := sha256.Sum256([]byte("admin123"))
	passwordS := fmt.Sprintf("%x", hash[:])
	return entities.User{
		Model: gorm.Model{
			ID:        1,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: gorm.DeletedAt{},
		},
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
		Model: gorm.Model{
			ID:        1,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: gorm.DeletedAt{},
		},
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
		Model: gorm.Model{
			ID:        1,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: gorm.DeletedAt{},
		},
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
		Model: gorm.Model{
			ID:        1,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: gorm.DeletedAt{},
		},
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
	return entities.Restaurant{ID: 1, Email: "restaurant1@outlook.my"}, nil
}

func (m mockRestaurantRepository) LoginRestaurant(email, password string) (entities.Restaurant, error) {
	hash := sha256.Sum256([]byte("resto123"))
	passwordS := fmt.Sprintf("%x", hash[:])
	return entities.Restaurant{ID: 1, Email: "restaurant1@outlook.my", Password: passwordS}, nil
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
	return []entities.RestaurantDetail{{
		Model: gorm.Model{
			ID:        1,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: gorm.DeletedAt{},
		},
		ID:             1,
		Name:           "Restaurant 1 TRUE",
		Open_Hour:      "15:00",
		Close_Hour:     "17:00",
		Open:           "Monday",
		Close:          "Tuesday,Wednesday,Thursday,Friday,Saturday,Sunday",
		Price:          10000,
		Latitude:       0,
		Longitude:      0,
		City:           "Jakarta",
		Address:        "Jl.Taman Daan Mogot no.7",
		PhoneNumber:    "088",
		ProfilePicture: "https://",
		Seats:          20,
		Status:         "OPEN",
		Description:    "Restaurant Satu",
	}}, nil
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

type mockFalseRestaurantRepository struct{}

func (m mockFalseRestaurantRepository) Register(newUser entities.Restaurant) (entities.Restaurant, error) {
	return entities.Restaurant{ID: 0, Email: "restaurant1@outlook.my"}, errors.New("")
}

func (m mockFalseRestaurantRepository) LoginRestaurant(email, password string) (entities.Restaurant, error) {
	hash := sha256.Sum256([]byte("resto123"))
	passwordS := fmt.Sprintf("%x", hash[:])
	return entities.Restaurant{ID: 0, Email: "restaurant1@outlook.my", Password: passwordS}, errors.New("")
}

func (m mockFalseRestaurantRepository) GetsWaiting() ([]entities.RestaurantDetail, error) {
	return []entities.RestaurantDetail{}, errors.New("")
}

func (m mockFalseRestaurantRepository) Approve(restaurantID uint, status string) (entities.RestaurantDetail, error) {
	return entities.RestaurantDetail{}, errors.New("")
}

func (m mockFalseRestaurantRepository) Get(restaurantID uint) (entities.Restaurant, entities.RestaurantDetail, error) {
	return entities.Restaurant{ID: 0}, entities.RestaurantDetail{ID: 0}, errors.New("")
}

func (m mockFalseRestaurantRepository) GetsByOpen(open int) ([]entities.RestaurantDetail, error) {
	return []entities.RestaurantDetail{}, errors.New("")
}

func (m mockFalseRestaurantRepository) GetExistSeat(restauranId uint, date_time string) ([]entities.Transaction, int, error) {
	return []entities.Transaction{{ID: 0}}, 1, errors.New("")
}

func (m mockFalseRestaurantRepository) Gets() ([]entities.RestaurantDetail, error) {
	return []entities.RestaurantDetail{{ID: 0}}, errors.New("")
}

func (m mockFalseRestaurantRepository) Update(restaurantID uint, updateUser entities.Restaurant) (entities.Restaurant, error) {
	return entities.Restaurant{ID: 0, Email: "restaurant1Update@outlook.my"}, errors.New("")
}

func (m mockFalseRestaurantRepository) UpdateDetail(restaurantID uint, updateUser entities.RestaurantDetail) (entities.RestaurantDetail, error) {
	return entities.RestaurantDetail{
		Model: gorm.Model{
			ID:        1,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: gorm.DeletedAt{},
		},
		ID:             1,
		Name:           "Restauran 1",
		Open_Hour:      "11:30",
		Close_Hour:     "17:00",
		Open:           "Monday",
		Close:          "Tuesday,Wednesday,Thursday,Friday,Saturday,Sunday",
		Price:          10000,
		Latitude:       0,
		Longitude:      0,
		City:           "",
		Address:        "",
		PhoneNumber:    "",
		ProfilePicture: "",
		Seats:          0,
		Status:         "",
		Description:    "",
	}, errors.New("")
}

func (m mockFalseRestaurantRepository) Delete(restaurantID uint) (entities.Restaurant, error) {
	return entities.Restaurant{ID: 0}, errors.New("")
}
