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
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/stretchr/testify/assert"
)

var jwtToken string

func TestUser(t *testing.T) {
	config := configs.GetConfig()
	fmt.Println(config)

	ec := echo.New()

	t.Run("Register User", func(t *testing.T) {
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

	t.Run("Login User", func(t *testing.T) {
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

	t.Run("Get User", func(t *testing.T) {
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

	t.Run("Update User", func(t *testing.T) {
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

	t.Run("Delete User", func(t *testing.T) {
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

func TestFalseUser(t *testing.T) {
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

	t.Run("Bad Request Login User", func(t *testing.T) {
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

	t.Run("Not Found Login User", func(t *testing.T) {
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

	t.Run("JWT Login User", func(t *testing.T) {
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

	t.Run("Not Found Get User", func(t *testing.T) {
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

	t.Run("Bad Request Update User", func(t *testing.T) {
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

	t.Run("Not Found Update User", func(t *testing.T) {
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

	t.Run("Not Found Delete User", func(t *testing.T) {
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

}

type mockFalseUserRepository struct{}

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
