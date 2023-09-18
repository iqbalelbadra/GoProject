package main

import (
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "GoProject/models"
	"GoProject/configs"
	"GoProject/routes"
	"log"
	"os"

	"github.com/labstack/echo/v4"
)



func main(){
	configs.InitDatabase()
	e := echo.New()
	routes.InitRoute(e)
	e.Start(getPort())
}

func getPort() string {
	port := os.Getenv("PORT")
    if port == "" {
        port = "8000"
    }
	return ":"+port
}