package controllers

import (
	"crypto/sha256"
	"encoding/hex"
	"GoProject/configs"
	"GoProject/middleware"
	"GoProject/models"
	"net/http"

	"github.com/labstack/echo/v4"
)


func RegisterController(c echo.Context) error {
	var userRequest models.User
	c.Bind(&userRequest)

	userRequest.Password, _ = generateSHA256Hash(userRequest.Password)

	result := configs.DB.Create(&userRequest)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Status: false,
			Message: "Failed to insert user",
			Data: nil,
		})
	}

	var userResponse = models.UserResponse {
		
		ID_User: userRequest.id_user,
		Email: userRequest.email,
		Nama: userRequest.nama,
		NoHP: userRequest.no_hp,
		Alamat: userRequest.alamat

	}

	userResponse.Token, _ = middleware.GenerateJwt(
		userResponse.nama,
		userResponse.id_user,
	)

	return c.JSON(http.StatusCreated, models.BaseResponse{
		Status: true,
		Message: "Success add user",
		Data: userResponse,
	})
}

func LoginController(c echo.Context) error {
	return c.JSON(http.StatusOK, nil)
}

func generateSHA256Hash(input string) (string, error) {
       hasher := sha256.New()

       _, err := hasher.Write([]byte(input))
    if err != nil {
        return "", err
    }

    hashBytes := hasher.Sum(nil)

    hashString := hex.EncodeToString(hashBytes)

    return hashString, nil
}

func GetUsersController(c echo.Context) error {
	var users []models.User

	return c.JSON(http.StatusOK, models.BaseResponse{
		Status: true,
		Message: "Berhasil get data",
		Data: users,
	})
}