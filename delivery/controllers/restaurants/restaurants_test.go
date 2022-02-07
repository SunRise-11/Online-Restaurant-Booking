package restaurants

// var jwtToken string
// var jwtTokenAdmin string

// func TestRestaurant(t *testing.T) {
// 	config := configs.GetConfig()
// 	fmt.Println(config)

// 	ec := echo.New()

// 	t.Run("Register Admin", func(t *testing.T) {
// 		reqBody, _ := json.Marshal(map[string]string{
// 			"email":    "admin@outlook.my",
// 			"password": "admin",
// 		})

// 		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
// 		res := httptest.NewRecorder()

// 		req.Header.Set("Content-Type", "application/json")

// 		context := ec.NewContext(req, res)
// 		context.SetPath("/admin/register")

// 		adminCtrl := auth.NewAdminControllers(mockUserRepository{})
// 		adminCtrl.RegisterAdminCtrl()(context)

// 		responses := RestaurantResponseFormat{}
// 		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

// 		assert.Equal(t, 200, responses.Code)
// 		assert.Equal(t, "Successful Operation", responses.Message)
// 	})

// 	t.Run("Login Admin", func(t *testing.T) {
// 		reqBody, _ := json.Marshal(map[string]string{
// 			"email":    "admin@outlook.my",
// 			"password": "admin",
// 		})

// 		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
// 		res := httptest.NewRecorder()

// 		req.Header.Set("Content-Type", "application/json")

// 		context := ec.NewContext(req, res)
// 		context.SetPath("/admin/login")

// 		adminCtrl := auth.NewAdminControllers(mockUserRepository{})
// 		adminCtrl.LoginAdminCtrl()(context)

// 		responses := LoginResponseFormat{}
// 		json.Unmarshal([]byte(res.Body.Bytes()), &responses)
// 		jwtTokenAdmin = responses.Token

// 		assert.Equal(t, 200, responses.Code)
// 		assert.Equal(t, "Successful Operation", responses.Message)
// 	})

// 	t.Run("Register Restaurant", func(t *testing.T) {
// 		reqBody, _ := json.Marshal(map[string]string{
// 			"email":    "restaurant1@outlook.my",
// 			"password": "resto123",
// 		})

// 		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
// 		res := httptest.NewRecorder()

// 		req.Header.Set("Content-Type", "application/json")

// 		context := ec.NewContext(req, res)
// 		context.SetPath("/restaurants/register")

// 		restoCtrl := NewRestaurantsControllers(mockRestaurantRepository{})
// 		restoCtrl.RegisterRestoCtrl()(context)

// 		responses := RestaurantResponseFormat{}
// 		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

// 		assert.Equal(t, 200, responses.Code)
// 		assert.Equal(t, "Successful Operation", responses.Message)
// 	})

// 	t.Run("Login Restaurant", func(t *testing.T) {
// 		reqBody, _ := json.Marshal(map[string]string{
// 			"email":    "restaurant1@outlook.my",
// 			"password": "resto123",
// 		})

// 		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
// 		res := httptest.NewRecorder()

// 		req.Header.Set("Content-Type", "application/json")

// 		context := ec.NewContext(req, res)
// 		context.SetPath("/restaurants/login")

// 		restoCtrl := NewRestaurantsControllers(mockRestaurantRepository{})
// 		restoCtrl.LoginRestoCtrl()(context)

// 		responses := LoginResponseFormat{}
// 		json.Unmarshal([]byte(res.Body.Bytes()), &responses)
// 		jwtToken = responses.Token

// 		assert.Equal(t, 200, responses.Code)
// 		assert.Equal(t, "Successful Operation", responses.Message)
// 	})

// 	t.Run("Get Restaurant", func(t *testing.T) {
// 		req := httptest.NewRequest(http.MethodGet, "/", nil)
// 		res := httptest.NewRecorder()

// 		req.Header.Set("Content-Type", "application/json")
// 		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

// 		context := ec.NewContext(req, res)
// 		context.SetPath("/restaurant")

// 		restaurantCtrl := NewRestaurantsControllers(mockRestaurantRepository{})
// 		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(restaurantCtrl.GetRestoByIdCtrl())(context); err != nil {
// 			log.Fatal(err)
// 			return
// 		}

// 		responses := RestaurantResponseFormat{}
// 		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

// 		assert.Equal(t, 200, responses.Code)
// 		assert.Equal(t, "Successful Operation", responses.Message)
// 	})

// 	t.Run("Update Restaurant", func(t *testing.T) {
// 		reqBody, _ := json.Marshal(map[string]string{
// 			"email":    "UPDATErestaurant@outlook.my",
// 			"password": "resto123",
// 		})

// 		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
// 		res := httptest.NewRecorder()

