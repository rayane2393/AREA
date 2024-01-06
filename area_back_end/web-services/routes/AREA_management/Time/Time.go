package Time

import (
	"API/DB_management"
	"API/call_API"
	"log"
	"errors"
	"time"
	"strings"
	"strconv"
)

type AreaTime struct {
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
	Title string `json:"title"`
	Content string `json:"content"`
}

func CheckActionCreation(action DB_management.Area) error {
	err := DB_management.VerifyValidServices(action)
	if err != nil {
		return err
	}
	return timeAREACreation(action)
}

func timeAREACreation(action DB_management.Area) error {
	if action.Interval != "" {
		var hour int
		var min int
		var err error
		if action.Interval == "*" {
			hour = 0
			min = 1
		} else {
			if strings.Split(action.Interval, ":")[0] == "*" {
				hour = 0
			} else {
				hour, err = strconv.Atoi(strings.Split(action.Interval, ":")[0])
				if err != nil {
					log.Println("error = ", err)
					return errors.New("error while creating area")
				}
			}
			if strings.Split(action.Interval, ":")[1] == "*" {
				min = 0
			} else {
				min, err = strconv.Atoi(strings.Split(action.Interval, ":")[1])
				if err != nil {
					log.Println("error = ", err)
					return errors.New("error while creating area")
				}
			}
		}
		action.Time_execute = time.Now().Add(time.Hour * time.Duration(hour) + time.Minute * time.Duration(min)).Format("15:04")
	}
	sqlStatement := `INSERT INTO areas_time (name, color, reaction, toggled, user_id, time_execute, action_name, reaction_name, title, content, interval) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`
	_, err := DB_management.DB.Exec(sqlStatement, action.Name, action.Color, action.Reaction, action.Toggled, action.UserID, action.Time_execute, action.Action_name, action.Reaction_name, action.Title, action.Content, action.Interval)
	if err != nil {
		log.Println("error = ", err)
		log.Println("error while creating area")
		return errors.New("error while creating area")
	}
	log.Println("area time successfully created")
	return nil
}

func CheckActionExecution(action DB_management.Area) (call_API.CallReactionStruct, error) {
	var title string
	var content string
	var user_id int
	var interval string
	sqlStatement := `SELECT title, content, user_id, interval FROM areas_time WHERE id = $1`
	err := DB_management.DB.QueryRow(sqlStatement, action.ID).Scan(&title, &content, &user_id, &interval)
	if err != nil {
		log.Println("error while getting title and content")
		log.Println("error = ", err)
		return call_API.CallReactionStruct{}, errors.New("error while getting title and content")
	}
	if action.Action_name == "repeat" {
		var hour int
		var min int
		var err error
		if interval == "*" {
			hour = 0
			min = 1
		} else {
			if strings.Split(interval, ":")[0] == "*" {
				hour = 0
			} else {
				hour, err = strconv.Atoi(strings.Split(interval, ":")[0])
				if err != nil {
					log.Println("error = ", err)
					return call_API.CallReactionStruct{}, errors.New("error while creating area")
				}
			}
			if strings.Split(interval, ":")[1] == "*" {
				min = 0
			} else {
				min, err = strconv.Atoi(strings.Split(interval, ":")[1])
				if err != nil {
					log.Println("error = ", err)
					return call_API.CallReactionStruct{}, errors.New("error while creating area")
				}
			}
		}
		action.Time_execute = time.Now().Add(time.Hour * time.Duration(hour) + time.Minute * time.Duration(min)).Format("15:04")
		sqlStatement = `UPDATE areas_time SET time_execute = $1 WHERE id = $2`
		_, err = DB_management.DB.Exec(sqlStatement, action.Time_execute, action.ID)
		if err != nil {
			log.Println("error = ", err)
			log.Println("error while modifying AREA")
			return call_API.CallReactionStruct{}, errors.New("error while modifying AREA")
		}
	}
	var github_email_db (*string)
	var github_email string
	var github_token_db (*string)
	var github_token string
	var github_pseudo_db (*string)
	var github_pseudo string
	sqlStatement = `SELECT github_email, github_token, github_pseudo FROM users WHERE id = $1`
	err = DB_management.DB.QueryRow(sqlStatement, user_id).Scan(&github_email_db, &github_token_db, &github_pseudo_db)
	if github_email_db == nil || github_token_db == nil || github_pseudo_db == nil {
		github_email = ""
		github_token = ""
		github_pseudo = ""
	} else {
		github_email = *github_email_db
		github_token = *github_token_db
		github_pseudo = *github_pseudo_db
	}
	result := call_API.TimeStruct{Title: title, Content: content}
	if err != nil {
		return call_API.CallReactionStruct{}, err
	}
	foremail := call_API.EmailStruct{UserID: user_id, Subject: title, Content: content}
	forgithub := call_API.GithubStruct{Owner: github_pseudo, RepoName: "AREA", FileName: title, Content: content, GithubToken: github_token, Email: github_email}
	return call_API.CallReactionStruct{AreaInfo: action, Time: result, Email: foremail, Github: forgithub}, nil
}

func CheckActionUpdate(action DB_management.Area) error {
	err := DB_management.VerifyValidServices(action)
	if err != nil {
		return err
	}
	return timeUpdate(action)
}

func timeUpdate(action DB_management.Area) error {
	sqlStatement := `UPDATE areas_time SET name = $1, color = $2, toggled = $3, time_execute = $4, title = $5, content = $6 WHERE id = $7`
	_, err := DB_management.DB.Exec(sqlStatement, action.Name, action.Color, action.Toggled, action.Time_execute, action.Title, action.Content, action.ID)
	if err != nil {
		log.Println("error = ", err)
		log.Println("error while modifying AREA")
		return errors.New("error while modifying AREA")
	}
	return nil
}

func GetAllActionFromUser(user_id int) ([]interface{}, error) {
	sqlStatement := `SELECT id, name, color, reaction, toggled, user_id, time_execute, action_name, reaction_name, time_execute, title, content FROM areas_time WHERE user_id = $1`
	rows, err := DB_management.DB.Query(sqlStatement, user_id)
	if err != nil {
		log.Println("error while getting AREA")
		return nil, errors.New("error while getting AREA please try again")
	}
	var listArea []interface{}
	for rows.Next() {
		var newArea AreaTime
		err = rows.Scan(&newArea.ID, &newArea.Name, &newArea.Color, &newArea.Reaction, &newArea.Toggled, &newArea.UserID, &newArea.Time_execute, &newArea.Action_name, &newArea.Reaction_name, &newArea.Time_execute, &newArea.Title, &newArea.Content)
		if err != nil {
			log.Println("error while getting AREA")
			return nil, errors.New("error while getting AREA please try again")
		}
		newArea.Action = "time"
		listArea = append(listArea, newArea)
	}
	defer rows.Close()
	return listArea, nil
}