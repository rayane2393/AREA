package Spotify

import (
	"API/DB_management"
	"API/call_API"

	"errors"
	// "fmt"
	"log"
)

type AreaSpotify struct {
	ID int `json:"id"`
	UserID int `json:"user_id"`
	Name string `json:"name"`
	Color string `json:"color"`
	Reaction string `json:"reaction"`
	Action string `json:"action"`
	Toggled bool `json:"toggled"`
	Time_execute string `json:"time_execute"`
	Action_name string `json:"action_name"`
	Reaction_name string `json:"reaction_name"`
	Artist_ID string `json:"artist_id"`
}

type SpotifyAction struct {
	Name string `json:"new_song"`
	Action func(int, string) (call_API.SpotifyStruct, error) `json:"action"`
}

var listAction = []SpotifyAction{
	{Name: "new_song", Action: call_API.CallCheckSpotifyRelease},
	{Name: "new_album_artist", Action: call_API.CallCheckNewAlbumArtist},
}

func CheckActionCreation(action DB_management.Area) error {
	err := DB_management.VerifyValidServices(action)
	if err != nil {
		return err
	}
	return spotifyAREACreation(action)
}

func spotifyAREACreation(action DB_management.Area) error {
	sqlStatement := `SELECT id FROM areas_spotify WHERE user_id = $1 AND action_name = $2 AND reaction_name = $3`
	// var id int
	// err := DB_management.DB.QueryRow(sqlStatement, action.UserID, action.Action_name, action.Reaction_name).Scan(&id)
	// if err == nil {
	// 	return errors.New("area already exist")
	// }

	action.Time_execute = "*"
	sqlStatement = `INSERT INTO areas_spotify (name, color, reaction, toggled, user_id, time_execute, action_name, reaction_name, new_release, artist_id, new_album_name) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`
	_, err := DB_management.DB.Exec(sqlStatement, action.Name, action.Color, action.Reaction, action.Toggled, action.UserID, action.Time_execute, action.Action_name, action.Reaction_name, "", action.Artist_ID, "")
	if err != nil {
		log.Println("error = ", err)
		log.Println("action = ", action)
		log.Println("error while creating area")
		return errors.New("error while creating area")
	}
	return nil
}

func CheckActionExecution(action DB_management.Area) (call_API.CallReactionStruct, error) {
	var user_id int
	sqlStatement := `SELECT user_id FROM areas_spotify WHERE id = $1`
	err := DB_management.DB.QueryRow(sqlStatement, action.ID).Scan(&user_id)
	if err != nil {
		log.Println("error while getting user id")
		log.Println("error = ", err)
		return call_API.CallReactionStruct{}, errors.New("error while getting user id")
	}
	var github_email string
	var github_token string
	var github_pseudo string
	sqlStatement = `SELECT github_email, github_token, github_pseudo FROM users WHERE id = $1`
	err = DB_management.DB.QueryRow(sqlStatement, user_id).Scan(&github_email, &github_token, &github_pseudo)
	for _, a := range listAction {
		if a.Name == action.Action_name {
			sqlStatement := `SELECT artist_id FROM areas_spotify WHERE id = $1`
			var artist_id string
			err := DB_management.DB.QueryRow(sqlStatement, action.ID).Scan(&artist_id)
			if err != nil {
				log.Println("error while getting artist id")
				log.Println("error = ", err)
				return call_API.CallReactionStruct{}, errors.New("error while getting artist id")
			}
			result, err := a.Action(action.ID, artist_id)
			if err != nil {
				return call_API.CallReactionStruct{}, err
			}
			foremail := call_API.EmailStruct{UserID: action.UserID, Subject: result.EmailSubject, Content: result.EmailContent}
			forgithub := call_API.GithubStruct{Owner: github_pseudo, RepoName: "AREA", FileName: result.GithubFile, Content: result.EmailContent, GithubToken: github_token, Email: github_email}
			return call_API.CallReactionStruct{AreaInfo: action, Spotify: result, Email: foremail, Github: forgithub}, nil
		}
	}
	return call_API.CallReactionStruct{}, errors.New("unknown reaction")
}

func CheckActionUpdate(action DB_management.Area) error {
	err := DB_management.VerifyValidServices(action)
	if err != nil {
		return err
	}
	return spotifyUpdate(action)
}

func spotifyUpdate(action DB_management.Area) error {
	sqlStatement := `UPDATE areas_spotify SET name = $1, color = $2, toggled = $3 WHERE id = $4`
	if action.Action_name == "new_album_artist" {
		sqlStatement = `UPDATE areas_spotify SET name = $1, color = $2, toggled = $3, artist_id = $4 WHERE id = $5`
	}
	_, err := DB_management.DB.Exec(sqlStatement, action.Name, action.Color, action.Toggled, action.ID)
	if err != nil {
		log.Println("error while modifying AREA")
		return errors.New("error while modifying AREA")
	}
	return nil
}

func GetAllActionFromUser(user_id int) ([]interface{}, error) {
	sqlStatement := `SELECT id, name, color, reaction, toggled, user_id, time_execute, action_name, reaction_name, artist_id FROM areas_spotify WHERE user_id = $1`
	rows, err := DB_management.DB.Query(sqlStatement, user_id)
	if err != nil {
		log.Println("error while getting AREA")
		return nil, errors.New("error while getting AREA please try again")
	}
	var listArea []interface{}
	for rows.Next() {
		var newArea AreaSpotify
		err = rows.Scan(&newArea.ID, &newArea.Name, &newArea.Color, &newArea.Reaction, &newArea.Toggled, &newArea.UserID, &newArea.Time_execute, &newArea.Action_name, &newArea.Reaction_name, &newArea.Artist_ID)
		if err != nil {
			log.Println("error while getting AREA")
			return nil, errors.New("error while getting AREA please try again")
		}
		newArea.Action = "spotify"
		listArea = append(listArea, newArea)
	}
	defer rows.Close()
	return listArea, nil
}