# BE Go Lang Online Store

RESTful API for an online store built with Go, Gin framework, Gorm and MySQL database. It provides endpoints for managing users, products, categories, carts, and transactions.


### RESTful API Design

The API provides the following endpoints:
- User
  - GET /users
  - GET /users/:id
  - POST /users
  - PUT /users/:id
  - DELETE /users/:id

- Product
  - GET /products
  - GET /products/:id
  - POST /products
  - PUT /products/:id
  - DELETE /products/:id

- Category
  - GET /categories
  - GET /categories/:id
  - POST /categories
  - PUT /categories/:id
  - DELETE /categories/:id

- Cart
  - GET /carts
  - POST /carts
  - DELETE /carts/:id

- Transaction
  - GET /transactions
  - POST /transactions


## Getting Started

### Prerequisites

- Go 1.16 or higher
- MySQL

### Installation

#### 1. Clone the repo
   `git clone https://github.com/fahmiidris007/go-on-store.git`

#### 2. Install Go packages
   `go mod tidy`

#### 3. Enter your MySQL database information in .env file
   `DB_USER = your_user`
   
   `DB_PASSWORD = your_password`
   
   `DB_NAME = your_database_name`
   
   `SECRET_KEY = inisecretkey // let this default`

#### 4. Run the server
   _i using air package, so just type this in the terminal to run code_
   
  `air` 




## Usage

Use Postman or any API client to interact with the API endpoints. The postman config link will be provided later.


## Contact

*Fahmi Idris* - fahmiidrismz@gmail.com

## Acknowledgements

- [Gin Web Framework](https://github.com/gin-gonic/gin)
- [Gorm](https://gorm.io/index.html)
- [Air - Live reload for Go apps](https://github.com/cosmtrek/air)
