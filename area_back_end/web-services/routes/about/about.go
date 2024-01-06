package about

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"log"
	"time"
	"strconv"
	"github.com/lib/pq"

	"API/DB_management"
)

// func About(c *gin.Context) {
// 	type client struct {
// 		Host string `json:"host"`
// 	}

// 	type reaction struct {
// 		Name string `json:"name"`
// 		Description string `json:"description"`
// 	}

// 	type action struct {
// 		Name string `json:"name"`
// 		Description string `json:"description"`
// 	}

// 	type service struct {
// 		Name string `json:"name"`
// 		Actions []action `json:"actions"`
// 		Reactions []reaction `json:"reactions"`
// 	}

// 	type server struct {
// 		Current_time string `json:"current_time"`
// 		Services []service `json:"services"`
// 	}

// 	type about struct {
// 		Client client `json:"client"`
// 		Server server `json:"server"`
// 	}

// 	var aboutInfo about
// 	var listServices []string = []string{"email", "lol", "pokemon", "spotify", "github", "discord", "calendar", "meteo"}
// 	var sqlStatement string

// 	aboutInfo.Client = client{}
// 	aboutInfo.Server = server{}
// 	aboutInfo.Server.Services = []service{}
// 	aboutInfo.Client.Host = c.ClientIP()
// 	aboutInfo.Server.Current_time = strconv.FormatInt(time.Now().Unix(), 10)
// 	for _, a := range listServices {
// 		var service service
// 		service.Name = a
// 		sqlStatement = "SELECT name, description FROM action WHERE services = $1"
// 		rows, err := DB_management.DB.Query(sqlStatement, a)
// 		if err != nil {
// 			log.Println("err = ", err)
// 			log.Println("error while getting actions")
// 			c.IndentedJSON(http.StatusNotAcceptable, gin.H{"message": "error while getting actions"})
// 			return
// 		}
// 		for rows.Next() {
// 			var action action
// 			err = rows.Scan(&action.Name, &action.Description)
// 			if err != nil {
// 				log.Println("error while getting actions")
// 				c.IndentedJSON(http.StatusNotAcceptable, gin.H{"message": "error while getting actions"})
// 				return
// 			}
// 			service.Actions = append(service.Actions, action)
// 		}
// 		defer rows.Close()
// 		sqlStatement = "SELECT name, description FROM reactions WHERE services = $1"
// 		rows, err = DB_management.DB.Query(sqlStatement, a)
// 		if err != nil {
// 			log.Println("error while getting reactions")
// 			c.IndentedJSON(http.StatusNotAcceptable, gin.H{"message": "error while getting reactions"})
// 			return
// 		}
// 		for rows.Next() {
// 			var reaction reaction
// 			err = rows.Scan(&reaction.Name, &reaction.Description)
// 			if err != nil {
// 				log.Println("error while getting reactions")
// 				c.IndentedJSON(http.StatusNotAcceptable, gin.H{"message": "error while getting reactions"})
// 				return
// 			}
// 			service.Reactions = append(service.Reactions, reaction)
// 		}
// 		defer rows.Close()
// 		aboutInfo.Server.Services = append(aboutInfo.Server.Services, service)
// 	}
// 	c.IndentedJSON(http.StatusOK, gin.H{"client": aboutInfo.Client, "server": aboutInfo.Server})
// }

type client struct {
	Host string `json:"host"`
}

type reaction struct {
	Name string `json:"name"`
	Description string `json:"description"`
}

type action struct {
	Name string `json:"name"`
	Description string `json:"description"`
	Arguments []string `json:"arguments"`
}

type service struct {
	Name string `json:"name"`
	Actions []action `json:"actions"`
	Reactions []reaction `json:"reactions"`
}

type server struct {
	Current_time string `json:"current_time"`
	Services []service `json:"services"`
}

type about struct {
	Client client `json:"client"`
	Server server `json:"server"`
}


// DetailedAbout godoc
// @Summary Get detailed information about the server and services
// @Description Returns detailed information about the client, server, and the services available with their actions and reactions.
// @Tags About
// @Accept  json
// @Produce  json
// @Success 200 {object} about "Detailed information about client and server"
// @Router /about.json [get]
func DetailedAbout(c *gin.Context) {

	var aboutInfo about
	var listServices []string = []string{"email", "lol", "spotify", "github", "discord", "meteo", "tft", "time", "steam"}
	var sqlStatement string

	aboutInfo.Client = client{}
	aboutInfo.Server = server{}
	aboutInfo.Server.Services = []service{}
	aboutInfo.Client.Host = c.ClientIP()
	aboutInfo.Server.Current_time = strconv.FormatInt(time.Now().Unix(), 10)
	for _, a := range listServices {
		var service service
		service.Name = a
		sqlStatement = "SELECT name, description, arguments FROM action WHERE services = $1"
		rows, err := DB_management.DB.Query(sqlStatement, a)
		if err != nil {
			log.Println("err = ", err)
			log.Println("error while getting actions")
			c.IndentedJSON(http.StatusNotAcceptable, gin.H{"message": "error while getting actions"})
			return
		}
		for rows.Next() {
			var action action
			err = rows.Scan(&action.Name, &action.Description, pq.Array(&action.Arguments))
			if err != nil {
				log.Println("a = ", a)
				log.Println("error while getting actions")
				c.IndentedJSON(http.StatusNotAcceptable, gin.H{"message": "error while getting actions"})
				return
			}
			service.Actions = append(service.Actions, action)
		}
		defer rows.Close()
		sqlStatement = "SELECT name, description FROM reactions WHERE services = $1"
		rows, err = DB_management.DB.Query(sqlStatement, a)
		if err != nil {
			log.Println("error while getting reactions")
			c.IndentedJSON(http.StatusNotAcceptable, gin.H{"message": "error while getting reactions"})
			return
		}
		for rows.Next() {
			var reaction reaction
			err = rows.Scan(&reaction.Name, &reaction.Description)
			if err != nil {
				log.Println("error while getting reactions")
				c.IndentedJSON(http.StatusNotAcceptable, gin.H{"message": "error while getting reactions"})
				return
			}
			service.Reactions = append(service.Reactions, reaction)
		}
		defer rows.Close()
		aboutInfo.Server.Services = append(aboutInfo.Server.Services, service)
	}
	c.IndentedJSON(http.StatusOK, gin.H{"client": aboutInfo.Client, "server": aboutInfo.Server})
}