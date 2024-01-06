package AREA_management

import (
	"API/API_key"
	"API/DB_management"
	"API/usefull_functions"
	"API/Spotify"
	"API/Email"
	"API/call_API"
	"API/Lol"
	"API/Meteo"
	"API/Github"
	"API/Tft"
	"API/Time"
	"API/Steam"
	"API/Discord"

	"net/http"
	"log"
	// "fmt"
	"time"

	"github.com/gin-gonic/gin"
)

type Timer struct {
	Time string `json:"time"`
}

type checkAction struct {
	name string `json:"name"`
	action func(DB_management.Area) (call_API.CallReactionStruct, error) `json:"action"`
}

type checkReaction struct {
	name string `json:"name"`
	reaction func(call_API.CallReactionStruct) error `json:"reaction"`
}

type areaCreation struct {
	name string `json:"name"`
	creation func(DB_management.Area) error `json:"creation"`
}

type areaUpdate struct {
	name string `json:"name"`
	update func(DB_management.Area) error `json:"update"`
}

type areaGet struct {
	name string `json:"name"`
	get func(int) ([]interface{}, error) `json:"get"`
}

var listAreaCreation = []areaCreation {
	{name: "spotify", creation: Spotify.CheckActionCreation},
	{name: "lol", creation: Lol.CheckActionCreation},
	{name: "meteo", creation: Meteo.CheckActionCreation},
	{name: "tft", creation: Tft.CheckActionCreation},
	{name: "time", creation: Time.CheckActionCreation},
	{name: "github", creation: Github.CheckActionCreation},
	{name: "steam", creation: Steam.CheckActionCreation},
}

var listCheckAction = []checkAction {
	{name: "spotify", action: Spotify.CheckActionExecution},
	{name: "lol", action: Lol.CheckActionExecution},
	{name: "meteo", action: Meteo.CheckActionExecution},
	{name: "tft", action: Tft.CheckActionExecution},
	{name: "time", action: Time.CheckActionExecution},
	{name: "github", action: Github.CheckActionExecution},
	{name: "steam", action: Steam.CheckActionExecution},
}

var listCheckReaction = []checkReaction {
	{name: "email", reaction: Email.CheckReactionExecution},
	{name: "github", reaction: Github.CheckReactionExecution},
	{name: "discord", reaction: Discord.CheckReactionExecution},
}

var listAreaUpdate = []areaUpdate {
	{name: "spotify", update: Spotify.CheckActionUpdate},
	{name: "lol", update: Lol.CheckActionUpdate},
	{name: "meteo", update: Meteo.CheckActionUpdate},
	{name: "tft", update: Tft.CheckActionUpdate},
	{name: "time", update: Time.CheckActionUpdate},
	{name: "github", update: Github.CheckActionUpdate},
	{name: "steam", update: Steam.CheckActionUpdate},
}

var listAreaGet = []areaGet {
	{name: "spotify", get: Spotify.GetAllActionFromUser},
	{name: "lol", get: Lol.GetAllActionFromUser},
	{name: "meteo", get: Meteo.GetAllActionFromUser},
	{name: "tft", get: Tft.GetAllActionFromUser},
	{name: "time", get: Time.GetAllActionFromUser},
	{name: "github", get: Github.GetAllActionFromUser},
	{name: "steam", get: Steam.GetAllActionFromUser},
}

var listAREA = []string {
	"spotify",
	"discord",
	"lol",
	"email",
	"github",
	"time",
	"meteo",
	"tft",
	"steam",
}


var listTimers = []Timer{}

func CheckAreaStart() {
	for _, a := range listAREA {
		sqlStatement := `SELECT time_execute FROM areas_` + a + ` WHERE toggled = true AND time_execute IS NOT NULL`
		rows, err := DB_management.DB.Query(sqlStatement)
		if err != nil {
			log.Println("error = ", err)
			log.Println("error while getting AREA")
			return
		}
		defer rows.Close()
		for rows.Next() {
			var timer Timer
			err = rows.Scan(&timer.Time)
			if err != nil {
				log.Println("error = ", err)
				log.Println("error while getting AREA")
				return
			}
			listTimers = append(listTimers, timer)
		}
	}
}

