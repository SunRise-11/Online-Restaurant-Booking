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
JWT_Secret_Key:"JWT Secret Key" 
Xendit_Secret_Key:""
Xendit_Callback_Token:""
UniPDF_Api_Key=""
Imgur_Client_ID=""
```

source : [Xendit Secret Key](https://dashboard.xendit.co/settings/developers#api-keys) , [Xendit Callback Token](https://dashboard.xendit.co/settings/developers#callbacks), [UniPDF Api Key](https://cloud.unidoc.io), [Imgur Client ID](https://apidocs.imgur.com)

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
