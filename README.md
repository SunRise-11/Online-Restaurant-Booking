[![Go.Dev reference](https://img.shields.io/badge/echo-reference-blue?logo=go&logoColor=blue)](https://github.com/labstack/echo)
[![Go.Dev reference](https://img.shields.io/badge/gorm-reference-blue?logo=go&logoColor=blue)](https://pkg.go.dev/gorm.io/gorm?tab=doc)
[![Contributors][contributors-shield]][contributors-url]

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
- [How to use](#executing-program)
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

### Executing program

- How to run the program

    ```console
    go run main.go    
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
    ├── .env                    # 
    ├── .gitignore              #
    ├── go.mod                  
    ├── go.sum                  
    ├── main.go                 # Main Program
    └── README.md               

## Authors

[Andrew Prasetyo](https://github.com/andrewptjio)

[Herlianto](https://github.com/herlianto-github)

[Ilham Junius](https://github.com/ilhamjunius)

## Version History

- 0.0.1
  - Initial Release

## Acknowledgments

- [Layered Architecture](https://www.oreilly.com/library/view/software-architecture-patterns/9781491971437/ch01.html)

[controbutors-shield]: https://img.shields.io/github/contributors/herlianto-github/Restobook.svg?style=for-the-badge
[contributors-url]: https://github.com/herlianto-github/Restobook/graphs/contributors
