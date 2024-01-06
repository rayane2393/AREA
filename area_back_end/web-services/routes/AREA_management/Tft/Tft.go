package Tft

import (
	"API/DB_management"
	"API/call_API"
	"log"
	"errors"
	"net/http"
	"io"
	"encoding/json"
	"os"
)

type AreaTft struct {
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
	Summoner_name string `json:"summoner_name"`
}

type TftAction struct {
	Name string `json:"name"`
	Action func(string, int) (call_API.TftStruct, error) `json:"action"`
}

var listAction = []TftAction{
	{Name: "new_rank", Action: call_API.CallCheckTftEloChange},
	{Name: "new_win", Action: call_API.CallCheckTftWinChange},
	{Name: "new_lose", Action: call_API.CallCheckTftLoseChange},
}

func CheckActionCreation(action DB_management.Area) error {
	err := DB_management.VerifyValidServices(action)
	if err != nil {
		return err
	}
	return tftAREACreation(action)
}

func tftAREACreation(action DB_management.Area) error {
	type TftRespBodyPlayer struct {
		ID string `json:"id"`
		AccountID string `json:"accountId"`
		Puuid string `json:"puuid"`
	}

	type TftRespBodyPlayerInfo struct {
		QueueType string `json:"queueType"`
		Tier string `json:"tier"`
		Rank string `json:"rank"`
		Wins int `json:"wins"`
		Losses int `json:"losses"`
	}

	sqlStatement := `SELECT id FROM areas_tft WHERE user_id = $1 AND action_name = $2 AND reaction_name = $3`
	// var id int
	// err := DB_management.DB.QueryRow(sqlStatement, action.UserID, action.Action_name, action.Reaction_name).Scan(&id)
	// if err == nil {
	// 	return errors.New("area already exist")
	// }

	var data = `https://euw1.api.riotgames.com/lol/summoner/v4/summoners/by-name/` + action.Summoner_name + "?api_key=" + os.Getenv("RIOT_API_KEY")
	req, err := http.NewRequest("GET", data, nil)
	if err != nil {
		log.Println("error = ", err)
		return errors.New("error while creating new request for summoner info")
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("error = ", err)
		return errors.New("error while getting summoner info")
	}
	if resp.StatusCode != 200 {
		return errors.New("summoner not found")
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("error = ", err)
		return errors.New("error while reading summoner info")
	}
	var summoner TftRespBodyPlayer
	json.Unmarshal(bodyText, &summoner)
	data = `https://euw1.api.riotgames.com/tft/league/v1/entries/by-summoner/` + summoner.ID + "?api_key=" + os.Getenv("RIOT_API_KEY")
	req, err = http.NewRequest("GET", data, nil)
	if err != nil {
		log.Println("error = ", err)
		return errors.New("error while creating new request for summoner info")
	}
	resp, err = client.Do(req)
	if err != nil {
		log.Println("error = ", err)
		return errors.New("error while getting summoner info")
	}
	defer resp.Body.Close()
	bodyText, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Println("error = ", err)
		return errors.New("error while reading summoner info")
	}
	var summoner_info []TftRespBodyPlayerInfo
	json.Unmarshal(bodyText, &summoner_info)
	var summoner_rank string
	var summoner_win int
	var summoner_lose int
	for _, info := range summoner_info {
		if info.QueueType == "RANKED_TFT" {
			summoner_rank = info.Tier + " " + info.Rank
			summoner_win = info.Wins
			summoner_lose = info.Losses
		}
	}
	action.Time_execute = "*"
	sqlStatement = `INSERT INTO areas_tft (name, color, reaction, toggled, user_id, time_execute, action_name, reaction_name, summoner_name, in_game, summoner_id, summoner_account_id, summoner_puuid, elo, win, lose) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16)`
	_, err = DB_management.DB.Exec(sqlStatement, action.Name, action.Color, action.Reaction, action.Toggled, action.UserID, action.Time_execute, action.Action_name, action.Reaction_name, action.Summoner_name, false, summoner.ID, summoner.AccountID, summoner.Puuid, summoner_rank, summoner_win, summoner_lose)
	if err != nil {
		log.Println("error = ", err)
		log.Println("error while creating area")
		return errors.New("error while creating area")
	}
	return nil
}

func CheckActionExecution(action DB_management.Area) (call_API.CallReactionStruct, error) {
	var summoner_name string
	var user_id int
	sqlStatement := `SELECT summoner_name, user_id FROM areas_tft WHERE id = $1`
	err := DB_management.DB.QueryRow(sqlStatement, action.ID).Scan(&summoner_name, &user_id)

	if err != nil {
		log.Println("error while getting summoner name")
		log.Println("error = ", err)
		return call_API.CallReactionStruct{}, errors.New("error while getting summoner name")
	}
	var github_email string
	var github_token string
	var github_pseudo string
	sqlStatement = `SELECT github_email, github_token, github_pseudo FROM users WHERE id = $1`
	err = DB_management.DB.QueryRow(sqlStatement, user_id).Scan(&github_email, &github_token, &github_pseudo)
	for _, a := range listAction {
		if a.Name == action.Action_name {
			result, err := a.Action(summoner_name, action.ID)
			if err != nil {
				return call_API.CallReactionStruct{}, err
			}
			foremail := call_API.EmailStruct{UserID: action.UserID, Subject: result.EmailSubject, Content: result.EmailContent}
			forgithub := call_API.GithubStruct{Owner: github_pseudo, RepoName: "AREA", FileName: result.GithubFile, Content: result.EmailContent, GithubToken: github_token, Email: github_email}
			return call_API.CallReactionStruct{AreaInfo: action, Tft: result, Email: foremail, Github: forgithub}, nil
		}
	}
	return call_API.CallReactionStruct{}, errors.New("unknown action")
}

func CheckActionUpdate(action DB_management.Area) error {
	err := DB_management.VerifyValidServices(action)
	if err != nil {
		return err
	}
	return tftUpdate(action)
}

func tftUpdate(action DB_management.Area) error {
	sqlStatement := `UPDATE areas_tft SET name = $1, color = $2, toggled = $3, summoner_name = $4 WHERE id = $5`
	_, err := DB_management.DB.Exec(sqlStatement, action.Name, action.Color, action.Toggled, action.Summoner_name, action.ID)
	if err != nil {
		log.Println("error while modifying AREA")
		return errors.New("error while modifying AREA")
	}
	return nil
}

func GetAllActionFromUser(user_id int) ([]interface{}, error) {
	sqlStatement := `SELECT id, name, color, reaction, toggled, user_id, time_execute, action_name, reaction_name, summoner_name FROM areas_tft WHERE user_id = $1`
	rows, err := DB_management.DB.Query(sqlStatement, user_id)
	if err != nil {
		log.Println("error while getting AREA")
		return nil, errors.New("error while getting AREA please try again")
	}
	var listArea []interface{}
	for rows.Next() {
		var newArea AreaTft
		err = rows.Scan(&newArea.ID, &newArea.Name, &newArea.Color, &newArea.Reaction, &newArea.Toggled, &newArea.UserID, &newArea.Time_execute, &newArea.Action_name, &newArea.Reaction_name, &newArea.Summoner_name)
		if err != nil {
			log.Println("error while getting AREA")
			return nil, errors.New("error while getting AREA please try again")
		}
		newArea.Action = "tft"
		listArea = append(listArea, newArea)
	}
	defer rows.Close()
	return listArea, nil
}