// 		req.Header.Set("Content-Type", "application/json")
// 		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

// 		context := ec.NewContext(req, res)
// 		context.SetPath("/restaurant")

// 		restaurantCtrl := NewRestaurantsControllers(mockRestaurantRepository{})
// 		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(restaurantCtrl.UpdateRestoByIdCtrl())(context); err != nil {
// 			log.Fatal(err)
// 			return
// 		}

// 		responses := RestaurantResponseFormat{}
// 		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

// 		assert.Equal(t, 200, responses.Code)
// 		assert.Equal(t, "Successful Operation", responses.Message)
// 	})

// 	t.Run("Create Detail Restaurant", func(t *testing.T) {
// 		reqBody, _ := json.Marshal(map[string]string{
// 			"name": "resto 1",
// 		})

// 		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
// 		res := httptest.NewRecorder()

// 		req.Header.Set("Content-Type", "application/json")
// 		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

// 		context := ec.NewContext(req, res)
// 		context.SetPath("/restaurant/detail")

// 		restaurantCtrl := NewRestaurantsControllers(mockRestaurantRepository{})
// 		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(restaurantCtrl.CreateDetailRestoByIdCtrl())(context); err != nil {
// 			log.Fatal(err)
// 			return
// 		}

// 		responses := RestaurantResponseFormat{}
// 		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

// 		assert.Equal(t, 200, responses.Code)
// 		assert.Equal(t, "Successful Operation", responses.Message)
// 	})

// 	t.Run("Update Detail Restaurant", func(t *testing.T) {
// 		now := time.Now()
// 		reqBody, _ := json.Marshal(map[string]interface{}{
// 			"open":  now,
// 			"close": now,
// 		})

// 		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(reqBody))
// 		res := httptest.NewRecorder()

// 		req.Header.Set("Content-Type", "application/json")
// 		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

// 		context := ec.NewContext(req, res)
// 		context.SetPath("/restaurant/detail")

// 		restaurantCtrl := NewRestaurantsControllers(mockRestaurantRepository{})
// 		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(restaurantCtrl.UpdateDetailRestoByIdCtrl())(context); err != nil {
// 			log.Fatal(err)
// 			return
// 		}

// 		responses := RestaurantResponseFormat{}
// 		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

// 		assert.Equal(t, 200, responses.Code)
// 		assert.Equal(t, "Successful Operation", responses.Message)
// 	})

// 	t.Run("Get Waiting Restaurant", func(t *testing.T) {

// 		req := httptest.NewRequest(http.MethodPut, "/", nil)
// 		res := httptest.NewRecorder()

// 		req.Header.Set("Content-Type", "application/json")
// 		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenAdmin))

// 		context := ec.NewContext(req, res)
// 		context.SetPath("/admin/waiting")

// 		restaurantCtrl := NewRestaurantsControllers(mockRestaurantRepository{})
// 		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(restaurantCtrl.GetsWaiting())(context); err != nil {
// 			log.Fatal(err)
// 			return
// 		}

// 		responses := RestaurantResponseFormat{}
// 		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

// 		assert.Equal(t, 200, responses.Code)
// 		assert.Equal(t, "Successful Operation", responses.Message)
// 	})

// 	t.Run("Approve Restaurant", func(t *testing.T) {

// 		req := httptest.NewRequest(http.MethodPut, "/", nil)
// 		res := httptest.NewRecorder()

// 		req.Header.Set("Content-Type", "application/json")
// 		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenAdmin))

// 		context := ec.NewContext(req, res)
// 		context.SetPath("/admin/approve")

// 		restaurantCtrl := NewRestaurantsControllers(mockRestaurantRepository{})
// 		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(restaurantCtrl.Approve())(context); err != nil {
// 			log.Fatal(err)
// 			return
// 		}

// 		responses := RestaurantResponseFormat{}
// 		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

// 		assert.Equal(t, 200, responses.Code)
// 		assert.Equal(t, "Successful Operation", responses.Message)
// 	})

// 	t.Run("Show All Restaurant", func(t *testing.T) {

// 		req := httptest.NewRequest(http.MethodGet, "/", nil)
// 		res := httptest.NewRecorder()

// 		req.Header.Set("Content-Type", "application/json")

// 		context := ec.NewContext(req, res)
// 		context.SetPath("/restaurants")

// 		restaurantCtrl := NewRestaurantsControllers(mockRestaurantRepository{})
// 		restaurantCtrl.Gets()(context)

// 		responses := RestaurantResponseFormat{}
// 		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

// 		assert.Equal(t, 200, responses.Code)
// 		assert.Equal(t, "Successful Operation", responses.Message)
// 	})

// 	t.Run("Show All open by day", func(t *testing.T) {

// 		req := httptest.NewRequest(http.MethodGet, "/", nil)
// 		res := httptest.NewRecorder()

