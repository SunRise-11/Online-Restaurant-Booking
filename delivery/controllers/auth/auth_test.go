package auth

import (
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
	"gorm.io/gorm"
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

func Test_Register_Admin(t *testing.T) {
	ec := echo.New()

	t.Run("400 Register Admin", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]int{
			"name":         1,
			"email":        1,
			"password":     1,
			"phone_number": 1,
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")

		context := ec.NewContext(req, res)
		context.SetPath("/admin/register")

		authCtrl := NewAdminControllers(mockAdminRepository{})
		authCtrl.RegisterAdminCtrl()(context)

		responses := AdminResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, 400, responses.Code)
		assert.Equal(t, "Bad Request", responses.Message)
	})

	t.Run("200 Register Admin", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]string{
			"name":        "Admin",
			"email":       "admin@outlook.my",
			"password":    "admin123",
			"phone_numer": "0877",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")

		context := ec.NewContext(req, res)
		context.SetPath("/admin/register")

		authCtrl := NewAdminControllers(mockAdminRepository{})
		authCtrl.RegisterAdminCtrl()(context)

		responses := AdminResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, 200, responses.Code)
		assert.Equal(t, "Successful Operation", responses.Message)
	})

	t.Run("500 Register Admin", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]string{
			"name":         "Admin",
			"email":        "admin@outlook.my",
			"password":     "admin123",
			"phone_number": "0877",
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

}

func Test_Login_Admin(t *testing.T) {
	ec := echo.New()

	t.Run("400 Login Admin", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]int{
			"email":    1,
			"password": 1,
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")

		context := ec.NewContext(req, res)
		context.SetPath("/admin/login")

		authCtrl := NewAdminControllers(mockAdminRepository{})
		authCtrl.LoginAdminCtrl()(context)

		responses := AdminResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, 400, responses.Code)
		assert.Equal(t, "Bad Request", responses.Message)
	})

	t.Run("404 Login Admin", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]string{
			"email":    "admin@outlook.my",
			"password": "admin123",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")

		context := ec.NewContext(req, res)
		context.SetPath("/admin/register")

		authCtrl := NewAdminControllers(mockFalseAdminRepository{})
		authCtrl.LoginAdminCtrl()(context)

		responses := AdminResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, 404, responses.Code)
		assert.Equal(t, "Not Found", responses.Message)
	})

	t.Run("200 Login Admin", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]string{
			"email":    "admin@outlook.my",
			"password": "admin123",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")

		context := ec.NewContext(req, res)
		context.SetPath("/admin/register")

		authCtrl := NewAdminControllers(mockAdminRepository{})
		authCtrl.LoginAdminCtrl()(context)

		responses := AdminResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, 200, responses.Code)
		assert.Equal(t, "Successful Operation", responses.Message)
	})

}

type mockAdminRepository struct{}

func (m mockAdminRepository) RegisterAdmin(newAdmin entities.User) (entities.User, error) {
	hash := sha256.Sum256([]byte("admin123"))
	passwordS := fmt.Sprintf("%x", hash[:])
	return entities.User{
		Model:      gorm.Model{},
		ID:         1,
		Email:      newAdmin.Email,
		Password:   passwordS,
		Name:       newAdmin.Name,
		Reputation: newAdmin.Reputation,
	}, nil
}

func (m mockAdminRepository) Register(newUser entities.User) (entities.User, error) {
	return entities.User{}, nil
}

func (m mockAdminRepository) LoginUser(email, password string) (entities.User, error) {
	hash := sha256.Sum256([]byte("admin123"))
	passwordS := fmt.Sprintf("%x", hash[:])
	return entities.User{
		Model:       gorm.Model{},
		ID:          1,
		Email:       email,
		Password:    passwordS,
		Name:        "admin",
		PhoneNumber: "0877",
	}, nil
}

func (m mockAdminRepository) Get(userID uint) (entities.User, error) {
	return entities.User{}, nil
}

func (m mockAdminRepository) Update(userID uint, updateUser entities.User) (entities.User, error) {
	return entities.User{}, nil
}

func (m mockAdminRepository) Delete(userID uint) (entities.User, error) {
	return entities.User{}, nil
}

type mockFalseAdminRepository struct{}

func (m mockFalseAdminRepository) RegisterAdmin(newAdmin entities.User) (entities.User, error) {
	return entities.User{}, errors.New("FAILED REGISTER ADMIN")
}

func (m mockFalseAdminRepository) Register(newUser entities.User) (entities.User, error) {
	return entities.User{}, errors.New("FAILED REGISTER USER")
}

func (m mockFalseAdminRepository) LoginUser(email, password string) (entities.User, error) {
	return entities.User{}, errors.New("FAILED LOGIN")
}

func (m mockFalseAdminRepository) Get(userID uint) (entities.User, error) {
	return entities.User{}, errors.New("FAILED GET")
}

func (m mockFalseAdminRepository) Update(userID uint, updateUser entities.User) (entities.User, error) {
	return entities.User{}, errors.New("FAILED UPDATE")
}

func (m mockFalseAdminRepository) Delete(userID uint) (entities.User, error) {
	return entities.User{}, errors.New("FAILED DELETE")
}
