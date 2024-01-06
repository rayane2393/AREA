package Github

import (
	"errors"
	"API/call_API"
	"API/DB_management"
	"log"
)

type reaction struct {
	name string `json:"name"`
	reaction func(call_API.GithubStruct) error `json:"reaction"`
}

type action struct {
	name string `json:"name"`
	action func(string, string, string, int) (call_API.GithubStruct, error) `json:"action"`
}

var listAction = []action {
	{name: "new_push", action: call_API.CallCheckGithubPush},
}

var listReaction = []reaction {
	{name: "fill_file", reaction: call_API.CreateFileGithub},
	{name: "create_repo", reaction: call_API.CreateRepoGithub},
}

func CheckActionCreation(action DB_management.Area) error {
	err := DB_management.VerifyValidServices(action)
	if err != nil {
		return err
	}
	return githubAREACreation(action)
}

func githubAREACreation(action DB_management.Area) error {
	sqlStatement := `SELECT github_token FROM users WHERE id=$1;`
	var token string
	err := DB_management.DB.QueryRow(sqlStatement, action.UserID).Scan(&token)
	if err != nil {
		log.Println("err = ", err)
		return errors.New("error while getting github token")
	}
	sqlStatement = `SELECT id FROM areas_github WHERE user_id = $1 AND action_name = $2 AND reaction_name = $3`
	// var id int
	// err := DB_management.DB.QueryRow(sqlStatement, action.UserID, action.Action_name, action.Reaction_name).Scan(&id)
	// if err == nil {
	// 	return errors.New("area already exist")
	// }

	action.Time_execute = "*"
	sqlStatement = `INSERT INTO areas_github (name, color, reaction, toggled, user_id, time_execute, action_name, reaction_name, repo_name) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`
	_, err = DB_management.DB.Exec(sqlStatement, action.Name, action.Color, action.Reaction, action.Toggled, action.UserID, action.Time_execute, action.Action_name, action.Reaction_name, action.Repo_name)
	if err != nil {
		log.Println("error = ", err)
		return errors.New("error while creating area")
	}
	return nil
}

func CheckActionExecution(action DB_management.Area) (call_API.CallReactionStruct, error) {
	type GithubStruct struct {
		GithubToken *string `json:"github_token"`
		GithubEmail *string `json:"github_email"`
		GithubPseudo *string `json:"github_pseudo"`
	}

	var user_id int
	var repo_name string
	sqlStatement := `SELECT user_id, repo_name FROM areas_github WHERE id = $1`
	err := DB_management.DB.QueryRow(sqlStatement, action.ID).Scan(&user_id, &repo_name)
	if err != nil {
		log.Println("error while getting user id")
		log.Println("error = ", err)
		return call_API.CallReactionStruct{}, errors.New("error while getting user id")
	}
	var github GithubStruct
	sqlStatement = `SELECT github_email, github_token, github_pseudo FROM users WHERE id = $1`
	err = DB_management.DB.QueryRow(sqlStatement, user_id).Scan(&github.GithubEmail, &github.GithubToken, &github.GithubPseudo)
	if err != nil {
		return call_API.CallReactionStruct{}, err
	}
	for _, a := range listAction {
		if a.name == action.Action_name {
			result, err := a.action(*github.GithubPseudo, repo_name, *github.GithubToken, action.ID)
			if err != nil {
				return call_API.CallReactionStruct{}, err
			}
			foremail := call_API.EmailStruct{UserID: action.UserID, Subject: result.EmailSubject, Content: result.Content}
			forgithub := call_API.GithubStruct{Owner: *github.GithubPseudo, RepoName: "AREA", FileName: result.GithubFile, Content: result.Content, GithubToken: *github.GithubToken, Email: *github.GithubEmail}
			return call_API.CallReactionStruct{AreaInfo: action, Email: foremail, Github: forgithub}, nil
		}
	}
	return call_API.CallReactionStruct{}, errors.New("unknown reaction")
}

func CheckReactionExecution(reaction call_API.CallReactionStruct) error {
	for _, a := range listReaction {
		if a.name == reaction.AreaInfo.Reaction_name {
			// log.Println("reaction email = ", reaction.Email)
			return a.reaction(reaction.Github)
		}
	}
	return errors.New("unknown reaction")
}

func CheckActionUpdate(action DB_management.Area) error {
	err := DB_management.VerifyValidServices(action)
	if err != nil {
		return err
	}
	return githubUpdate(action)
}

func githubUpdate(action DB_management.Area) error {
	sqlStatement := `UPDATE areas_github SET name = $1, color = $2, toggled = $3, repo_name = $4 WHERE id = $5`
	_, err := DB_management.DB.Exec(sqlStatement, action.Name, action.Color, action.Toggled, action.Repo_name, action.ID)
	if err != nil {
		log.Println("error while modifying AREA")
		return errors.New("error while modifying AREA")
	}
	return nil
}

func GetAllActionFromUser(user_id int) ([]interface{}, error) {
	sqlStatement := `SELECT id, name, color, reaction, toggled, user_id, time_execute, action_name, reaction_name, repo_name FROM areas_github WHERE user_id = $1`
	rows, err := DB_management.DB.Query(sqlStatement, user_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var actions []interface{}
	for rows.Next() {
		var action DB_management.Area
		err = rows.Scan(&action.ID, &action.Name, &action.Color, &action.Reaction, &action.Toggled, &action.UserID, &action.Time_execute, &action.Action_name, &action.Reaction_name, &action.Repo_name)
		if err != nil {
			return nil, err
		}
		action.Action = "github"
		actions = append(actions, action)
	}
	return actions, nil
}