// 		req.Header.Set("Content-Type", "application/json")

// 		context := ec.NewContext(req, res)
// 		context.SetPath("/restaurants/open?open=Monday&operational_hour=10:00")

// 		restaurantCtrl := NewRestaurantsControllers(mockRestaurantRepository{})
// 		restaurantCtrl.GetsByOpen()(context)

// 		responses := RestaurantResponseFormat{}
// 		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

// 		assert.Equal(t, 200, responses.Code)
// 		assert.Equal(t, "Successful Operation", responses.Message)
// 	})

// 	t.Run("Delete Restaurant", func(t *testing.T) {

// 		req := httptest.NewRequest(http.MethodDelete, "/", nil)
// 		res := httptest.NewRecorder()

// 		req.Header.Set("Content-Type", "application/json")
// 		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenAdmin))

// 		context := ec.NewContext(req, res)
// 		context.SetPath("/restaurant")

// 		restaurantCtrl := NewRestaurantsControllers(mockRestaurantRepository{})
// 		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(restaurantCtrl.DeleteRestaurantCtrl())(context); err != nil {
// 			log.Fatal(err)
// 			return
// 		}

// 		responses := RestaurantResponseFormat{}
// 		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

// 		assert.Equal(t, 200, responses.Code)
// 		assert.Equal(t, "Successful Operation", responses.Message)
// 	})

// }

// type mockUserRepository struct{}

// func (m mockUserRepository) RegisterAdmin(newUser entities.User) (entities.User, error) {
// 	return entities.User{ID: 1, Name: "admin"}, nil
// }

// func (m mockUserRepository) Register(newUser entities.User) (entities.User, error) {
// 	return entities.User{ID: 1, Name: "herlianto"}, nil
// }

// func (m mockUserRepository) LoginUser(email, password string) (entities.User, error) {
// 	hash := sha256.Sum256([]byte("herlianto123"))
// 	passwordS := fmt.Sprintf("%x", hash[:])
// 	return entities.User{ID: 1, Name: "herlianto", Password: passwordS, Email: "herlianto@outlook.my"}, nil
// }

// func (m mockUserRepository) Get(userID uint) (entities.User, error) {
// 	return entities.User{ID: 1, Name: "herlianto"}, nil
// }

// func (m mockUserRepository) Update(userID uint, updateUser entities.User) (entities.User, error) {
// 	return entities.User{ID: 1, Name: "andrew"}, nil
// }

// func (m mockUserRepository) Delete(userID uint) (entities.User, error) {
// 	return entities.User{ID: 0}, nil
// }

// type mockRestaurantRepository struct{}

// func (m mockRestaurantRepository) Register(newUser entities.Restaurant) (entities.Restaurant, error) {
// 	return entities.Restaurant{ID: 1, Email: "restaurant1@outlook.my"}, nil
// }

// func (m mockRestaurantRepository) LoginRestaurant(email, password string) (entities.Restaurant, error) {
// 	hash := sha256.Sum256([]byte("resto123"))
// 	passwordS := fmt.Sprintf("%x", hash[:])
// 	return entities.Restaurant{ID: 1, Email: "restaurant1@outlook.my", Password: passwordS}, nil
// }

// func (m mockRestaurantRepository) GetsWaiting() ([]entities.RestaurantDetail, error) {
// 	return []entities.RestaurantDetail{{ID: 1}}, nil
// }

// func (m mockRestaurantRepository) Approve(restaurantID uint, status string) (entities.RestaurantDetail, error) {
// 	return entities.RestaurantDetail{ID: 1}, nil
// }

// func (m mockRestaurantRepository) Get(restaurantID uint) (entities.Restaurant, entities.RestaurantDetail, error) {
// 	return entities.Restaurant{ID: 1}, entities.RestaurantDetail{ID: 1}, nil
// }

// func (m mockRestaurantRepository) GetsByOpen(open int) ([]entities.RestaurantDetail, error) {
// 	return []entities.RestaurantDetail{{ID: 1}}, nil
// }
// func (m mockRestaurantRepository) Gets() ([]entities.RestaurantDetail, error) {
// 	return []entities.RestaurantDetail{{ID: 1}}, nil
// }

// func (m mockRestaurantRepository) Update(restaurantID uint, updateUser entities.Restaurant) (entities.Restaurant, error) {
// 	return entities.Restaurant{ID: 1, Email: "restaurant1Update@outlook.my"}, nil
// }

// func (m mockRestaurantRepository) UpdateDetail(restaurantID uint, updateUser entities.RestaurantDetail) (entities.RestaurantDetail, error) {
// 	return entities.RestaurantDetail{ID: 1}, nil
// }

