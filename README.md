# pretest privy

## How to setup
- `touch .env` in terminal (copy paste from .env.example)
```
    # Local
    APP_PORT = 45001
    SECRET_KEY = hahahahihihihi!

    # DATABASE
    DB_HOST=localhost
    DB_NAME=mybalance
    DB_PORT=5432
    DB_USER=postgres
    DB_PWD=postgres
    DB_SSL_MODE=disable
```
- create new database with name *mybalance*
- run with `go run main.go` in your terminal
- gorm will automatically migrate the database
- you can test API with postman

you can test this API with import this url in postman
https://api.postman.com/collections/18647539-77ba0c39-7809-4fa7-8b39-febd3cfdbf31?access_key=PMAT-01GJHETY7KJ9YCA242W5PYBD2B


## Techstack:
- Golang
- Framework: Gin-gonic (https://gin-gonic.com)
- ORM: GORM (https://gorm.io)
- Postgres (https://www.postgresql.org)

### Flow Top Up
- Create Account
- Login with the registered email
- after login you get token to create account bank (don't forget to use the *code* to make a transaction)
- after create account bank. next, you can top up with the code that has been registered at the bank.

### Flow Transfer balance
- Login with the registered email
- after login get token to create transaction(transfer balance). Make sure the recipient's email is registered

