package Meteo

import (
	"API/DB_management"
	"API/call_API"
	"log"
	"errors"
)

type AreaMeteo struct {
	ID int `json:"id"`
	UserID int `json:"user_id"`
	Name string `json:"name"`
	Color string `json:"color"`
	Action string `json:"action"`
	Reaction string `json:"reaction"`
	Toggled bool `json:"toggled"`
	Time_execute string `json:"time_execute"`
	Action_name string `json:"action_name"`
	Reaction_name string `json:"reaction_name"`
	City string `json:"city"`
}

type MeteoAction struct {
	Name string `json:"name"`
	Action func(string) (call_API.MeteoStruct, error) `json:"action"`
}

var listAction = []MeteoAction{
	{Name: "daily_weather", Action: call_API.CallCheckMeteo},
}

func CheckActionCreation(action DB_management.Area) error {
	err := DB_management.VerifyValidServices(action)
	if err != nil {
		return err
	}
	return meteoAREACreation(action)
}

func meteoAREACreation(action DB_management.Area) error {
	sqlStatement := `SELECT id FROM areas_meteo WHERE user_id = $1 AND action_name = $2 AND reaction_name = $3 AND city = $4`
	// var id int
	// err := DB_management.DB.QueryRow(sqlStatement, action.UserID, action.Action_name, action.Reaction_name).Scan(&id)
	// if err == nil {
	// 	return errors.New("area already exist")
	// }

	action.Time_execute = "*"
	if action.Action_name == "daily_weather" {
		action.Time_execute = "00:00"
	}
	sqlStatement = `INSERT INTO areas_meteo (name, color, reaction, toggled, user_id, time_execute, action_name, reaction_name, city) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`
	_, err := DB_management.DB.Exec(sqlStatement, action.Name, action.Color, action.Reaction, action.Toggled, action.UserID, action.Time_execute, action.Action_name, action.Reaction_name, action.City)
	if err != nil {
		log.Println("error = ", err)
		log.Println("error while creating area")
		return errors.New("error while creating area")
	}
	log.Println("area meteo successfully created")
	return nil
}

func CheckActionExecution(action DB_management.Area) (call_API.CallReactionStruct, error) {
	var city string
	var user_id int
	sqlStatement := `SELECT city, user_id FROM areas_meteo WHERE id = $1`
	err := DB_management.DB.QueryRow(sqlStatement, action.ID).Scan(&city, &user_id)
	if err != nil {
		log.Println("error while getting city")
		log.Println("error = ", err)
		return call_API.CallReactionStruct{}, errors.New("error while getting city")
	}
	var github_email string
	var github_token string
	var github_pseudo string
	sqlStatement = `SELECT github_email, github_token, github_pseudo FROM users WHERE id = $1`
	err = DB_management.DB.QueryRow(sqlStatement, user_id).Scan(&github_email, &github_token, &github_pseudo)
	for _, a := range listAction {
		if a.Name == action.Action_name {
			result, err := a.Action(city)
			if err != nil {
				return call_API.CallReactionStruct{}, err
			}
			foremail := call_API.EmailStruct{UserID: action.UserID, Subject: result.EmailSubject, Content: result.EmailContent}
			forgithub := call_API.GithubStruct{Owner: github_pseudo, RepoName: "AREA", FileName: result.GithubFile, Content: result.EmailContent, GithubToken: github_token, Email: github_email}
			return call_API.CallReactionStruct{AreaInfo: action, Meteo: result, Email: foremail, Github: forgithub}, nil
		}
	}
	return call_API.CallReactionStruct{}, errors.New("unknown action")
}

func CheckActionUpdate(action DB_management.Area) error {
	err := DB_management.VerifyValidServices(action)
	if err != nil {
		return err
	}
	return meteoUpdate(action)
}

func meteoUpdate(action DB_management.Area) error {
	sqlStatement := `UPDATE areas_meteo SET name = $1, color = $2, toggled = $3, city = $4 WHERE id = $5`
	_, err := DB_management.DB.Exec(sqlStatement, action.Name, action.Color, action.Toggled, action.City, action.ID)
	if err != nil {
		log.Println("error = ", err)
		log.Println("error while modifying AREA")
		return errors.New("error while modifying AREA")
	}
	return nil
}

func GetAllActionFromUser(user_id int) ([]interface{}, error) {
	sqlStatement := `SELECT id, name, color, reaction, toggled, user_id, time_execute, action_name, reaction_name, city FROM areas_meteo WHERE user_id = $1`
	rows, err := DB_management.DB.Query(sqlStatement, user_id)
	if err != nil {
		log.Println("error while getting AREA")
		return nil, errors.New("error while getting AREA please try again")
	}
	var listArea []interface{}
	for rows.Next() {
		var newArea AreaMeteo
		err = rows.Scan(&newArea.ID, &newArea.Name, &newArea.Color, &newArea.Reaction, &newArea.Toggled, &newArea.UserID, &newArea.Time_execute, &newArea.Action_name, &newArea.Reaction_name, &newArea.City)
		if err != nil {
			log.Println("error while getting AREA")
			return nil, errors.New("error while getting AREA please try again")
		}
		newArea.Action = "meteo"
		listArea = append(listArea, newArea)
	}
	defer rows.Close()
	return listArea, nil
}