// func (m mockRestaurantRepository) Delete(restaurantID uint) (entities.Restaurant, error) {
// 	return entities.Restaurant{ID: 1}, nil
// }

// func TestFalseRestaurant(t *testing.T) {
// 	config := configs.GetConfig()
// 	fmt.Println(config)

// 	ec := echo.New()

// 	t.Run("Register Admin", func(t *testing.T) {
// 		reqBody, _ := json.Marshal(map[string]string{
// 			"email":    "admin@outlook.my",
// 			"password": "admin",
// 		})

// 		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
// 		res := httptest.NewRecorder()

// 		req.Header.Set("Content-Type", "application/json")

// 		context := ec.NewContext(req, res)
// 		context.SetPath("/admin/register")

// 		adminCtrl := auth.NewAdminControllers(mockUserRepository{})
// 		adminCtrl.RegisterAdminCtrl()(context)

// 		responses := RestaurantResponseFormat{}
// 		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

// 		assert.Equal(t, 200, responses.Code)
// 		assert.Equal(t, "Successful Operation", responses.Message)
// 	})

// 	t.Run("Login Admin", func(t *testing.T) {
// 		reqBody, _ := json.Marshal(map[string]string{
// 			"email":    "admin@outlook.my",
// 			"password": "admin",
// 		})

// 		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
// 		res := httptest.NewRecorder()

// 		req.Header.Set("Content-Type", "application/json")

// 		context := ec.NewContext(req, res)
// 		context.SetPath("/admin/login")

// 		adminCtrl := auth.NewAdminControllers(mockUserRepository{})
// 		adminCtrl.LoginAdminCtrl()(context)

// 		responses := LoginResponseFormat{}
// 		json.Unmarshal([]byte(res.Body.Bytes()), &responses)
// 		jwtTokenAdmin = responses.Token

// 		assert.Equal(t, 200, responses.Code)
// 		assert.Equal(t, "Successful Operation", responses.Message)
// 	})

// 	t.Run("Register Restaurant", func(t *testing.T) {
// 		reqBody, _ := json.Marshal(map[string]string{
// 			"email":    "restaurant1@outlook.my",
// 			"password": "resto123",
// 		})

// 		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
// 		res := httptest.NewRecorder()

// 		req.Header.Set("Content-Type", "application/json")

// 		context := ec.NewContext(req, res)
// 		context.SetPath("/restaurants/register")

// 		restoCtrl := NewRestaurantsControllers(mockRestaurantRepository{})
// 		restoCtrl.RegisterRestoCtrl()(context)

// 		responses := RestaurantResponseFormat{}
// 		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

// 		assert.Equal(t, 200, responses.Code)
// 		assert.Equal(t, "Successful Operation", responses.Message)
// 	})

// 	t.Run("FALSE Register Restaurant", func(t *testing.T) {
// 		reqBody, _ := json.Marshal(map[string]int{
// 			"email": 1,
// 		})

// 		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
// 		res := httptest.NewRecorder()

// 		req.Header.Set("Content-Type", "application/json")

// 		context := ec.NewContext(req, res)
// 		context.SetPath("/restaurants/register")

// 		restoCtrl := NewRestaurantsControllers(mockRestaurantRepository{})
// 		restoCtrl.RegisterRestoCtrl()(context)

// 		responses := RestaurantResponseFormat{}
// 		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

// 		assert.Equal(t, 400, responses.Code)
// 		assert.Equal(t, "Bad Request", responses.Message)
// 	})

// 	t.Run("FALSE Register Restaurant", func(t *testing.T) {
// 		reqBody, _ := json.Marshal(map[string]string{
// 			"email":    "restaurant1@outlook.my",
// 			"password": "resto123",
// 		})

// 		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
// 		res := httptest.NewRecorder()

// 		req.Header.Set("Content-Type", "application/json")

// 		context := ec.NewContext(req, res)
// 		context.SetPath("/restaurants/register")

// 		restoCtrl := NewRestaurantsControllers(mockFalseRestaurantRepository{})
// 		restoCtrl.RegisterRestoCtrl()(context)

// 		responses := RestaurantResponseFormat{}
// 		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

// 		assert.Equal(t, 500, responses.Code)
// 		assert.Equal(t, "Internal Server Error", responses.Message)
// 	})

// 	t.Run("FALSE Login Restaurant", func(t *testing.T) {
// 		reqBody, _ := json.Marshal(map[string]int{
// 			"email": 1,
// 		})

// 		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
// 		res := httptest.NewRecorder()

// 		req.Header.Set("Content-Type", "application/json")

// 		context := ec.NewContext(req, res)
// 		context.SetPath("/restaurants/login")

// 		restoCtrl := NewRestaurantsControllers(mockFalseRestaurantRepository{})
// 		restoCtrl.LoginRestoCtrl()(context)

