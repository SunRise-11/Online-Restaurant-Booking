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
- [How to use](#how-to-use)
- [How to contribute](#how-to-contribute)
- [Endpoints](#endpoints)
- [Help](#help)
- [Authors](#authors)

## Getting Started

### Dependencies

- [Git](https://git-scm.com)
- [Golang](https://go.dev)
- [Visual Studio Code](https://code.visualstudio.com)

### How To Contribute

- Fork this repository

    ```console
    $ git clone https://github.com/YOUR_USERNAME/Restobook.git
    > Cloning into `Restobook`...
    > remote: Counting objects: 10, done.
    > remote: Compressing objects: 100% (8/8), done.
    > remove: Total 10 (delta 1), reused 10 (delta 1)
    > Unpacking objects: 100% (10/10), done.
    ```

    ```console
    cd Restobook
    ```

- Simple run  

    ```console
    go mod init Restobook
    ```

    ```console
    touch main.go    
    ```

    ```console
    echo 'package main 
    
    import "fmt"
    
    func main(){
    
        fmt.Println("Hello World")
    
    }' >> main.go
    ```

    ```console
    go run main.go
    ```

- Important

    ```console
    git checkout -b feature-name 
    ```

    Always create new branch when develop something

    ```console
    git add .    
    ```

    ```console
    git commit -m "feature description"
    ```

    ```console
    $ git remote -v
    > origin  https://github.com/YOUR_USERNAME/Restobook.git (fetch)
    > origin  https://github.com/YOUR_USERNAME/Restobook.git (push)
    ```

    ```console
    git remote add upstream https://github.com/herlianto-github/Restobook.git
    ```

    ```console
    $ git remote -v
    > origin    https://github.com/YOUR_USERNAME/Restobook.git (fetch)
    > origin    https://github.com/YOUR_USERNAME/Restobook.git (push)
    > upstream  https://github.com/herlianto-github/Restobook.git (fetch)
    > upstream  https://github.com/herlianto-github/Restobook.git (push)
    ```

    ```console
    git push -u origin feature-name    
    ```

### How To Use

- How to run the program

  - Create Database

    ```sh
    mysql -u 'database name' -p 'database password'
    create database 'database name'
    ```

  - Create your .env file

    ```sh
    APP_PORT:8000
    DB_DRIVER:"mysql"
    DB_ADDRESS:"localhost"
    DB_PORT:3306
    DB_USERNAME:"database username"
    DB_PASSWORD:"database password"
    DB_NAME:"database name"
    JWT_Secret_Key:"JWT Secret" 
    Xendit_Secret_Key:"https://dashboard.xendit.co/settings/developers#api-keys"
    Xendit_Callback_Token:"https://dashboard.xendit.co/settings/developers#callbacks"
    UniPDF_Api_Key="https://cloud.unidoc.io"
    Imgur_Client_ID="https://apidocs.imgur.com"
    ```

  - Create Config.go file

    ```sh
    package configs

    import (
      "Restobook/delivery/common"
      "log"
      "os"
      "sync"

      "github.com/joho/godotenv"
      "github.com/unidoc/unipdf/v3/common/license"
      "github.com/xendit/xendit-go"
    )

    type AppConfig struct {
      Port     string
      Database struct {
        Driver   string
        Name     string
        Address  string
        Port     string
        Username string
        Password string
      }
    }

    var lock = &sync.Mutex{}
    var appConfig *AppConfig

    func GetConfig() *AppConfig {
      lock.Lock()
      defer lock.Unlock()

      if appConfig == nil {
        appConfig = initConfig()
      }

      return appConfig
    }

    func initConfig() *AppConfig {

      err := godotenv.Load()
      if err != nil {
        log.Fatal("Error loading .env file")
      }

      var defaultConfig AppConfig
      defaultConfig.Port = os.Getenv("APP_PORT")
      defaultConfig.Database.Driver = os.Getenv("DB_DRIVER")
      defaultConfig.Database.Name = os.Getenv("DB_NAME")
      defaultConfig.Database.Address = os.Getenv("DB_ADDRESS")
      defaultConfig.Database.Port = os.Getenv("DB_PORT")
      defaultConfig.Database.Username = os.Getenv("DB_USERNAME")
      defaultConfig.Database.Password = os.Getenv("DB_PASSWORD")
      common.JWT_SECRET_KEY = os.Getenv("JWT_Secret_Key")
      common.XENDIT_SECRET_KEY = os.Getenv("Xendit_Secret_Key")
      common.XENDIT_CALLBACK_TOKEN = os.Getenv("Xendit_Callback_Token")
      common.IMGUR_CLIENTID = os.Getenv("Imgur_Client_ID")
      common.UNIPDF_API_KEY = os.Getenv("UniPDF_Api_Key")

      xendit.Opt.SecretKey = common.XENDIT_SECRET_KEY
      license.SetMeteredKey(common.UNIPDF_API_KEY)

      return &defaultConfig

    }

    ```

  - Create utils file

    ```sh
    package utils

    import (
      "Restobook/configs"
      "Restobook/entities"
      "fmt"

      "gorm.io/driver/mysql"
      "gorm.io/gorm"
    )

    func InitDB(config *configs.AppConfig) *gorm.DB {

      connectionString :=
        fmt.Sprintf(
          "%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local",
          config.Database.Username,
          config.Database.Password,
          config.Database.Address,
          config.Database.Port,
          config.Database.Name,
        )

      db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})

      if err != nil {
        panic(err)
      }

      InitialMigration(db)
      return db
    }
    func InitialMigration(db *gorm.DB) {

      db.Migrator().DropTable(&entities.User{})

      db.AutoMigrate(entities.User{})

    }

    ```

  - Create main.go file

    ```sh
    package main

    import (
      "Restobook/configs"
      "Restobook/delivery/controllers/users"
      "Restobook/delivery/routes"
      usersRepo "Restobook/repository/users"

      "Restobook/utils"
      "fmt"

      "github.com/labstack/echo/v4"
    )

    func main() {

      config := configs.GetConfig()
      db := utils.InitDB(config)      

      e := echo.New()
      usersRepo := usersRepo.NewUsersRepo(db)
      usersCtrl := users.NewUsersControllers(usersRepo)

      routes.RegisterPath(e, usersCtrl)

      e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", config.Port)))
    }

    ```

  - go run main.go

    ```sh
      ____    __
      / __/___/ /  ___
    / _// __/ _ \/ _ \
    /___/\__/_//_/\___/ v4.6.3
    High performance, minimalist Go web framework
    https://echo.labstack.com
    ____________________________________O/_______
                                        O\
    ⇨ http server started on [::]:8000
    ```

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

## Help

- **Configs**<br/>Contain database and http configuration
- **Delivery (API)**<br/>API http handlers or controllers
- **Entities**<br/>Contain database model
- **Repository** <br/> Contain implementation entities database anq query with ORM.
- **Utils**<br/>Contain database driver (mySQL)

## Structuring

    .
    ├── configs                
      ├──config.go              # Configs files
    ├── delivery                # Endpoints handlers or controllers
      ├──common
        ├── global.go           # Constant variable
        ├── http_responses.go   # Default http code, status, message
      ├──controllers
        ├── users
          ├── formatter_req.go  # Default request format for spesific controllers
          ├── formatter_res.go  # Default response format for spesific controllers
          ├── users_test.go     # Unit tests for spesific controllers
          ├── users.go          # Spesific controller
      ├──helpers
        ├── helper.go           # Helper Function
      ├──routes  
        ├── routes.go           # Endpoints
    ├── entities                
      ├── users.go              # database model
    ├── repository              
      ├── interface.go          # Repository Interface for controllers
      ├── users_test.go         # Unit test for spesific repository
      ├── users.go              # Spesific Repository
    ├── utils                 
      ├── databasedriver.go     # Database driver
    ├── .env                    # Individual working environment variables
    ├── .gitignore              # Which files to ignore when committing
    ├── go.mod                  
    ├── go.sum                  
    ├── main.go                 # Main Program
    └── README.md               

## Authors

[Andrew Prasetyo](https://github.com/andrewptjio) (Person In Charge and maintainer)

[Herlianto](https://github.com/herlianto-github) (Author and maintainer)

[Ilham Junius](https://github.com/ilhamjunius) (Author and maintainer)

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
