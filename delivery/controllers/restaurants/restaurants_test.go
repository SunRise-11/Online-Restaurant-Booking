package restaurants

import (
	"Restobook/configs"
	"Restobook/delivery/common"
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
)

var jwtToken string

func TestRestaurant(t *testing.T) {
	config := configs.GetConfig()
	fmt.Println(config)

	ec := echo.New()

	t.Run("Register Restaurant", func(t *testing.T) {
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

	t.Run("Login Restaurant", func(t *testing.T) {
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
		jwtToken = responses.Token

		assert.Equal(t, 200, responses.Code)
		assert.Equal(t, "Successful Operation", responses.Message)
	})

	t.Run("Get Restaurant", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

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

	t.Run("Update Restaurant", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]string{
			"email":    "UPDATErestaurant@outlook.my",
			"password": "resto123",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

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

	t.Run("Create Detail Restaurant", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]string{
			"name": "resto 1",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

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

	t.Run("Update Detail Restaurant", func(t *testing.T) {
		now := time.Now()
		reqBody, _ := json.Marshal(map[string]interface{}{
			"open":  now,
			"close": now,
		})

		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

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

type mockRestaurantRepository struct{}

func (m mockRestaurantRepository) Register(newUser entities.Restaurant) (entities.Restaurant, error) {
	return entities.Restaurant{ID: 1, Email: "restaurant1@outlook.my"}, nil
}

func (m mockRestaurantRepository) LoginRestaurant(email, password string) (entities.Restaurant, error) {
	hash := sha256.Sum256([]byte("resto123"))
	passwordS := fmt.Sprintf("%x", hash[:])
	return entities.Restaurant{ID: 1, Email: "restaurant1@outlook.my", Password: passwordS}, nil
}

func (m mockRestaurantRepository) Get(restaurantID uint) (entities.Restaurant, entities.RestaurantDetail, error) {
	return entities.Restaurant{ID: 1}, entities.RestaurantDetail{ID: 1}, nil
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

func TestFalseRestaurant(t *testing.T) {
	config := configs.GetConfig()
	fmt.Println(config)

	ec := echo.New()

	t.Run("Register Restaurant", func(t *testing.T) {
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

	t.Run("FALSE Register Restaurant", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]int{
			"email": 1,
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

		assert.Equal(t, 400, responses.Code)
		assert.Equal(t, "Bad Request", responses.Message)
	})

	t.Run("FALSE Register Restaurant", func(t *testing.T) {
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

	t.Run("FALSE Login Restaurant", func(t *testing.T) {
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
		jwtToken = responses.Token

		assert.Equal(t, 400, responses.Code)
		assert.Equal(t, "Bad Request", responses.Message)
	})

	t.Run("FALSE Login Restaurant", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]string{
			"email":    "restaurant1@outlook.com",
			"password": "resto",
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
		jwtToken = responses.Token

		assert.Equal(t, 404, responses.Code)
		assert.Equal(t, "Not Found", responses.Message)
	})

	t.Run("Login Restaurant", func(t *testing.T) {
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
		jwtToken = responses.Token

		assert.Equal(t, 200, responses.Code)
		assert.Equal(t, "Successful Operation", responses.Message)
	})

	t.Run("FALSE Get Restaurant", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

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

	t.Run("FALSE Update Restaurant", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]int{
			"email": 1,
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

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

	t.Run("FALSE Update Restaurant", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]string{
			"email":    "restaurant1@outlook.my",
			"password": "resto12",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

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

	t.Run("FALSE Create Detail Restaurant", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]int{
			"name": 1,
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := ec.NewContext(req, res)
		context.SetPath("/restaurant/detail")

		restaurantCtrl := NewRestaurantsControllers(mockRestaurantRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(restaurantCtrl.CreateDetailRestoByIdCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}

		responses := RestaurantResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, 400, responses.Code)
		assert.Equal(t, "Bad Request", responses.Message)
	})

	t.Run("FALSE Create Detail Restaurant", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]string{
			"name": "restoFalse",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

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

	t.Run("FALSE Update Detail Restaurant", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]int{
			"open": 1,
		})

		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

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

	t.Run("FALSE Update Detail Restaurant", func(t *testing.T) {
		now := time.Now()
		reqBody, _ := json.Marshal(map[string]interface{}{
			"open": now,
		})

		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

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

func (m mockFalseRestaurantRepository) Get(restaurantID uint) (entities.Restaurant, entities.RestaurantDetail, error) {
	return entities.Restaurant{ID: 0}, entities.RestaurantDetail{ID: 0}, errors.New("")
}

func (m mockFalseRestaurantRepository) Update(restaurantID uint, updateUser entities.Restaurant) (entities.Restaurant, error) {
	return entities.Restaurant{ID: 0, Email: "restaurant1Update@outlook.my"}, errors.New("")
}

func (m mockFalseRestaurantRepository) UpdateDetail(restaurantID uint, updateUser entities.RestaurantDetail) (entities.RestaurantDetail, error) {
	return entities.RestaurantDetail{ID: 0}, errors.New("")
}

func (m mockFalseRestaurantRepository) Delete(restaurantID uint) (entities.Restaurant, error) {
	return entities.Restaurant{ID: 0}, errors.New("")
}
