package services

import "github.com/Leon4rdoMonteiro/gofiber.api/models"

func GetUserService() models.Users {
	result := models.Users{}

	user1 := models.User{Id: "1", Name: "Leonardo", Email: "leonardo.monteiro@email.com", Password: "123456789"}
	user2 := models.User{Id: "2", Name: "Maria", Email: "maria.monteiro@email.com", Password: "987654321"}

	result.Users = append(result.Users, user1, user2)

	return result
}
