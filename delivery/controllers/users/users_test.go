package users

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

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/stretchr/testify/assert"
)

var jwtToken string

func TestResisterUser(t *testing.T) {
	config := configs.GetConfig()
	fmt.Println(config)

	ec := echo.New()

	t.Run("400 Register User", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]int{
			"name": 1,
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")

		context := ec.NewContext(req, res)
		context.SetPath("/users/register")

		userCtrl := NewUsersControllers(mockFalseUserRepository{})
		userCtrl.RegisterUserCtrl()(context)

		responses := UserResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, 400, responses.Code)
		assert.Equal(t, "Bad Request", responses.Message)
	})

	t.Run("500 Register User", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]string{
			"name":     "herlianto",
			"email":    "herlianto@outlook.my",
			"password": "herlianto123",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")

		context := ec.NewContext(req, res)
		context.SetPath("/users/register")

		userCtrl := NewUsersControllers(mockFalseUserRepository{})
		userCtrl.RegisterUserCtrl()(context)

		responses := UserResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, 500, responses.Code)
		assert.Equal(t, "Internal Server Error", responses.Message)
	})

	t.Run("200 Register User", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]string{
			"name":     "herlianto",
			"email":    "herlianto@outlook.my",
			"password": "herlianto123",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")

		context := ec.NewContext(req, res)
		context.SetPath("/users/register")

		userCtrl := NewUsersControllers(mockUserRepository{})
		userCtrl.RegisterUserCtrl()(context)

		responses := UserResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, 200, responses.Code)
		assert.Equal(t, "Successful Operation", responses.Message)
	})

}

func TestLoginUser(t *testing.T) {
	config := configs.GetConfig()
	fmt.Println(config)
	ec := echo.New()

	t.Run("400 Login User", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]int{
			"email": 1,
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")

		context := ec.NewContext(req, res)
		context.SetPath("/users/login")

		userCtrl := NewUsersControllers(mockFalseUserRepository{})
		userCtrl.LoginAuthCtrl()(context)

		responses := LoginResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, 400, responses.Code)
		assert.Equal(t, "Bad Request", responses.Message)
	})

	t.Run("404 Login User", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]string{
			"email":    "herlianto@outlook.my",
			"password": "herlianto1232",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")

		context := ec.NewContext(req, res)
		context.SetPath("/users/login")

		userCtrl := NewUsersControllers(mockFalseUserRepository{})
		userCtrl.LoginAuthCtrl()(context)

		responses := LoginResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, 404, responses.Code)
		assert.Equal(t, "Not Found", responses.Message)
	})

	t.Run("200 Login User", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]string{
			"email":    "herlianto@outlook.my",
			"password": "herlianto123",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")

		context := ec.NewContext(req, res)
		context.SetPath("/users/login")

		userCtrl := NewUsersControllers(mockUserRepository{})
		userCtrl.LoginAuthCtrl()(context)

		responses := LoginResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)
		jwtToken = responses.Token

		assert.Equal(t, 200, responses.Code)
		assert.Equal(t, "Successful Operation", responses.Message)
	})

}

func TestGetUserByID(t *testing.T) {
	config := configs.GetConfig()
	fmt.Println(config)
	ec := echo.New()

	t.Run("404 Get User", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/", nil)
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := ec.NewContext(req, res)
		context.SetPath("/users")

		userCtrl := NewUsersControllers(mockFalseUserRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(userCtrl.GetUserByIdCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}

		responses := UserResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, 404, responses.Code)
		assert.Equal(t, "Not Found", responses.Message)
	})

	t.Run("200 Get User", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/", nil)
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := ec.NewContext(req, res)
		context.SetPath("/users")

		userCtrl := NewUsersControllers(mockUserRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(userCtrl.GetUserByIdCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}

		responses := UserResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, 200, responses.Code)
		assert.Equal(t, "Successful Operation", responses.Message)
	})

}