func launchArea(timer string, c *gin.Context, inter call_API.CallReactionStruct) {
	for _, a := range listAREA {
		sqlStatement := `SELECT user_id, reaction, action_name, reaction_name, id FROM areas_` + a + ` WHERE toggled = true AND time_execute = $1`
		rows, err := DB_management.DB.Query(sqlStatement, timer)
		if !usefull_functions.CheckError(err, "error while getting AREA", "error while getting AREA from DB", c) {
			return
		}
		for rows.Next() {
			var newArea DB_management.Area
			err = rows.Scan(&newArea.UserID, &newArea.Reaction, &newArea.Action_name, &newArea.Reaction_name, &newArea.ID)
			if !usefull_functions.CheckError(err, "error while getting AREA", "error while getting AREA from DB", c) {
				return
			}
			for _, b := range listCheckAction {
				if b.name == a {
					inter, err = b.action(newArea)
					if err != nil {
						log.Println("error = ", err)
					}
					continue
				}
			}
			if err != nil {
				log.Println("pass reaction")
				continue
			}
			for _, b := range listCheckReaction {
				if b.name == newArea.Reaction {
					err = b.reaction(inter)
					if err != nil {
						log.Println("error = ", err)
					}
					continue
				}
			}
		}
		defer rows.Close()
	}
}


func CheckArea(c *gin.Context) {
	if API_key.CheckApiKeyAdmin(c) == false {
		return
	}

	var inter call_API.CallReactionStruct

	new := time.Now()
	launchArea(new.Format("15:04"), c, inter)
	launchArea("*", c, inter)
	return
}

// PostArea godoc
// @Summary Create a new AREA
// @Description Create a new AREA with specified parameters.
// @Tags AREA
// @Accept  json
// @Produce  json
// @Param   x-api-key   header    string             true  "API Key"
// @Param   bearer      header    string             true  "Bearer Token"
// @Param   area        body      AreaCreationRequest true  "AREA Creation Data"
// @Success 201 {object} AreaCreationSuccessResponse "AREA created successfully"
// @Failure 400 {object} AreaCreationErrorResponse   "Invalid input data"
// @Failure 500 {string} string                     "Internal Server Error"
// @Router /create-area [post]
func PostArea(c *gin.Context) {
	if API_key.CheckApiKey(c) == false {
		return
	}

	var newArea DB_management.Area
	actionCorrect := false
	reactionCorrect := false

	if !usefull_functions.CheckError(c.ShouldBindJSON(&newArea), "no argument valid in json send by user", "no argument valid in json send by user", c) {
		return
	}

	for _, a := range listAREA {
		if a == newArea.Action {
			actionCorrect = true
		}
		if a == newArea.Reaction {
			reactionCorrect = true
		}
	}

	if actionCorrect == false  || reactionCorrect == false {
		log.Println("action not valid")
		c.IndentedJSON(http.StatusNotAcceptable, gin.H{"message": "AREA not valid"})
		return
	}


	sqlStatement := `SELECT id FROM users WHERE token = $1`
	var userID int
	err := DB_management.DB.QueryRow(sqlStatement, c.Request.Header.Get("bearer")).Scan(&userID)
	if !usefull_functions.CheckError(err, "Error with user token", "error while getting user id", c) {
		return
	}
	newArea.UserID = userID
	for _, a := range listAreaCreation {
		if a.name == newArea.Action {
			err = a.creation(newArea)
			if err != nil {
				log.Println("error = ", err)
				log.Println("error while creating AREA")
				c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error :" + err.Error()})
				return
			}
		}
	}
	log.Println("AREA created")
	sqlStatement = `SELECT id, time_execute FROM areas_` + newArea.Action + ` WHERE action_name = $1 AND reaction_name = $2 AND user_id = $3`
	err = DB_management.DB.QueryRow(sqlStatement, newArea.Action_name, newArea.Reaction_name, newArea.UserID).Scan(&newArea.ID, &newArea.Time_execute)
	if err != nil {
		log.Println("error = ", err)
		log.Println("error while getting AREA id")
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error while creating AREA please try again"})
		return
	}
	if newArea.Time_execute != "" {
		listTimers = append(listTimers, Timer{Time: newArea.Time_execute})
	}
	c.IndentedJSON(http.StatusCreated, gin.H{"message": "AREA created succesfully", "id": newArea.ID})
}