// 		responses := LoginResponseFormat{}
// 		json.Unmarshal([]byte(res.Body.Bytes()), &responses)
// 		jwtToken = responses.Token

// 		assert.Equal(t, 400, responses.Code)
// 		assert.Equal(t, "Bad Request", responses.Message)
// 	})

// 	t.Run("FALSE Login Restaurant", func(t *testing.T) {
// 		reqBody, _ := json.Marshal(map[string]string{
// 			"email":    "restaurant1@outlook.com",
// 			"password": "resto",
// 		})

// 		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
// 		res := httptest.NewRecorder()

// 		req.Header.Set("Content-Type", "application/json")

// 		context := ec.NewContext(req, res)
// 		context.SetPath("/restaurants/login")

// 		restoCtrl := NewRestaurantsControllers(mockFalseRestaurantRepository{})
// 		restoCtrl.LoginRestoCtrl()(context)

// 		responses := LoginResponseFormat{}
// 		json.Unmarshal([]byte(res.Body.Bytes()), &responses)
// 		jwtToken = responses.Token

// 		assert.Equal(t, 404, responses.Code)
// 		assert.Equal(t, "Not Found", responses.Message)
// 	})

// 	t.Run("Login Restaurant", func(t *testing.T) {
// 		reqBody, _ := json.Marshal(map[string]string{
// 			"email":    "restaurant1@outlook.my",
// 			"password": "resto123",
// 		})

// 		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
// 		res := httptest.NewRecorder()

// 		req.Header.Set("Content-Type", "application/json")

// 		context := ec.NewContext(req, res)
// 		context.SetPath("/restaurants/login")

// 		restoCtrl := NewRestaurantsControllers(mockRestaurantRepository{})
// 		restoCtrl.LoginRestoCtrl()(context)

// 		responses := LoginResponseFormat{}
// 		json.Unmarshal([]byte(res.Body.Bytes()), &responses)
// 		jwtToken = responses.Token

// 		assert.Equal(t, 200, responses.Code)
// 		assert.Equal(t, "Successful Operation", responses.Message)
// 	})

// 	t.Run("FALSE Get Restaurant", func(t *testing.T) {
// 		req := httptest.NewRequest(http.MethodGet, "/", nil)
// 		res := httptest.NewRecorder()

// 		req.Header.Set("Content-Type", "application/json")
// 		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

// 		context := ec.NewContext(req, res)
// 		context.SetPath("/restaurant")

// 		restaurantCtrl := NewRestaurantsControllers(mockFalseRestaurantRepository{})
// 		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(restaurantCtrl.GetRestoByIdCtrl())(context); err != nil {
// 			log.Fatal(err)
// 			return
// 		}

// 		responses := RestaurantResponseFormat{}
// 		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

// 		assert.Equal(t, 404, responses.Code)
// 		assert.Equal(t, "Not Found", responses.Message)
// 	})

// 	t.Run("FALSE Get Waiting Restaurant", func(t *testing.T) {

// 		req := httptest.NewRequest(http.MethodGet, "/", nil)
// 		res := httptest.NewRecorder()

// 		req.Header.Set("Content-Type", "application/json")
// 		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

// 		context := ec.NewContext(req, res)
// 		context.SetPath("/admin/waiting")

// 		restaurantCtrl := NewRestaurantsControllers(mockFalseRestaurantRepository{})
// 		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(restaurantCtrl.GetsWaiting())(context); err != nil {
// 			log.Fatal(err)
// 			return
// 		}

// 		responses := RestaurantResponseFormat{}
// 		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

// 		assert.Equal(t, 406, responses.Code)
// 		assert.Equal(t, "Not Accepted", responses.Message)
// 	})

// 	t.Run("FALSE Get Waiting Restaurant", func(t *testing.T) {

// 		req := httptest.NewRequest(http.MethodGet, "/", nil)
// 		res := httptest.NewRecorder()

// 		req.Header.Set("Content-Type", "application/json")
// 		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenAdmin))

// 		context := ec.NewContext(req, res)
// 		context.SetPath("/admin/waiting")

// 		restaurantCtrl := NewRestaurantsControllers(mockFalseRestaurantRepository{})
// 		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(restaurantCtrl.GetsWaiting())(context); err != nil {
// 			log.Fatal(err)
// 			return
// 		}

// 		responses := RestaurantResponseFormat{}
// 		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

// 		assert.Equal(t, 404, responses.Code)
// 		assert.Equal(t, "Not Found", responses.Message)
// 	})

// 	t.Run("FALSE Update Restaurant", func(t *testing.T) {
// 		reqBody, _ := json.Marshal(map[string]int{
// 			"email": 1,
// 		})

// 		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
// 		res := httptest.NewRecorder()

