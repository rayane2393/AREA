package login

import (
	"net/http"
	// "database/sql"
	"os/exec"
	"log"
	"API/API_key"
	"API/DB_management"
	"API/OAuth"
	// "strconv"

    "github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

// func createUser(sqlStatement string, args ...interface{}) (error) {
// 	// github_id := 0
// 	_, err := DB_management.DB.Exec(sqlStatement, args...)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

func createToken() (string, error) {
	newTokenBytes, err := exec.Command("uuidgen").Output()
	if err != nil {
		log.Println("error while generating Token")
		return "", err
	}
	newToken := string(newTokenBytes)
	if last := len(newToken) - 1; last >= 0 && newToken[last] == '\n' {
        newToken = newToken[:last]
    }
	return newToken, nil
}




// Register godoc
// @Summary User registration
// @Description Register a new user with a username, email, and password.
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param   x-api-key   header    string             true  "API Key"
// @Param   user        body      RegisterCredentials  true  "User Registration Data"
// @Success 201 {object} RegisterSuccessResponse   "Account registered successfully"
// @Failure 400 {object} RegisterErrorResponse     "Invalid input data"
// @Failure 500 {string} string                    "Internal Server Error"
// @Router /register [post]
func Register(c *gin.Context) {
	var newUsers DB_management.User

	if API_key.CheckApiKey(c) == false {
		return
	}
	if err := c.BindJSON(&newUsers); err != nil {
		log.Println("error = ", err)
		log.Println("no argument valid in json send by user")
		c.IndentedJSON(http.StatusNotAcceptable, gin.H{"message": "parameters not fill"})
		return
	}
	for _, a := range DB_management.Users {
		if a.Username == newUsers.Username {
			log.Println("username already taken")
			c.IndentedJSON(http.StatusNotAcceptable, gin.H{"message": "Username already taken"})
			return
		}
		if a.Email == newUsers.Email {
			log.Println("email already taken")
			c.IndentedJSON(http.StatusNotAcceptable, gin.H{"message": "Email already taken"})
			return
		}
	}
	log.Printf("new user = %v", newUsers)
	if newUsers.Phone == nil {
		newUsers.Phone = new(string)
	}
	if DB_management.CreateUser(`INSERT INTO users (username, email, password, phone) VALUES($1, $2, $3, $4)`, newUsers.Username, *newUsers.Email, *newUsers.Password, *newUsers.Phone) != nil {
		log.Println("error while creating user")
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error while creating account please try again"})
		return
	}
	DB_management.Users = append(DB_management.Users, newUsers)
	c.IndentedJSON(http.StatusCreated, gin.H{"message": "Accout registered succesfully"})
	log.Println("user created and send correctly")
}


type RegisterCredentials struct {
    Username string `json:"username" example:"newuser"`
    Email    string `json:"email" example:"newuser@example.com"`
    Password string `json:"password" example:"securePassword123"`
}

// RegisterSuccessResponse defines the structure for a successful registration response
type RegisterSuccessResponse struct {
    Message string `json:"message" example:"Account registered successfully"`
}

// RegisterErrorResponse defines the structure for a registration error response
type RegisterErrorResponse struct {
    Message string `json:"message" example:"Username already taken"`
}

// Login godoc
// @Summary User login
// @Description Authenticate user and provide a token upon successful authentication.
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param   x-api-key   header    string  true  "API Key"
// @Param   credentials body      LoginCredentials  true  "User Credentials"
// @Success 200 {object} LoginSuccessResponse   "Login successful"
// @Failure 400 {object} LoginErrorResponse     "Invalid username or password"
// @Failure 500 {string} string                 "Internal Server Error"
// @Router /login [post]
func Login(c *gin.Context) {
	var connexionInformation DB_management.User

	if API_key.CheckApiKey(c) == false {
		return
	}
	if err := c.BindJSON(&connexionInformation); err != nil {
		log.Println("no argument valid in json send by user")
		c.IndentedJSON(http.StatusNotAcceptable, gin.H{"message": "Parameters not filled"})
		return
	}

	for i, a := range DB_management.Users {
		if (a.Username == connexionInformation.Username || (a.Email != nil && *a.Email == connexionInformation.Username)) && *a.Password == *connexionInformation.Password {
			log.Println("Login valid")
			newToken, err := createToken()
			if err != nil {
				c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error while creating account please try again"})
				return
			}
			a.Token = &newToken
			DB_management.Users[i].Token = a.Token
			log.Println("Token created")
			sqlStatement := `UPDATE users SET token = $1 WHERE username = $2;`
			_, err = DB_management.DB.Exec(sqlStatement, a.Token, a.Username)
			if err != nil {
				log.Println("error while updating token")
				panic(err)
			}
			c.IndentedJSON(http.StatusOK, gin.H{"message": "Login successful", "token": *a.Token})
			return
		}
	}
	log.Println("Invalid Login")
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Incorrect credential"})
}

type LoginCredentials struct {
	Password string `json:"password" example:"password123"`
    Username string `json:"username" example:"user@example.com"`
}

// LoginSuccessResponse defines the structure for a successful login response
type LoginSuccessResponse struct {
    Message string `json:"message" example:"Login successful"`
    Token   string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
}

// LoginErrorResponse defines the structure for a login error response
type LoginErrorResponse struct {
    Message string `json:"message" example:"Incorrect credentials"`
}

// GithubLogin godoc
// @Summary Login with GitHub OAuth
// @Description Login or register a new user using GitHub OAuth by providing a GitHub code.
// @Tags Auth
// @Accept json
// @Produce json
// @Param x-api-key header string true "API key for access control"
// @Param Bearer header string false "Bearer token if the user is already logged in"
// @Param code body Code true "GitHub code for OAuth"
// @Success 200 {object} LoginResponse "Login successful with token"
// @Failure 400 {object} ErrorResponse "Bad request if the code is not provided or invalid"
// @Failure 401 {object} ErrorResponse "Unauthorized if the API key is not valid"
// @Failure 500 {object} ErrorResponse "Internal Server Error when login fails"
// @Router /login/github [post]
func GithubLogin(c *gin.Context) {
	if API_key.CheckApiKey(c) == false {
		return
	}

	var codeData map[string]interface{}
	var newUser DB_management.User
	var err error

	if err := c.ShouldBindJSON(&codeData); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	code, exists := codeData["code"].(string)
    if !exists {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "code not found in the request body"})
        return
    }
	token, exists := c.Request.Header["Bearer"]
	if exists {
		newUser, err = OAuth.Github(code, token[0])
		if err != nil {
			log.Println("error while getting user from github")
			log.Println("error = ", err)
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
		c.IndentedJSON(http.StatusOK, gin.H{"message": "Login successful"})
		return
	} else {
		newUser, err = OAuth.Github(code, "")
	}
	if err != nil {
		log.Println("error while getting user from github")
		log.Println("error = ", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	for i, a := range DB_management.Users {
		if a.Github_id!= nil && a.Github_id == newUser.Github_id {
			newToken, err := createToken()
			if err != nil {
				c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error with github login please try again"})
				return
			}
			a.Token = &newToken
			DB_management.Users[i].Token = &newToken
			sqlStatement := `UPDATE users SET token = $1 WHERE username = $2;`
			_, err = DB_management.DB.Exec(sqlStatement, a.Token, a.Username)
			if err != nil {
				log.Println("error while updating token")
				panic(err)
			}
			c.IndentedJSON(http.StatusOK, gin.H{"message": "Login successful", "token": *a.Token})
			return
		}
	}
}

type Code struct {
    Code string `json:"code"`
}

// LoginResponse represents the response sent back to the user upon successful login
type LoginResponse struct {
    Message string `json:"message"`
    Token   string `json:"token,omitempty"`
}

// ErrorResponse represents the structure of an error response
type ErrorResponse struct {
    Error string `json:"error"`
}


// DiscordLogin godoc
// @Summary Login with discord OAuth
// @Description Login or register a new user using discord OAuth by providing a discord code.
// @Tags Auth
// @Accept json
// @Produce json
// @Param x-api-key header string true "API key for access control"
// @Param Bearer header string false "Bearer token if the user is already logged in"
// @Param code body Code true "discord code for OAuth"
// @Success 200 {object} LoginResponse "Login successful with token"
// @Failure 400 {object} ErrorResponse "Bad request if the code is not provided or invalid"
// @Failure 401 {object} ErrorResponse "Unauthorized if the API key is not valid"
// @Failure 500 {object} ErrorResponse "Internal Server Error when login fails"
// @Router /login/discord [post]
func DiscordLogin(c *gin.Context) {
	if API_key.CheckApiKey(c) == false {
		return
	}

	var codeData map[string]interface{}
	var newUser DB_management.User
	var err error
	if err := c.ShouldBindJSON(&codeData); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	code, exists := codeData["code"].(string)
	if !exists {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "code not found in the request body"})
		return
	}
	token, exists := c.Request.Header["Bearer"]
	if exists {
		newUser, err = OAuth.Discord(code, token[0])
		if err != nil {
			log.Println("error while getting user from discord")
			log.Println("error = ", err)
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
		c.IndentedJSON(http.StatusOK, gin.H{"message": "Login successful"})
		return
	} else {
		newUser, err = OAuth.Discord(code, "")
		if err != nil {
			log.Println("error while getting user from discord")
			log.Println("error = ", err)
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
		token, err := createToken()
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error with discord login please try again"})
			return
		}
		newUser.Token = &token
		sqlStatement := `UPDATE users SET token = $1 WHERE username = $2;`
		_, err = DB_management.DB.Exec(sqlStatement, newUser.Token, newUser.Username)
		if err != nil {
			log.Println("error while updating token")
			panic(err)
		}
		c.IndentedJSON(http.StatusOK, gin.H{"message": "Login successful", "token": *newUser.Token})
		return
	}
}

func FakeRouteGoogle(c *gin.Context) {
	log.Println("FakeRouteGoogle")
	return
}