type AreaCreationRequest struct {
    Name          string `json:"name" example:"agagou"`
    Color         string `json:"color" example:"#FF0000"`
    Action        string `json:"action" example:"spotify"`
    Reaction      string `json:"reaction" example:"email"`
    Toggled       bool   `json:"toggled" example:true`
    TimeExecute   string `json:"time_execute" example:"*"`
    ActionName    string `json:"action_name" example:"new_song"`
    ReactionName  string `json:"reaction_name" example:"send_mail"`
}

// AreaCreationSuccessResponse defines the structure for a successful AREA creation response
type AreaCreationSuccessResponse struct {
    Message string `json:"message" example:"AREA created successfully"`
    ID      int    `json:"id" example:1`
}

// AreaCreationErrorResponse defines the structure for an AREA creation error response
type AreaCreationErrorResponse struct {
    Message string `json:"message" example:"AREA not valid"`
}


// PostModifyArea godoc
// @Summary Modify an existing AREA
// @Description Modify an existing AREA with new details provided by the user.
// @Tags AREA
// @Accept json
// @Produce json
// @Param bearer header string true "Bearer token for authorization"
// @Param x-api-key header string true "API key for access control"
// @Param area body Area true "Area information to be modified"
// @Success 201 {object} Area "AREA modified successfully"
// @Failure 400 {object} ErrorResponse "Bad request when the parameters are not filled"
// @Failure 401 {object} ErrorResponse "Unauthorized if the token is not valid or missing"
// @Failure 500 {object} ErrorResponse "Internal Server Error when AREA modification fails"
// @Router /modify-area [post]
func PostModifyArea (c *gin.Context) {
	var admin bool

	if API_key.CheckApiKeyAdmin(c) == true {
		admin = true
	} else if API_key.CheckApiKeyUser(c) == false {
		return
	}

	var modifyArea DB_management.Area
	userToken := c.Request.Header.Get("bearer")

	if err := c.BindJSON(&modifyArea); err != nil {
		log.Println("no argument valid in json send by user")
		c.IndentedJSON(http.StatusNotAcceptable, gin.H{"message": "Parameters not filled"})
		return
	}
	userID, err := DB_management.GetIntFromDB(`SELECT id FROM users WHERE token = $1`, userToken)
	if err != nil {
		log.Println("error while getting user id")
		c.IndentedJSON(http.StatusNotAcceptable, gin.H{"message": "Error with user token"})
		return
	}
	oldArea, err := DB_management.GetAreaWithId(modifyArea.ID, modifyArea.Action)
	if err != nil {
		log.Println("error while getting old area")
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error while modifying AREA please try again"})
		return
	}
	if (oldArea.UserID != userID && admin == false) {
		log.Println("token not valid")
		c.IndentedJSON(http.StatusNotAcceptable, gin.H{"message": "Token not valid"})
		return
	}
	for _, a := range listAreaUpdate {
		if a.name == modifyArea.Action {
			err = a.update(modifyArea)
			if err != nil {
				c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
				return
			}
		}
	}
	if modifyArea.Time_execute != "" {
		listTimers = append(listTimers, Timer{Time: modifyArea.Time_execute})
	}
	c.IndentedJSON(http.StatusCreated, gin.H{"message": "AREA modified succesfully"})
}

type Area struct {
    Name          string `json:"name"`
    Color         string `json:"color"`
    Action        string `json:"action"`
    Reaction      string `json:"reaction"`
    Toggled       bool   `json:"toggled"`
    ID            int    `json:"id"`
    ActionName    string `json:"action_name"`
    ReactionName  string `json:"reaction_name"`
}

// ErrorResponse represents the structure of an error response
type ErrorResponse struct {
    Message string `json:"message"`
}

