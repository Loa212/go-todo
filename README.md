# My first handbook

How I set up this project

## Init

```bash
go mod init
```

## Install some stuff

1. install [GORM](https://gorm.io/docs/) with postgres

```bash
go get -u gorm.io/gorm

go get -u gorm.io/driver/postgres
```

2. install [GIN](https://gin-gonic.com/docs/)

```bash
go get -u github.com/gin-gonic/gin
```

3. install BCrypt & JWT

```bash
go get -u golang.org/x/crypto/bcrypt

go get -u github.com/golang-jwt/jwt/v4

```

4. install [godotenv](https://pkg.go.dev/github.com/joho/godotenv#section-readme)

```bash
go get github.com/joho/godotenv

```

5. install [CompileDaemon](https://github.com/githubnemo/CompileDaemon) & install it so that u can use it as a command.

```bash
go get github.com/githubnemo/CompileDaemon

go install github.com/githubnemo/CompileDaemon


```

## Usage

in the terminal run the following command (will run on port:3000)

```bash
compiledaemon --command="./go-todo"
```

## License

[MIT](https://choosealicense.com/licenses/mit/)
