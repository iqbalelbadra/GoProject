package controllers

import (
	"GoProject/configs"
	"GoProject/models"
	"net/http"

	"github.com/labstack/echo/v4"
)


func AddTransaksiController(c echo.Context) error {
	var transaksiRequest models.Transaksi

	c.Bind(&transaksiRequest)

	result := configs.DB.Create(&transaksiRequest)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Status: false,
			Message: "Failed insert data transaksi to database",
			Data: nil,
		})
	}

	return c.JSON(http.StatusCreated, models.BaseResponse{
		Status: true,
		Message: "Success insert data transaksi to database",
		Data: transaksiRequest,
	})
}

func GetDetailTransaksiController(c echo.Context) error {
	var id = c.Param("id_transaksi")
	var data = models.Transaksi{}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Status: true,
		Message: id,
		Data: data,
	})
}

func GetTransaksiController(c echo.Context) error {
	var data []models.transaksi

	result := configs.DB.Find(&data)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Status: false,
			Message: "Failed get data from transaksi",
			Data: nil,
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Status: true,
		Message: "Success get Data",
		Data: data,
	})
}