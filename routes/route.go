package routes

import (
	"GoProject/controllers"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRoute(e *echo.Echo) {
	e.POST("/register", controllers.RegisterController)
	e.POST("/login", controllers.LoginController)

	eAuth := e.Group("")
	eAuth.Use(echojwt.JWT([]byte(os.Getenv("PRIVATE_KEY_JWT"))))

	eLog := eAuth.Group("")
	eLog.Use(middleware.Logger())
	eLog.GET("/transaksi", controllers.GetTransaksiController)
	eLog.GET("/product", controllers.GetProductController)
	
	eAuth.GET("/users", controllers.GetUsersController)
	eAuth.POST("/transaksi", controllers.AddTransaksiController)
	eAuth.GET("/transaksi/:id", controllers.GetDetailTransaksiController)
	eAuth.POST("/product", controllers.AddProductController)
	eAuth.GET("/product/:id", controllers.GetDetailProductController)
}