package user_management_admin

import (
	"API/API_key"
	"API/DB_management"

	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func GetUser(c *gin.Context) {
	type User struct {
		Username string `json:"username"`
		Email (*string) `json:"email"`
		Token string `json:"token"`
		Phone (*string) `json:"phone"`
		Github_id (*int) `json:"github_id"`
		Discord_id (*string) `json:"discord_id"`
	}
	if API_key.CheckApiKey(c) == false {
		return
	}
	var token string
	token = c.Request.Header.Get("Bearer")
	if token == "" {
		log.Println("No token found")
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "No token found"})
		return
	}
	sqlStatement := `SELECT username, email, token, phone, github_id, discord_id FROM users WHERE token = $1`
	var user User
	err := DB_management.DB.QueryRow(sqlStatement, token).Scan(&user.Username, &user.Email, &user.Token, &user.Phone, &user.Github_id, &user.Discord_id)
	if err != nil {
		log.Println("error while getting user")
		log.Println("error = ", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error while getting user"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "User found", "user": user})
}