// 		req.Header.Set("Content-Type", "application/json")
// 		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

// 		context := ec.NewContext(req, res)
// 		context.SetPath("/restaurant")

// 		restaurantCtrl := NewRestaurantsControllers(mockFalseRestaurantRepository{})
// 		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(restaurantCtrl.UpdateRestoByIdCtrl())(context); err != nil {
// 			log.Fatal(err)
// 			return
// 		}

// 		responses := RestaurantResponseFormat{}
// 		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

// 		assert.Equal(t, 400, responses.Code)
// 		assert.Equal(t, "Bad Request", responses.Message)
// 	})

// 	t.Run("FALSE Update Restaurant", func(t *testing.T) {
// 		reqBody, _ := json.Marshal(map[string]string{
// 			"email":    "restaurant1@outlook.my",
// 			"password": "resto12",
// 		})

// 		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
// 		res := httptest.NewRecorder()

// 		req.Header.Set("Content-Type", "application/json")
// 		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

// 		context := ec.NewContext(req, res)
// 		context.SetPath("/restaurant")

// 		restaurantCtrl := NewRestaurantsControllers(mockFalseRestaurantRepository{})
// 		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(restaurantCtrl.UpdateRestoByIdCtrl())(context); err != nil {
// 			log.Fatal(err)
// 			return
// 		}

// 		responses := RestaurantResponseFormat{}
// 		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

// 		assert.Equal(t, 404, responses.Code)
// 		assert.Equal(t, "Not Found", responses.Message)
// 	})

// 	t.Run("FALSE Create Detail Restaurant", func(t *testing.T) {
// 		reqBody, _ := json.Marshal(map[string]int{
// 			"name": 1,
// 		})

// 		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
// 		res := httptest.NewRecorder()

// 		req.Header.Set("Content-Type", "application/json")
// 		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

// 		context := ec.NewContext(req, res)
// 		context.SetPath("/restaurant/detail")

// 		restaurantCtrl := NewRestaurantsControllers(mockRestaurantRepository{})
// 		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(restaurantCtrl.CreateDetailRestoByIdCtrl())(context); err != nil {
// 			log.Fatal(err)
// 			return
// 		}

// 		responses := RestaurantResponseFormat{}
// 		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

// 		assert.Equal(t, 400, responses.Code)
// 		assert.Equal(t, "Bad Request", responses.Message)
// 	})

// 	t.Run("FALSE Create Detail Restaurant", func(t *testing.T) {
// 		reqBody, _ := json.Marshal(map[string]string{
// 			"name": "restoFalse",
// 		})

// 		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
// 		res := httptest.NewRecorder()

// 		req.Header.Set("Content-Type", "application/json")
// 		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

// 		context := ec.NewContext(req, res)
// 		context.SetPath("/restaurant/detail")

// 		restaurantCtrl := NewRestaurantsControllers(mockFalseRestaurantRepository{})
// 		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(restaurantCtrl.CreateDetailRestoByIdCtrl())(context); err != nil {
// 			log.Fatal(err)
// 			return
// 		}

// 		responses := RestaurantResponseFormat{}
// 		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

// 		assert.Equal(t, 404, responses.Code)
// 		assert.Equal(t, "Not Found", responses.Message)
// 	})

// 	t.Run("FALSE Update Detail Restaurant", func(t *testing.T) {
// 		reqBody, _ := json.Marshal(map[string]int{
// 			"open": 1,
// 		})

// 		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(reqBody))
// 		res := httptest.NewRecorder()

// 		req.Header.Set("Content-Type", "application/json")
// 		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

// 		context := ec.NewContext(req, res)
// 		context.SetPath("/restaurant/detail")

// 		restaurantCtrl := NewRestaurantsControllers(mockFalseRestaurantRepository{})
// 		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(restaurantCtrl.UpdateDetailRestoByIdCtrl())(context); err != nil {
// 			log.Fatal(err)
// 			return
// 		}

// 		responses := RestaurantResponseFormat{}
// 		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

// 		assert.Equal(t, 400, responses.Code)
// 		assert.Equal(t, "Bad Request", responses.Message)
// 	})

// 	t.Run("FALSE Update Detail Restaurant", func(t *testing.T) {
// 		now := time.Now()
// 		reqBody, _ := json.Marshal(map[string]interface{}{
// 			"open": now,
// 		})

// 		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(reqBody))
// 		res := httptest.NewRecorder()

// 		req.Header.Set("Content-Type", "application/json")
// 		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

// 		context := ec.NewContext(req, res)
// 		context.SetPath("/restaurant/detail")

// 		restaurantCtrl := NewRestaurantsControllers(mockFalseRestaurantRepository{})
// 		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(restaurantCtrl.UpdateDetailRestoByIdCtrl())(context); err != nil {
// 			log.Fatal(err)
// 			return
// 		}

