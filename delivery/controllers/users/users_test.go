package users

import (
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
	"gorm.io/gorm"
)

var jwtTokenUser string

func Test_Resister_User(t *testing.T) {

	ec := echo.New()

	t.Run("400 Register User", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]interface{}{
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
			"name":         "user",
			"email":        "user@outlook.my",
			"password":     "user123",
			"phone_number": "0877",
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
			"name":     "user",
			"email":    "user@outlook.my",
			"password": "user123",
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

func Test_Login_User(t *testing.T) {

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
			"email":    "user@outlook.my",
			"password": "user123",
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
			"email":    "user@outlook.my",
			"password": "user123",
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
		jwtTokenUser = responses.Token

		assert.Equal(t, 200, responses.Code)
		assert.Equal(t, "Successful Operation", responses.Message)
	})

}

func Test_Get_User(t *testing.T) {

	ec := echo.New()

	t.Run("404 Get User", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/", nil)
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))

		context := ec.NewContext(req, res)
		context.SetPath("/users")

		userCtrl := NewUsersControllers(mockFalseUserRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(userCtrl.GetUserCtrl())(context); err != nil {
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
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))

		context := ec.NewContext(req, res)
		context.SetPath("/users")

		userCtrl := NewUsersControllers(mockUserRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(userCtrl.GetUserCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}

		responses := UserResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, 200, responses.Code)
		assert.Equal(t, "Successful Operation", responses.Message)
	})

}

func Test_Update_User(t *testing.T) {

	ec := echo.New()

	t.Run("400 Update User", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]int{
			"name": 1,
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))

		context := ec.NewContext(req, res)
		context.SetPath("/user")

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
			"name":     "user",
			"email":    "user@outlook.my",
			"password": "user123",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))

		context := ec.NewContext(req, res)
		context.SetPath("/user")

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
			"name":         "herlianto",
			"email":        "herlianto@outlook.my",
			"password":     "herlianto123",
			"phone_number": "0877",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))

		context := ec.NewContext(req, res)
		context.SetPath("/user")

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

func Test_Delete_User(t *testing.T) {

	ec := echo.New()

	t.Run("404 Delete User", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/", nil)
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))

		context := ec.NewContext(req, res)
		context.SetPath("/user")

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
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))

		context := ec.NewContext(req, res)
		context.SetPath("/user")

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
	return entities.User{}, nil
}

func (m mockUserRepository) Register(newUser entities.User) (entities.User, error) {
	hash := sha256.Sum256([]byte(newUser.Password))
	passwordS := fmt.Sprintf("%x", hash[:])
	return entities.User{
		Model:       gorm.Model{},
		ID:          1,
		Email:       newUser.Email,
		Password:    passwordS,
		Name:        newUser.Name,
		PhoneNumber: newUser.PhoneNumber,
		Reputation:  80,
		Balance:     0,
	}, nil
}

func (m mockUserRepository) LoginUser(email, password string) (entities.User, error) {
	return entities.User{
		Model:       gorm.Model{},
		ID:          1,
		Email:       email,
		Name:        "user",
		PhoneNumber: "0877",
		Reputation:  80,
		Balance:     0,
	}, nil
}

func (m mockUserRepository) Get(userID uint) (entities.User, error) {
	return entities.User{
		Model:       gorm.Model{},
		ID:          userID,
		Email:       "user@outlook.my",
		Name:        "user",
		PhoneNumber: "0877",
		Reputation:  80,
		Balance:     0,
	}, nil
}

func (m mockUserRepository) Update(userID uint, updateUser entities.User) (entities.User, error) {
	hash := sha256.Sum256([]byte(updateUser.Password))
	passwordS := fmt.Sprintf("%x", hash[:])
	return entities.User{
		Model:       gorm.Model{},
		ID:          userID,
		Email:       updateUser.Email,
		Password:    passwordS,
		Name:        updateUser.Name,
		PhoneNumber: updateUser.PhoneNumber,
		Reputation:  80,
		Balance:     0,
	}, nil
}

func (m mockUserRepository) Delete(userID uint) (entities.User, error) {
	return entities.User{
		Model:       gorm.Model{},
		ID:          userID,
		Email:       "user@outlook.my",
		Name:        "user",
		PhoneNumber: "0877",
		Reputation:  80,
		Balance:     0,
	}, nil
}

type mockFalseUserRepository struct{}

func (m mockFalseUserRepository) RegisterAdmin(newUser entities.User) (entities.User, error) {
	return entities.User{}, errors.New("FAILED REGISTER ADMIN")
}

func (m mockFalseUserRepository) Register(newUser entities.User) (entities.User, error) {
	return entities.User{}, errors.New("FAILED REGISTER USER")
}

func (m mockFalseUserRepository) LoginUser(email, password string) (entities.User, error) {
	return entities.User{}, errors.New("FAILED LOGIN")
}

func (m mockFalseUserRepository) Get(userID uint) (entities.User, error) {
	return entities.User{}, errors.New("FAILED GET")
}

func (m mockFalseUserRepository) Update(userID uint, updateUser entities.User) (entities.User, error) {
	return entities.User{}, errors.New("FAILED UPDATE")
}

func (m mockFalseUserRepository) Delete(userID uint) (entities.User, error) {
	return entities.User{}, errors.New("FAILED DELETE")
}
