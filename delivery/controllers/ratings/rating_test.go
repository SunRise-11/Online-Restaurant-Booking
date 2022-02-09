package ratings

import (
	"Restobook/delivery/common"
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

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/stretchr/testify/assert"
)

var jwtTokenUser string

func TestCreateRating(t *testing.T) {
	ec := echo.New()

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

		userCtrl := users.NewUsersControllers(mockUserRepository{})
		userCtrl.LoginAuthCtrl()(context)

		responses := users.LoginResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)
		jwtTokenUser = responses.Token
	})

	t.Run("Create Rating", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]interface{}{
			"rating":  1,
			"comment": "Mantap",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))

		context := ec.NewContext(req, res)
		context.SetPath("/ratings")

		ratingCtrl := NewRatingController(mockRatingRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(ratingCtrl.Create())(context); err != nil {
			log.Fatal(err)
			return
		}

		responses := users.LoginResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, 200, responses.Code)
		assert.Equal(t, "Successful Operation", responses.Message)
	})

	t.Run("Create Rating fail but Update Success", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]interface{}{
			"rating":  1,
			"comment": "Mantap",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))

		context := ec.NewContext(req, res)
		context.SetPath("/ratings")

		ratingCtrl := NewRatingController(mockFalseCreateRatingRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(ratingCtrl.Create())(context); err != nil {
			log.Fatal(err)
			return
		}

		responses := users.LoginResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, 200, responses.Code)
		assert.Equal(t, "Successful Operation", responses.Message)
	})
}

func TestUpdateRating(t *testing.T) {
	ec := echo.New()

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

		userCtrl := users.NewUsersControllers(mockUserRepository{})
		userCtrl.LoginAuthCtrl()(context)

		responses := users.LoginResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)
		jwtTokenUser = responses.Token
	})

	t.Run("Update Rating", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]interface{}{
			"rating":  1,
			"comment": "Mantap",
		})

		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))

		context := ec.NewContext(req, res)
		context.SetPath("/ratings")
		context.SetParamNames("restaurantId")
		context.SetParamValues("1")

		ratingCtrl := NewRatingController(mockRatingRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(ratingCtrl.Update())(context); err != nil {
			log.Fatal(err)
			return
		}

		responses := users.LoginResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, 200, responses.Code)
		assert.Equal(t, "Successful Operation", responses.Message)
	})
}

func TestDeleteRating(t *testing.T) {
	ec := echo.New()

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

		userCtrl := users.NewUsersControllers(mockUserRepository{})
		userCtrl.LoginAuthCtrl()(context)

		responses := users.LoginResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)
		jwtTokenUser = responses.Token
	})

	t.Run("Delete Rating", func(t *testing.T) {

		req := httptest.NewRequest(http.MethodDelete, "/", nil)
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))

		context := ec.NewContext(req, res)
		context.SetPath("/ratings")
		context.SetParamNames("restaurantId")
		context.SetParamValues("1")

		ratingCtrl := NewRatingController(mockRatingRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(ratingCtrl.Delete())(context); err != nil {
			log.Fatal(err)
			return
		}

		responses := users.LoginResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, 200, responses.Code)
		assert.Equal(t, "Successful Operation", responses.Message)
	})
}

func TestFalseCreateRating(t *testing.T) {

	ec := echo.New()

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

		userCtrl := users.NewUsersControllers(mockUserRepository{})
		userCtrl.LoginAuthCtrl()(context)

		responses := users.LoginResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)
		jwtTokenUser = responses.Token
	})

	t.Run("FALSE Create Rating", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]int{
			"rating":  1,
			"comment": 1,
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))

		context := ec.NewContext(req, res)
		context.SetPath("/ratings")

		ratingCtrl := NewRatingController(mockFalseRatingRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(ratingCtrl.Create())(context); err != nil {
			log.Fatal(err)
			return
		}

		responses := users.LoginResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, 400, responses.Code)
		assert.Equal(t, "Bad Request", responses.Message)
	})

	t.Run("FALSE Create Rating", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]interface{}{
			"rating":  1,
			"comment": "Mantap",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))

		context := ec.NewContext(req, res)
		context.SetPath("/ratings")

		ratingCtrl := NewRatingController(mockFalseRatingRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(ratingCtrl.Create())(context); err != nil {
			log.Fatal(err)
			return
		}

		responses := users.LoginResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, 404, responses.Code)
		assert.Equal(t, "Not Found", responses.Message)
	})

	t.Run("FALSE Create Rating", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]interface{}{
			"restaurant_id": 2,
			"rating":        1,
			"comment":       "Mantap",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))

		context := ec.NewContext(req, res)
		context.SetPath("/ratings")

		ratingCtrl := NewRatingController(mockFalseRatingRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(ratingCtrl.Create())(context); err != nil {
			log.Fatal(err)
			return
		}

		responses := users.LoginResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, 404, responses.Code)
		assert.Equal(t, "Not Found", responses.Message)
	})
}

