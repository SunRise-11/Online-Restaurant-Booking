package common

var JWT_SECRET_KEY string
var XENDIT_SECRET_KEY string
var XENDIT_CALLBACK_TOKEN string
var UNIPDF_API_KEY string
var IMGUR_CLIENTID string
var FTP_ADDRESS string
var FTP_USERNAME string
var FTP_PASSWORD string

var Daytoint = []struct {
	Day string
	No  int
}{
	{"Monday", 0},
	{"Tuesday", 1},
	{"Wednesday", 2},
	{"Thursday", 3},
	{"Friday", 4},
	{"Saturday", 5},
	{"Sunday", 6},
}
