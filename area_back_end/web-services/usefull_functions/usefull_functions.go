package usefull_functions

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckError(err error, message_user string, message_log string, c *gin.Context) (bool) {
	if err != nil {
		log.Println("error = ", err)
		log.Println(message_log)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error": message_user,
		})
		return false
	}
	return true
}