func TestFalseUpdateRating(t *testing.T) {

	ec := echo.New()

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

		userCtrl := users.NewUsersControllers(mockUserRepository{})
		userCtrl.LoginAuthCtrl()(context)

		responses := users.LoginResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)
		jwtTokenUser = responses.Token
	})

	t.Run("FALSE Update Rating", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]interface{}{
			"rating":  1,
			"comment": "Mantap",
		})

		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))

		context := ec.NewContext(req, res)
		context.SetPath("/ratings/1")

		ratingCtrl := NewRatingController(mockFalseRatingRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(ratingCtrl.Update())(context); err != nil {
			log.Fatal(err)
			return
		}

		responses := users.LoginResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, 400, responses.Code)
		assert.Equal(t, "Bad Request", responses.Message)
	})

	t.Run("FALSE Update Rating", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]int{
			"rating":  1,
			"comment": 1,
		})

		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))

		context := ec.NewContext(req, res)
		context.SetPath("/ratings")
		context.SetParamNames("restaurantId")
		context.SetParamValues("1")

		ratingCtrl := NewRatingController(mockFalseRatingRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(ratingCtrl.Update())(context); err != nil {
			log.Fatal(err)
			return
		}

		responses := users.LoginResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, 400, responses.Code)
		assert.Equal(t, "Bad Request", responses.Message)
	})

	t.Run("FALSE Update Rating", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]int{
			"rating": 1,
		})

		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))

		context := ec.NewContext(req, res)
		context.SetPath("/ratings")
		context.SetParamNames("restaurantId")
		context.SetParamValues("1")

		ratingCtrl := NewRatingController(mockFalseRatingRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(ratingCtrl.Update())(context); err != nil {
			log.Fatal(err)
			return
		}

		responses := users.LoginResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, 404, responses.Code)
		assert.Equal(t, "Not Found", responses.Message)
	})
}

func TestFalseDeleteRating(t *testing.T) {

	ec := echo.New()

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

		userCtrl := users.NewUsersControllers(mockUserRepository{})
		userCtrl.LoginAuthCtrl()(context)

		responses := users.LoginResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)
		jwtTokenUser = responses.Token
	})

	t.Run("FALSE Delete Rating", func(t *testing.T) {

		req := httptest.NewRequest(http.MethodDelete, "/", nil)
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))

		context := ec.NewContext(req, res)
		context.SetPath("/ratings")
		context.SetParamNames("a")

		ratingCtrl := NewRatingController(mockFalseRatingRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(ratingCtrl.Delete())(context); err != nil {
			log.Fatal(err)
			return
		}

		responses := users.LoginResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, 400, responses.Code)
		assert.Equal(t, "Bad Request", responses.Message)
	})

	t.Run("FALSE Delete Rating", func(t *testing.T) {

		req := httptest.NewRequest(http.MethodDelete, "/", nil)
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))

		context := ec.NewContext(req, res)
		context.SetPath("/ratings")
		context.SetParamNames("restaurantId")
		context.SetParamValues("2")

		ratingCtrl := NewRatingController(mockFalseRatingRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(ratingCtrl.Delete())(context); err != nil {
			log.Fatal(err)
			return
		}

		responses := users.LoginResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, 404, responses.Code)
		assert.Equal(t, "Not Found", responses.Message)
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

type mockRatingRepository struct{}

func (m mockRatingRepository) Create(newRating entities.Rating) (entities.Rating, error) {
	return entities.Rating{ID: 1, Rating: 5, Comment: "Mantap"}, nil
}

func (m mockRatingRepository) Update(updateRating entities.Rating) (entities.Rating, error) {
	return entities.Rating{ID: 1}, nil
}

func (m mockRatingRepository) Delete(userID, restaurantID int) (entities.Rating, error) {
	return entities.Rating{ID: 1}, nil
}

func (m mockRatingRepository) IsCanGiveRating(userID, restaurantID int) (bool, error) {
	return true, nil
}

type mockFalseRatingRepository struct{}

func (m mockFalseRatingRepository) Create(newRating entities.Rating) (entities.Rating, error) {
	return entities.Rating{ID: 0, Rating: 5, Comment: "Mantap"}, errors.New("")
}

func (m mockFalseRatingRepository) Update(updateRating entities.Rating) (entities.Rating, error) {
	return entities.Rating{ID: 0}, errors.New("")
}

func (m mockFalseRatingRepository) Delete(userID, restaurantID int) (entities.Rating, error) {
	return entities.Rating{ID: 0}, errors.New("")
}

func (m mockFalseRatingRepository) IsCanGiveRating(userID, restaurantID int) (bool, error) {
	return false, errors.New("")
}

type mockFalseCreateRatingRepository struct{}

func (m mockFalseCreateRatingRepository) Create(newRating entities.Rating) (entities.Rating, error) {
	return entities.Rating{ID: 0, Rating: 5, Comment: "Mantap"}, errors.New("")
}

func (m mockFalseCreateRatingRepository) Update(updateRating entities.Rating) (entities.Rating, error) {
	return entities.Rating{ID: 0}, nil
}

func (m mockFalseCreateRatingRepository) Delete(userID, restaurantID int) (entities.Rating, error) {
	return entities.Rating{ID: 0}, errors.New("")
}

func (m mockFalseCreateRatingRepository) IsCanGiveRating(userID, restaurantID int) (bool, error) {
	return true, errors.New("")
}