// PostGetAllArea godoc
// @Summary Get all areas for a user
// @Description Retrieves all the areas associated with the user identified by the provided token.
// @Tags AREA
// @Accept  json
// @Produce  json
// @Param   bearer   header    string  true  "Bearer Token"
// @Success 200 {object} map[string]interface{} "Successful retrieval of areas"
// @Failure 400 {object} map[string]string "Invalid request data"
// @Failure 401 {object} map[string]string "Invalid API key or bearer token"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /get-area [post]
func PostGetAllArea(c *gin.Context) {
	var userID int
	var err error
	admin := false

	if API_key.CheckApiKeyAdmin(c) == true {
		admin = true
	} else if API_key.CheckApiKeyUser(c) == false {
		c.IndentedJSON(http.StatusNotAcceptable, gin.H{"message": "invalid API key"})
		return
	}

	if admin == false {
		userToken := c.Request.Header.Get("bearer")
		userID, err = DB_management.GetIntFromDB(`SELECT id FROM users WHERE token = $1`, userToken)
		if err != nil {
			log.Println("error while getting user id")
			c.IndentedJSON(http.StatusNotAcceptable, gin.H{"message": "Error with user id"})
			return
		}
	} else {
		type RequestData struct {
			ID int `json:"user_id"`
		}
		var requestData RequestData
		if err := c.ShouldBindJSON(&requestData); err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		userID = requestData.ID
	}
	var listArea []interface{}
	for _, a := range listAreaGet {
		res, err := a.get(userID)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
		listArea = append(listArea, res...)
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "AREA get succesfully", "areas": listArea})
}

// PostDeleteArea godoc
// @Summary Delete an AREA
// @Description Delete an existing AREA by its ID.
// @Tags AREA
// @Accept json
// @Produce json
// @Param x-api-key header string true "API key for access"
// @Param bearer header string true "Bearer token for user authentication"
// @Param body body DeleteAreaRequest true "Request body for deleting an AREA"
// @Success 200 {object} map[string]interface{} "message: AREA deleted successfully"
// @Failure 400 {object} map[string]interface{} "error: error message"
// @Failure 401 {object} map[string]interface{} "message: invalid API key or Token not valid"
// @Failure 500 {object} map[string]interface{} "message: Error while deleting AREA please try again"
// @Router /delete-area [post]
func PostDeleteArea(c *gin.Context) {
	type RequestData struct {
		ID int `json:"area_id"`
		Action string `json:"action"`
	}
	var id int
	admin := false

	if API_key.CheckApiKeyAdmin(c) == true {
		admin = true
	} else if API_key.CheckApiKeyUser(c) == false {
		c.IndentedJSON(http.StatusNotAcceptable, gin.H{"message": "invalid API key"})
		return
	}

	var requestData RequestData
	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id = requestData.ID
	action := requestData.Action
	if admin == false {
		userToken := c.Request.Header.Get("bearer")
		userID, err := DB_management.GetIntFromDB(`SELECT id FROM users WHERE token = $1`, userToken)
		if err != nil {
			log.Println("error while getting user id from token")
			c.IndentedJSON(http.StatusNotAcceptable, gin.H{"message": "Error with token"})
			return
		}
		sqlStatement := `SELECT user_id FROM areas_` + action + ` WHERE id = $1`
		resultID, err := DB_management.GetIntFromDB(sqlStatement, id)
		if err != nil {
			log.Println("error while getting user id")
			c.IndentedJSON(http.StatusNotAcceptable, gin.H{"message": "Error with user id"})
			return
		}
		if userID != resultID {
			log.Println("token not valid")
			c.IndentedJSON(http.StatusNotAcceptable, gin.H{"message": "Token not valid"})
			return
		}
	}
	sqlStatement := `DELETE FROM areas_` + action + ` WHERE id = $1`
	_, err := DB_management.DB.Exec(sqlStatement, id)
	if err != nil {
		log.Println("error = ", err)
		log.Println("error while deleting AREA")
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error while deleting AREA please try again"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "AREA deleted succesfully"})
}

type DeleteAreaRequest struct {
    AreaID int `json:"area_id" binding:"required"`
}