// 		responses := RestaurantResponseFormat{}
// 		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

// 		assert.Equal(t, 404, responses.Code)
// 		assert.Equal(t, "Not Found", responses.Message)
// 	})

// 	t.Run("FALSE Approve Restaurant", func(t *testing.T) {

// 		reqBody, _ := json.Marshal(map[string]int{
// 			"resto_id": 1,
// 		})

// 		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(reqBody))
// 		res := httptest.NewRecorder()

// 		req.Header.Set("Content-Type", "application/json")
// 		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenAdmin))

// 		context := ec.NewContext(req, res)
// 		context.SetPath("/admin/approve")

// 		restaurantCtrl := NewRestaurantsControllers(mockFalseRestaurantRepository{})
// 		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(restaurantCtrl.Approve())(context); err != nil {
// 			log.Fatal(err)
// 			return
// 		}

// 		responses := RestaurantResponseFormat{}
// 		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

// 		assert.Equal(t, 404, responses.Code)
// 		assert.Equal(t, "Not Found", responses.Message)
// 	})

// 	t.Run("FALSE Approve Restaurant", func(t *testing.T) {

// 		reqBody, _ := json.Marshal(map[string]int{
// 			"resto_id": 1,
// 		})

// 		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(reqBody))
// 		res := httptest.NewRecorder()

// 		req.Header.Set("Content-Type", "application/json")
// 		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

// 		context := ec.NewContext(req, res)
// 		context.SetPath("/admin/approve")

// 		restaurantCtrl := NewRestaurantsControllers(mockFalseRestaurantRepository{})
// 		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(restaurantCtrl.Approve())(context); err != nil {
// 			log.Fatal(err)
// 			return
// 		}

// 		responses := RestaurantResponseFormat{}
// 		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

// 		assert.Equal(t, 406, responses.Code)
// 		assert.Equal(t, "Not Accepted", responses.Message)
// 	})

// 	t.Run("FALSE Approve Restaurant", func(t *testing.T) {

// 		reqBody, _ := json.Marshal(map[string]string{
// 			"resto_id": "asd",
// 		})

// 		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(reqBody))
// 		res := httptest.NewRecorder()

// 		req.Header.Set("Content-Type", "application/json")
// 		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenAdmin))

// 		context := ec.NewContext(req, res)
// 		context.SetPath("/admin/approve")

// 		restaurantCtrl := NewRestaurantsControllers(mockFalseRestaurantRepository{})
// 		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(restaurantCtrl.Approve())(context); err != nil {
// 			log.Fatal(err)
// 			return
// 		}

// 		responses := RestaurantResponseFormat{}
// 		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

// 		assert.Equal(t, 400, responses.Code)
// 		assert.Equal(t, "Bad Request", responses.Message)
// 	})

// 	t.Run("FALSE Show All Restaurant", func(t *testing.T) {

// 		req := httptest.NewRequest(http.MethodGet, "/", nil)
// 		res := httptest.NewRecorder()

// 		req.Header.Set("Content-Type", "application/json")

// 		context := ec.NewContext(req, res)
// 		context.SetPath("/restaurants")

// 		restaurantCtrl := NewRestaurantsControllers(mockFalseRestaurantRepository{})
// 		restaurantCtrl.Gets()(context)

// 		responses := RestaurantResponseFormat{}
// 		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

// 		assert.Equal(t, 404, responses.Code)
// 		assert.Equal(t, "Not Found", responses.Message)
// 	})

// 	t.Run("Show All open by day", func(t *testing.T) {

// 		req := httptest.NewRequest(http.MethodGet, "/", nil)
// 		res := httptest.NewRecorder()

// 		req.Header.Set("Content-Type", "application/json")

// 		context := ec.NewContext(req, res)
// 		context.SetPath("/restaurants/open?open=Monday&operational_hour=10:00")

// 		restaurantCtrl := NewRestaurantsControllers(mockFalseRestaurantRepository{})
// 		restaurantCtrl.GetsByOpen()(context)

// 		responses := RestaurantResponseFormat{}
// 		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

// 		assert.Equal(t, 404, responses.Code)
// 		assert.Equal(t, "Not Found", responses.Message)
// 	})

// 	t.Run("FALSE Delete Restaurant", func(t *testing.T) {

// 		reqBody, _ := json.Marshal(map[string]string{
// 			"resto_id": "a",
// 		})

// 		req := httptest.NewRequest(http.MethodDelete, "/", bytes.NewBuffer(reqBody))
// 		res := httptest.NewRecorder()

// 		req.Header.Set("Content-Type", "application/json")
// 		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenAdmin))

// 		context := ec.NewContext(req, res)
// 		context.SetPath("/restaurant")

