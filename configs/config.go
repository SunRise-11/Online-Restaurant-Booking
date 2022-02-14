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
	// common.UNIPDF_API_KEY = os.Getenv("UniPDF_Api_Key")

	xendit.Opt.SecretKey = common.XENDIT_SECRET_KEY
	license.SetMeteredKey(os.Getenv("UniPDF_Api_Key"))

	return &defaultConfig

}
