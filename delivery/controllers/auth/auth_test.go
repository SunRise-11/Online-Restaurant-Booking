package auth

import (
	"Restobook/configs"
	"Restobook/entities"
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var jwtTokenAdmin, jwtTokenUser, jwtTokenRestaurant string

func TestToken(t *testing.T) {
	t.Run("Create Token Admin", func(t *testing.T) {
		res, _ := CreateTokenAuthAdmin(1)
		jwtTokenAdmin = res
		assert.Equal(t, jwtTokenAdmin, res)
	})
	t.Run("Create Token User", func(t *testing.T) {
		res, _ := CreateTokenAuthUser(1)
		jwtTokenUser = res
		assert.Equal(t, jwtTokenUser, res)
	})
	t.Run("Create Token Restaurant", func(t *testing.T) {
		res, _ := CreateTokenAuthRestaurant(1)
		jwtTokenRestaurant = res
		assert.Equal(t, jwtTokenRestaurant, res)
	})
}

func TestAdmin(t *testing.T) {
	config := configs.GetConfig()
	fmt.Println(config)

	ec := echo.New()

	t.Run("Register Admin", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]string{
			"name":     "Admin",
			"email":    "admin@outlook.my",
			"password": "admin123",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")

		context := ec.NewContext(req, res)
		context.SetPath("/users/register")

		authCtrl := NewAdminControllers(mockAdminRepository{})
		authCtrl.RegisterAdminCtrl()(context)

		responses := AdminResponseFormat{}
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
		context.SetPath("/users/login")

		authCtrl := NewAdminControllers(mockAdminRepository{})
		authCtrl.LoginAuthCtrl()(context)

		responses := LoginResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, 200, responses.Code)
		assert.Equal(t, "Successful Operation", responses.Message)
	})

}

type mockAdminRepository struct{}

func (m mockAdminRepository) RegisterAdmin(newAdmin entities.User) (entities.User, error) {
	return entities.User{ID: 1, Name: "Admin"}, nil
}

func (m mockAdminRepository) Register(newUser entities.User) (entities.User, error) {
	return entities.User{ID: 1, Name: "Admin"}, nil
}

func (m mockAdminRepository) LoginUser(email, password string) (entities.User, error) {
	hash := sha256.Sum256([]byte("admin123"))
	passwordS := fmt.Sprintf("%x", hash[:])
	return entities.User{ID: 1, Name: "Admin", Password: passwordS, Email: "admin@outlook.my"}, nil
}

func (m mockAdminRepository) Get(userID uint) (entities.User, error) {
	return entities.User{ID: 1, Name: "admin"}, nil
}

func (m mockAdminRepository) Update(userID uint, updateUser entities.User) (entities.User, error) {
	return entities.User{ID: 1, Name: "andrew"}, nil
}

func (m mockAdminRepository) Delete(userID uint) (entities.User, error) {
	return entities.User{ID: 0}, nil
}

func TestFalseAdmin(t *testing.T) {
	config := configs.GetConfig()
	fmt.Println(config)

	ec := echo.New()

	t.Run("FALSE Register Admin", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]int{
			"name": 1,
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")

		context := ec.NewContext(req, res)
		context.SetPath("/admin/register")

		authCtrl := NewAdminControllers(mockFalseAdminRepository{})
		authCtrl.RegisterAdminCtrl()(context)

		responses := AdminResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, 400, responses.Code)
		assert.Equal(t, "Bad Request", responses.Message)
	})
	t.Run("FALSE Register Admin", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]string{
			"name":     "Admin",
			"email":    "admin@outlook.my",
			"password": "admin123",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")

		context := ec.NewContext(req, res)
		context.SetPath("/admin/register")

		authCtrl := NewAdminControllers(mockFalseAdminRepository{})
		authCtrl.RegisterAdminCtrl()(context)

		responses := AdminResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, 500, responses.Code)
		assert.Equal(t, "Internal Server Error", responses.Message)
	})

	t.Run("FALSE Login Admin", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]int{
			"email": 1,
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")

		context := ec.NewContext(req, res)
		context.SetPath("/admin/login")

		authCtrl := NewAdminControllers(mockFalseAdminRepository{})
		authCtrl.LoginAuthCtrl()(context)

		responses := LoginResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, 400, responses.Code)
		assert.Equal(t, "Bad Request", responses.Message)
	})

	t.Run("FALSE Login Admin", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]string{
			"email":    "admisn@outlook.my",
			"password": "admin123",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")

		context := ec.NewContext(req, res)
		context.SetPath("/admin/login")

		authCtrl := NewAdminControllers(mockFalseAdminRepository{})
		authCtrl.LoginAuthCtrl()(context)

		responses := LoginResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, 404, responses.Code)
		assert.Equal(t, "Not Found", responses.Message)
	})

}

type mockFalseAdminRepository struct{}

func (m mockFalseAdminRepository) RegisterAdmin(newAdmin entities.User) (entities.User, error) {
	return entities.User{ID: 0, Name: "admin"}, errors.New("")
}

func (m mockFalseAdminRepository) Register(newUser entities.User) (entities.User, error) {
	return entities.User{ID: 1, Name: "admin"}, errors.New("")
}

func (m mockFalseAdminRepository) LoginUser(email, password string) (entities.User, error) {
	hash := sha256.Sum256([]byte("admin123"))
	passwordS := fmt.Sprintf("%x", hash[:])
	return entities.User{ID: 0, Name: "admin", Password: passwordS, Email: "admin@outlook.my"}, errors.New("")
}

func (m mockFalseAdminRepository) Get(userID uint) (entities.User, error) {
	return entities.User{ID: 1, Name: "admin"}, errors.New("")
}

func (m mockFalseAdminRepository) Update(userID uint, updateUser entities.User) (entities.User, error) {
	return entities.User{ID: 1, Name: "andrew"}, errors.New("")
}

func (m mockFalseAdminRepository) Delete(userID uint) (entities.User, error) {
	return entities.User{ID: 0}, errors.New("")
}