// 		restaurantCtrl := NewRestaurantsControllers(mockFalseRestaurantRepository{})
// 		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(restaurantCtrl.DeleteRestaurantCtrl())(context); err != nil {
// 			log.Fatal(err)
// 			return
// 		}

// 		responses := RestaurantResponseFormat{}
// 		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

// 		assert.Equal(t, 400, responses.Code)
// 		assert.Equal(t, "Bad Request", responses.Message)
// 	})

// 	t.Run("FALSE Delete Restaurant", func(t *testing.T) {

// 		reqBody, _ := json.Marshal(map[string]int{
// 			"resto_id": 1,
// 		})

// 		req := httptest.NewRequest(http.MethodDelete, "/", bytes.NewBuffer(reqBody))
// 		res := httptest.NewRecorder()

// 		req.Header.Set("Content-Type", "application/json")
// 		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenAdmin))

// 		context := ec.NewContext(req, res)
// 		context.SetPath("/restaurant")

// 		restaurantCtrl := NewRestaurantsControllers(mockFalseRestaurantRepository{})
// 		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(restaurantCtrl.DeleteRestaurantCtrl())(context); err != nil {
// 			log.Fatal(err)
// 			return
// 		}

// 		responses := RestaurantResponseFormat{}
// 		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

// 		assert.Equal(t, 404, responses.Code)
// 		assert.Equal(t, "Not Found", responses.Message)
// 	})

// 	t.Run("FALSE Delete Restaurant", func(t *testing.T) {

// 		reqBody, _ := json.Marshal(map[string]int{
// 			"resto_id": 1,
// 		})

// 		req := httptest.NewRequest(http.MethodDelete, "/", bytes.NewBuffer(reqBody))
// 		res := httptest.NewRecorder()

// 		req.Header.Set("Content-Type", "application/json")
// 		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

// 		context := ec.NewContext(req, res)
// 		context.SetPath("/restaurant")

// 		restaurantCtrl := NewRestaurantsControllers(mockFalseRestaurantRepository{})
// 		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(restaurantCtrl.DeleteRestaurantCtrl())(context); err != nil {
// 			log.Fatal(err)
// 			return
// 		}

// 		responses := RestaurantResponseFormat{}
// 		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

// 		assert.Equal(t, 406, responses.Code)
// 		assert.Equal(t, "Not Accepted", responses.Message)
// 	})

// }

// type mockFalseRestaurantRepository struct{}

// func (m mockFalseRestaurantRepository) Register(newUser entities.Restaurant) (entities.Restaurant, error) {
// 	return entities.Restaurant{ID: 0, Email: "restaurant1@outlook.my"}, errors.New("")
// }

// func (m mockFalseRestaurantRepository) LoginRestaurant(email, password string) (entities.Restaurant, error) {
// 	hash := sha256.Sum256([]byte("resto123"))
// 	passwordS := fmt.Sprintf("%x", hash[:])
// 	return entities.Restaurant{ID: 0, Email: "restaurant1@outlook.my", Password: passwordS}, errors.New("")
// }

// func (m mockFalseRestaurantRepository) GetsWaiting() ([]entities.RestaurantDetail, error) {
// 	return []entities.RestaurantDetail{}, errors.New("")
// }

// func (m mockFalseRestaurantRepository) Approve(restaurantID uint, status string) (entities.RestaurantDetail, error) {
// 	return entities.RestaurantDetail{}, errors.New("")
// }

// func (m mockFalseRestaurantRepository) Get(restaurantID uint) (entities.Restaurant, entities.RestaurantDetail, error) {
// 	return entities.Restaurant{ID: 0}, entities.RestaurantDetail{ID: 0}, errors.New("")
// }

// func (m mockFalseRestaurantRepository) GetsByOpen(open int) ([]entities.RestaurantDetail, error) {
// 	return []entities.RestaurantDetail{}, errors.New("")
// }

// func (m mockFalseRestaurantRepository) Gets() ([]entities.RestaurantDetail, error) {
// 	return []entities.RestaurantDetail{}, errors.New("")
// }

// func (m mockFalseRestaurantRepository) Update(restaurantID uint, updateUser entities.Restaurant) (entities.Restaurant, error) {
// 	return entities.Restaurant{ID: 0, Email: "restaurant1Update@outlook.my"}, errors.New("")
// }

// func (m mockFalseRestaurantRepository) UpdateDetail(restaurantID uint, updateUser entities.RestaurantDetail) (entities.RestaurantDetail, error) {
// 	return entities.RestaurantDetail{ID: 0}, errors.New("")
// }

// func (m mockFalseRestaurantRepository) Delete(restaurantID uint) (entities.Restaurant, error) {
// 	return entities.Restaurant{ID: 0}, errors.New("")
// }
