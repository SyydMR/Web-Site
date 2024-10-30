package main

// go get -u github.com/gin-gonic/gin
// go get -u gorm.io/driver/postgres
// go get -u github.com/spf13/viper
// go get -u github.com/golang-jwt/jwt/v4

import (
	"github.com/SyydMR/Web-Site/src/configs"
	"github.com/SyydMR/Web-Site/src/routes"
)

func main() {
    configs.Init()
    router := routes.GetRoute()
    router.Run(":8080")
}