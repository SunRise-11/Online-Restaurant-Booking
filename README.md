[![Go.Dev reference](https://img.shields.io/badge/echo-reference-blue?logo=go&logoColor=blue)](https://github.com/labstack/echo)
[![Go.Dev reference](https://img.shields.io/badge/gorm-reference-blue?logo=go&logoColor=blue)](https://pkg.go.dev/gorm.io/gorm?tab=doc)

[![Contributors](https://img.shields.io/github/contributors/herlianto-github/Restobook.svg?style=for-the-badge)](https://github.com/herlianto-github/Restobook/graphs/contributors)

# Restobook

<!-- Description -->
<br/>
<div align="center">
  <a href="https://github.com/herlianto-github/Restobook/IMAGES/Restobook.png">
    <img src="IMAGES/Restobook.png" alt="Logo" width="200" height="200">
  </a>
</div>
<div>
  <h3 align="center">Restobook</h3>
  <p align="center">
    Online Restaurant Booking
    <br/>
    <a href="https://github.com/herlianto-github/Restobook"><strong>Explore the docs »</strong></a>
    <br/>
    <!-- <br /> -->
    <a href="https://whimsical.com/online-order-QJZTHKQp4jGWeVMxMsmLiX">Wireframe</a>
    ·
    <a href="https://github.com/herlianto-github/Restobook/issues">Report Bug</a>
    ·
    <a href="https://github.com/herlianto-github/Restobook/issues">Request Feature</a>
  </p>
</div>

# Table of Content

- [Description](#restobook)
- [How To Use](HOW_TO_USE.md)
- [How To Contribute](CONRTIBUTING.md)
- [Roadmap](ROADMAP.md)
- [Entity Relationship Model](#entity–relationship-model)
- [Endpoints](#endpoints)
- [Folder Structure](#structuring)
- [Version History](#version-history)
- [Acknowledgments](#acknowledgments)
- [Authors](#authors)

## Entity–relationship model
  <!-- ERD -->
  <br/>
  <div align="center">
    <a href="https://github.com/herlianto-github/Restobook/ERD/erd_Resto.png">
      <img src="ERD/erd_Resto.png" alt="Logo">
    </a>
  </div>

## Endpoints

| Method | Endpoint | Description| Authentication | Authorization
|:-----|:--------|:----------| :----------:| :----------:|
| POST   | /admin/register   | Register a new admin                      | No | No
| POST   | /admin/login      | Login existing admin                      | No | No
| POST   | /admin/approve    | Approve Restaurant                        | Yes | Yes
| GET    | /admin/waiting    | Get all restaurant waiting for approval   | Yes | Yes
|---|---|---|---|---|
| POST   | /user/register    | Register a new user            | No | No
| POST   | /user/login       | Login existing user            | No | No
| GET    | /user             | Get current user profile       | Yes | Yes
| PUT    | /user             | Update current user profile    | Yes | Yes
|---|---|---|---|---|
| POST   | /topup            | Topup Balance for user         | Yes | Yes
| POST   | /topup/callback   | Callback for payment gateway   | No | No
| GET    | /topup/pending    | Get current user topup status  | Yes | Yes
| GET    | /topup/history    | Get current user topup history | Yes | Yes
|---|---|---|---|---|
| POST   | /restaurants/Register    | Register a new restaurant         | No | No
| POST   | /restaurants/Login       | Login existing restaurant         | No | No
| POST   | /myrestaurant/detail     | Create current restaurant detail  | Yes | Yes
| POST   | /restaurant/report?day=07&month=02&year=2022&export=EXCEL    | Create current restaurant Report PDF or EXCEL  | Yes | Yes
| GET    | /myrestaurant            | Get current restaurant profile    | Yes | Yes
| GET    | /restaurants             | Get all restaurant                | Yes | Yes
| GET    | /restaurants/open?date_time=2022-02-07 11:30:00              | Get all restaurant open by date | Yes | Yes
| PUT    | /myrestaurant            | Update current user profile       | Yes | Yes
| PUT    | /myrestaurant/detail     | Update current restaurant detail  | Yes | Yes
|---|---|---|---|---|
| POST   | /transaction                    | Create new transaction by user                   | Yes | Yes
| GET    | /transaction/waiting            | Get all list of current user waiting transaction | Yes | Yes
| GET    | /transaction/restaurant/waiting | Get list of all user waiting in restaurant       | Yes | Yes
| GET    | /transaction/restaurant/waiting | Get list of all user waiting in restaurant       | Yes | Yes
| GET    | /transaction/accepted           | Get accepted booking by user                     | Yes | Yes
| GET    | /transaction/history            | Get all transaction history booking by user      | Yes | Yes
| GET    | /transaction/rejected           | Get all transaction rejected booking by user     | Yes | Yes
| GET    | /transaction/success            | Get all transaction success booking by user      | Yes | Yes
| GET    | /transaction/fail               | Get all transaction fail booking by user         | Yes | Yes
| GET    | /transaction/cancel             | Get all transaction cancel booking by user       | Yes | Yes
|---|---|---|---|---|
| POST   | /ratings                        | Create new rating by user                        | Yes | Yes
| PUT    | /ratings/:restaurantId          | Update current rating                            | Yes | Yes
| DELETE | /ratings/:restaurantId          | Delete current rating                            | Yes | Yes
|---|---|---|---|---|

## Structuring

  ```sh
    Restobook
    ├── configs                
    │     └──config.go           # Configs files
    ├── delivery                 # Endpoints handlers or controllers
    │     └──common
    │     │   ├── global.go           # Constant variable
    │     │   └── http_responses.go   # Default http code, status, message
    │     └──controllers
    │     │   ├── users
    │     │   ├── formatter_req.go    # Default request format for spesific controllers
    │     │   ├── formatter_res.go    # Default response format for spesific controllers
    │     │   ├── users_test.go       # Unit tests for spesific controllers
    │     │   └── users.go            # Spesific controller
    │     └──helpers
    │     │   └── helper.go           # Helper Function
    │     └──routes  
    │         └── routes.go           # Endpoints list
    ├── entities                
    │     └── users.go          # database model
    ├── repository              
    │     ├── interface.go      # Repository Interface for controllers
    │     ├── users_test.go     # Unit test for spesific repository
    │     └── users.go          # Spesific Repository
    ├── utils                 
    │     └── driver.go         # Database driver
    ├── .env                    # Individual working environment variables
    ├── .gitignore              # Which files to ignore when committing
    ├── go.mod                  
    ├── go.sum                  
    ├── main.go                 # Main Program
    └── README.md    
  ```

## Version History

- 0.0.1
  - Endpoint Admin
    - **Register**<br/> /admin/register
    - **Login**<br/> /admin/login
  - Endpoint User
    - **Register**<br/> /users/register
    - **Login**<br/> /users/login
    - **Show User**<br/> /user
    - **Update User**<br/> /user
    - **Delete User**<br/> /user

- 0.0.2
  - Endpoint Restaurant
    - **Register**<br/> /restaurants/register
    - **Login**<br/> /restaurants/login
    - **Show Restaurant**<br/> /myrestaurant
    - **Update Restaurant**<br/> /myrestaurant
    - **Create Restaurant Detail**<br/> /myrestaurant/detail
    - **Update Restaurant Detail**<br/> /myrestaurant/detail
- 0.0.3
  - Third Release
    - coming soon

## Acknowledgments

- [Layered Architecture](https://www.oreilly.com/library/view/software-architecture-patterns/9781491971437/ch01.html)

## Authors

- [Andrew Prasetyo](https://github.com/andrewptjio) (Person In Charge and maintainer)
- [Herlianto](https://github.com/herlianto-github) (Author and maintainer)
- [Ilham Junius](https://github.com/ilhamjunius) (Author and maintainer)