func TestUpdateUser(t *testing.T) {
	config := configs.GetConfig()
	fmt.Println(config)
	ec := echo.New()

	t.Run("400 Update User", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]int{
			"name": 1,
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := ec.NewContext(req, res)
		context.SetPath("/users")

		userCtrl := NewUsersControllers(mockFalseUserRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(userCtrl.UpdateUserCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}

		responses := UserResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, 400, responses.Code)
		assert.Equal(t, "Bad Request", responses.Message)
	})

	t.Run("404 Update User", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]string{
			"name":     "herlianto",
			"email":    "herlianto@outlook.my",
			"password": "herlianto123",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := ec.NewContext(req, res)
		context.SetPath("/users")

		userCtrl := NewUsersControllers(mockFalseUserRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(userCtrl.UpdateUserCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}

		responses := UserResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, 404, responses.Code)
		assert.Equal(t, "Not Found", responses.Message)
	})

	t.Run("200 Update User", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]string{
			"name":     "herlianto",
			"email":    "herlianto@outlook.my",
			"password": "herlianto123",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := ec.NewContext(req, res)
		context.SetPath("/users")

		userCtrl := NewUsersControllers(mockUserRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(userCtrl.UpdateUserCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}

		responses := UserResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, 200, responses.Code)
		assert.Equal(t, "Successful Operation", responses.Message)
	})

}

func TestDeleteUser(t *testing.T) {

	config := configs.GetConfig()
	fmt.Println(config)
	ec := echo.New()

	t.Run("404 Delete User", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/", nil)
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := ec.NewContext(req, res)
		context.SetPath("/users")

		userCtrl := NewUsersControllers(mockFalseUserRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(userCtrl.DeleteUserCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}

		responses := UserResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, 404, responses.Code)
		assert.Equal(t, "Not Found", responses.Message)
	})

	t.Run("200 Delete User", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/", nil)
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := ec.NewContext(req, res)
		context.SetPath("/users")

		userCtrl := NewUsersControllers(mockUserRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(userCtrl.DeleteUserCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}

		responses := UserResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, 200, responses.Code)
		assert.Equal(t, "Successful Operation", responses.Message)
	})

}

type mockUserRepository struct{}

func (m mockUserRepository) RegisterAdmin(newUser entities.User) (entities.User, error) {
	return entities.User{ID: 1, Name: "admin"}, nil
}

func (m mockUserRepository) Register(newUser entities.User) (entities.User, error) {
	return entities.User{ID: 1, Name: "herlianto"}, nil
}

func (m mockUserRepository) LoginUser(email, password string) (entities.User, error) {
	hash := sha256.Sum256([]byte("herlianto123"))
	passwordS := fmt.Sprintf("%x", hash[:])
	return entities.User{ID: 1, Name: "herlianto", Password: passwordS, Email: "herlianto@outlook.my"}, nil
}

func (m mockUserRepository) Get(userID uint) (entities.User, error) {
	return entities.User{ID: 1, Name: "herlianto"}, nil
}

func (m mockUserRepository) Update(userID uint, updateUser entities.User) (entities.User, error) {
	return entities.User{ID: 1, Name: "andrew"}, nil
}

func (m mockUserRepository) Delete(userID uint) (entities.User, error) {
	return entities.User{ID: 0}, nil
}

type mockFalseUserRepository struct{}

func (m mockFalseUserRepository) RegisterAdmin(newUser entities.User) (entities.User, error) {
	return entities.User{ID: 0, Name: "admin"}, errors.New("")
}

func (m mockFalseUserRepository) Register(newUser entities.User) (entities.User, error) {
	return entities.User{ID: 0, Name: "herlianto"}, errors.New("")
}

func (m mockFalseUserRepository) LoginUser(email, password string) (entities.User, error) {
	hash := sha256.Sum256([]byte("herlianto123"))
	passwordS := fmt.Sprintf("%x", hash[:])
	return entities.User{ID: 1, Name: "herlianto", Password: passwordS, Email: "herlianto@outlook.my"}, errors.New("")
}

func (m mockFalseUserRepository) Get(userID uint) (entities.User, error) {
	return entities.User{ID: 0, Name: "herlianto"}, errors.New("")
}

func (m mockFalseUserRepository) Update(userID uint, updateUser entities.User) (entities.User, error) {
	return entities.User{ID: 0, Name: "andrew"}, errors.New("")
}

func (m mockFalseUserRepository) Delete(userID uint) (entities.User, error) {
	return entities.User{ID: 0}, errors.New("")
}
