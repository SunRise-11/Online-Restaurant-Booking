package topup

import (
	"Restobook/configs"
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
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/stretchr/testify/assert"
)

var jwtToken string

func TestTopupXenditError(t *testing.T) {
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
		jwtToken = responses.Token
	})

	t.Run("Topup Internal Server Error Xendit", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]int{
			"total": 10000,
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := ec.NewContext(req, res)
		context.SetPath("/topup")

		topupCtrl := NewTopUpControllers(mockTopupRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(topupCtrl.TopUp())(context); err != nil {
			log.Fatal(err)
			return
		}

		responses := TopUpResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, 500, responses.Code)
		assert.Equal(t, "Internal Server Error", responses.Message)
	})
}
func TestTopup(t *testing.T) {
	configs.GetConfig()
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
		jwtToken = responses.Token
	})

	t.Run("Topup Success", func(t *testing.T) {

		reqBody, _ := json.Marshal(map[string]int{
			"total": 10000,
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := ec.NewContext(req, res)
		context.SetPath("/topup")

		topupCtrl := NewTopUpControllers(mockTopupRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(topupCtrl.TopUp())(context); err != nil {
			log.Fatal(err)
			return
		}

		responses := TopUpResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, 200, responses.Code)
		assert.Equal(t, "Successful Operation", responses.Message)
	})

	t.Run("Topup Bad Request at Binding", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]string{
			"total": "10000",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := ec.NewContext(req, res)
		context.SetPath("/topup")

		topupCtrl := NewTopUpControllers(mockTopupRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(topupCtrl.TopUp())(context); err != nil {
			log.Fatal(err)
			return
		}

		responses := TopUpResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, 400, responses.Code)
		assert.Equal(t, "Bad Request", responses.Message)
	})

	t.Run("Topup Internal Server Error at Create", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]int{
			"total": 10000,
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := ec.NewContext(req, res)
		context.SetPath("/topup")

		topupCtrl := NewTopUpControllers(mockFalseTopupRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(topupCtrl.TopUp())(context); err != nil {
			log.Fatal(err)
			return
		}

		responses := TopUpResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, 500, responses.Code)
		assert.Equal(t, "Internal Server Error", responses.Message)
	})
}

func TestGetAllWaiting(t *testing.T) {
	configs.GetConfig()
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
		jwtToken = responses.Token
	})

	t.Run("Get All Waiting Success ", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/", nil)
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := ec.NewContext(req, res)
		context.SetPath("/topup/pending")

		topupCtrl := NewTopUpControllers(mockTopupRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(topupCtrl.GetAllWaiting())(context); err != nil {
			log.Fatal(err)
			return
		}

		responses := TopUpResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, 200, responses.Code)
		assert.Equal(t, "Successful Operation", responses.Message)
	})

	t.Run("Get All Waiting Not Found Error ", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/", nil)
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := ec.NewContext(req, res)
		context.SetPath("/topup/pending")

		topupCtrl := NewTopUpControllers(mockFalseTopupRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(topupCtrl.GetAllWaiting())(context); err != nil {
			log.Fatal(err)
			return
		}

		responses := TopUpResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, 404, responses.Code)
		assert.Equal(t, "Not Found", responses.Message)
	})
}

func TestGetAllPaid(t *testing.T) {
	configs.GetConfig()
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
		jwtToken = responses.Token
	})

	t.Run("Get All Waiting Success ", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/", nil)
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := ec.NewContext(req, res)
		context.SetPath("/topup/pending")

		topupCtrl := NewTopUpControllers(mockTopupRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(topupCtrl.GetAllPaid())(context); err != nil {
			log.Fatal(err)
			return
		}

		responses := TopUpResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, 200, responses.Code)
		assert.Equal(t, "Successful Operation", responses.Message)
	})

	t.Run("Get All Waiting Not Found Error ", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/", nil)
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := ec.NewContext(req, res)
		context.SetPath("/topup/pending")

		topupCtrl := NewTopUpControllers(mockFalseTopupRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(topupCtrl.GetAllPaid())(context); err != nil {
			log.Fatal(err)
			return
		}

		responses := TopUpResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, 404, responses.Code)
		assert.Equal(t, "Not Found", responses.Message)
	})
}

