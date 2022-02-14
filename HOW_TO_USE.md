# How to use

- Create Database

```sql
mysql -u 'database name' -p 'database password'
create database 'database name'
```

- Create your .env file

```go
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

```go
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

- Create driver.go

```go
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

```go
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
