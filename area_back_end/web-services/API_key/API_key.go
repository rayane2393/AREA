package API_key

import (
	"log"
	"os/exec"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ApiKeys struct {
	Key string `json:"api key"`
}

var ApiKeysAdmin = []ApiKeys{
	{Key: "167bb87e-60ff-11ee-8c99-0242ac120002"},
}

var ApiKeysUsers = []ApiKeys{
	{Key: "ef4f85d0-6128-11ee-8c99-0242ac120002"},
}

func CreateKey(c *gin.Context) string {
	newTokenBytes, err := exec.Command("uuidgen").Output()
	if err != nil {
		log.Println("error while creating key")
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "error while creating Key please try again"})
		return ""
	}
	newToken := string(newTokenBytes)
	if last := len(newToken) - 1; last >= 0 && newToken[last] == '\n' {
		newToken = newToken[:last]
	}
	return newToken
}

func CheckApiKeyUser(c *gin.Context) bool {
	key := c.Request.Header.Get("X-API-Key")
	for _, userKey := range ApiKeysUsers {
		if key == userKey.Key {
			log.Println("API KEY user valid")
			return true
		}
	}
	log.Println("API KEY user not valid")
	return false
}

func CheckApiKeyAdmin(c *gin.Context) bool {
	key := c.Request.Header.Get("X-API-Key")
	for _, adminKey := range ApiKeysAdmin {
		if key == adminKey.Key {
			log.Println("API KEY admin valid")
			return true
		}
	}
	log.Println("API KEY admin not valid")
	return false
}

func CheckApiKey(c *gin.Context) bool {
	if CheckApiKeyAdmin(c) == false && CheckApiKeyUser(c) == false {
		c.IndentedJSON(http.StatusNotAcceptable, gin.H{"message": "invalid API key"})
		return false
	}
	return true
}