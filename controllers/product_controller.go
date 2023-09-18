package controllers

import (
	"GoProject/configs"
	"GoProject/models"
	"net/http"

	"github.com/labstack/echo/v4"
)


func AddProductController(c echo.Context) error {
	var productRequest models.Product

	c.Bind(&productRequest)

	result := configs.DB.Create(&productRequest)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Status: false,
			Message: "Failed insert data to database",
			Data: nil,
		})
	}

	return c.JSON(http.StatusCreated, models.BaseResponse{
		Status: true,
		Message: "Success insert data to database",
		Data: productRequest,
	})
}

func GetDetailProductController(c echo.Context) error {
	var id = c.Param("id_product")
	var data = models.Product{}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Status: true,
		Message: id,
		Data: data,
	})
}

func GetProductController(c echo.Context) error {
	var data []models.Product

	result := configs.DB.Find(&data)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Status: false,
			Message: "Failed get data from product",
			Data: nil,
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Status: true,
		Message: "Success get Data",
		Data: data,
	})
}