func TestCallback(t *testing.T) {
	configs.GetConfig()
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
		jwtToken = responses.Token
	})

	t.Run("Callback Success", func(t *testing.T) {

		reqBody, _ := json.Marshal(map[string]string{
			"external_id": "XENDIT INVOICE ID",
			"status":      "PAID",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("X-Callback-Token", common.XENDIT_CALLBACK_TOKEN)

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := ec.NewContext(req, res)
		context.SetPath("/topup/callback")

		topupCtrl := NewTopUpControllers(mockTopupRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(topupCtrl.Callback())(context); err != nil {
			log.Fatal(err)
			return
		}

		responses := TopUpResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, 200, responses.Code)
		assert.Equal(t, "Successful Operation", responses.Message)
	})

	t.Run("Callback Status Not Accepted", func(t *testing.T) {

		reqBody, _ := json.Marshal(map[string]string{
			"external_id": "XENDIT INVOICE ID",
			"status":      "PAID",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := ec.NewContext(req, res)
		context.SetPath("/topup/callback")

		topupCtrl := NewTopUpControllers(mockTopupRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(topupCtrl.Callback())(context); err != nil {
			log.Fatal(err)
			return
		}

		responses := TopUpResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, 406, responses.Code)
		assert.Equal(t, "Not Accepted", responses.Message)
	})

	t.Run("Callback Not Found", func(t *testing.T) {

		reqBody, _ := json.Marshal(map[string]string{
			"external_id": "XENDIT INVOICE ID",
			"status":      "PAID",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("X-Callback-Token", common.XENDIT_CALLBACK_TOKEN)

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := ec.NewContext(req, res)
		context.SetPath("/topup/callback")

		topupCtrl := NewTopUpControllers(mockFalseTopupRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(topupCtrl.Callback())(context); err != nil {
			log.Fatal(err)
			return
		}

		responses := TopUpResponseFormat{}
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

type mockTopupRepository struct{}

func (mt mockTopupRepository) Create(topup entities.TopUp) (entities.TopUp, error) {
	invoiceId := strings.ToUpper(strings.ReplaceAll(uuid.New().String(), "-", ""))
	return entities.TopUp{ID: 1, UserID: 1, InvoiceID: "UNIT TESTING " + invoiceId, PaymentUrl: "", Total: 10000, Status: "PENDING"}, nil
}

func (mt mockTopupRepository) GetAllWaiting(userId uint) ([]entities.TopUp, error) {
	return []entities.TopUp{{ID: 1, UserID: 1, InvoiceID: "XENDIT INVOICE ID", PaymentUrl: "XENDIT PAYMENT URL", Total: 10000, Status: "PENDING"}}, nil
}

func (mt mockTopupRepository) GetAllPaid(userId uint) ([]entities.TopUp, error) {
	return []entities.TopUp{{ID: 1, UserID: 1, InvoiceID: "XENDIT INVOICE ID", PaymentUrl: "XENDIT PAYMENT URL", Total: 10000, Status: "PAID"}}, nil
}

func (mt mockTopupRepository) Update(extId string, topUp entities.TopUp) (entities.TopUp, error) {
	return entities.TopUp{ID: 1, UserID: 1, InvoiceID: "XENDIT INVOICE ID", PaymentUrl: "XENDIT PAYMENT URL", Total: 10000, Status: "PAID"}, nil
}

func (mt mockTopupRepository) GetByInvoice(extId string) (entities.TopUp, error) {
	return entities.TopUp{ID: 1, UserID: 1, InvoiceID: "XENDIT INVOICE ID", PaymentUrl: "XENDIT PAYMENT URL", Total: 10000, Status: "PENDING"}, nil
}

func (mt mockTopupRepository) GetUser(userId int) (entities.User, error) {
	return entities.User{ID: 1, Balance: 10000}, nil
}

func (mt mockTopupRepository) UpdateUserBalance(userId int, user entities.User) (entities.User, error) {
	return entities.User{ID: 1, Balance: 10000}, nil
}

type mockFalseTopupRepository struct{}

func (mt mockFalseTopupRepository) Create(topup entities.TopUp) (entities.TopUp, error) {
	invoiceId := strings.ToUpper(strings.ReplaceAll(uuid.New().String(), "-", ""))
	return entities.TopUp{ID: 1, UserID: 1, InvoiceID: "UNIT TESTING " + invoiceId, PaymentUrl: "", Total: 10000, Status: "PENDING"}, errors.New("")
}

func (mt mockFalseTopupRepository) GetAllWaiting(userId uint) ([]entities.TopUp, error) {
	return []entities.TopUp{{ID: 1, UserID: 1, InvoiceID: "XENDIT INVOICE ID", PaymentUrl: "XENDIT PAYMENT URL", Total: 10000, Status: "PENDING"}}, errors.New("")
}

func (mt mockFalseTopupRepository) GetAllPaid(userId uint) ([]entities.TopUp, error) {
	return []entities.TopUp{{ID: 1, UserID: 1, InvoiceID: "XENDIT INVOICE ID", PaymentUrl: "XENDIT PAYMENT URL", Total: 10000, Status: "PAID"}}, errors.New("")
}

func (mt mockFalseTopupRepository) Update(extId string, topUp entities.TopUp) (entities.TopUp, error) {
	return entities.TopUp{ID: 1, UserID: 1, InvoiceID: "XENDIT INVOICE ID", PaymentUrl: "XENDIT PAYMENT URL", Total: 10000, Status: "PAID"}, errors.New("")
}

func (mt mockFalseTopupRepository) GetByInvoice(extId string) (entities.TopUp, error) {
	return entities.TopUp{ID: 1, UserID: 1, InvoiceID: "XENDIT INVOICE ID", PaymentUrl: "XENDIT PAYMENT URL", Total: 10000, Status: "PENDING"}, errors.New("")
}

func (mt mockFalseTopupRepository) GetUser(userId int) (entities.User, error) {
	return entities.User{ID: 1, Balance: 10000}, errors.New("")
}

func (mt mockFalseTopupRepository) UpdateUserBalance(userId int, user entities.User) (entities.User, error) {
	return entities.User{ID: 1, Balance: 10000}, errors.New("")
}
