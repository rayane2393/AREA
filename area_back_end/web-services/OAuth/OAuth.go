package OAuth

import (
	"net/http"
	"io/ioutil"
	"os"
	"log"
	"encoding/json"
	"API/DB_management"
	"bytes"
	"errors"
	"net/url"
	"fmt"
)

func Github(code string, token string) (DB_management.User, error) {
	type githubAccessTokenResponse struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
		Scope       string `json:"scope"`
	}

	type githubUserResponse struct {
		Login string `json:"login"`
		ID int `json:"id"`
		Email string `json:"email"`
	}

	type githubEmailResponse struct {
		Email string `json:"email"`
		Primary bool `json:"primary"`
		Verified bool `json:"verified"`
		Visibility string `json:"visibility"`
	}

	requestBodyMap := map[string]string{
        "client_id": os.Getenv("GITHUB_ID"),
        "client_secret": os.Getenv("GITHUB_SECRET"),
        "code": code,
    }
    requestJSON, _ := json.Marshal(requestBodyMap)

    req, err := http.NewRequest(
        "POST",
        "https://github.com/login/oauth/access_token",
        bytes.NewBuffer(requestJSON),
    )
    if err != nil {
        log.Panic("Request creation failed")
		return DB_management.User{}, errors.New("Error while creating account please try again")
    }
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Accept", "application/json")

    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        log.Panic("Request failed")
		return DB_management.User{}, errors.New("Error while creating account please try again")
    }

    respbody, _ := ioutil.ReadAll(resp.Body)

    var ghresp githubAccessTokenResponse
    json.Unmarshal(respbody, &ghresp)

    req, err = http.NewRequest(
        "GET",
        "https://api.github.com/user",
        nil,
    )
    if err != nil {
        log.Panic("API Request creation failed")
		return DB_management.User{}, errors.New("Error while creating account please try again")
    }

    authorizationHeaderValue := fmt.Sprintf("token %s", ghresp.AccessToken)
    req.Header.Set("Authorization", authorizationHeaderValue)

    resp, resperr := http.DefaultClient.Do(req)
    if resperr != nil {
        log.Panic("Request failed")
		return DB_management.User{}, errors.New("Error while creating account please try again")
    }

    respbody, _ = ioutil.ReadAll(resp.Body)

	var githubUser githubUserResponse

	json.Unmarshal(respbody, &githubUser)

	req, err = http.NewRequest(
		"GET",
        "https://api.github.com/user/emails",
        nil,
    )
    if err != nil {
		log.Panic("API Request creation failed")
		return DB_management.User{}, errors.New("Error while creating account please try again")
    }

    authorizationHeaderValue = fmt.Sprintf("Bearer %s", ghresp.AccessToken)
    req.Header.Set("Authorization", authorizationHeaderValue)
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")

    resp, resperr = http.DefaultClient.Do(req)
    if resperr != nil {
        log.Panic("Request failed")
		return DB_management.User{}, errors.New("Error while creating account please try again")
    }

    respbody, _ = ioutil.ReadAll(resp.Body)

	var githubEmail []githubEmailResponse
	json.Unmarshal(respbody, &githubEmail)

	for _, a := range githubEmail {
		if a.Primary == true && a.Verified == true {
			githubUser.Email = a.Email
		}
	}

	if githubUser.ID == 0 {
		log.Println("error while getting github user")
		log.Println("respbody = ", string(respbody))
		return DB_management.User{}, errors.New("Error while getting github accout please try again")
	}

	id, err := DB_management.GetIntFromDB(`SELECT id FROM users WHERE github_id = $1`, githubUser.ID)
	log.Println("id = ", id)
	if id == 0 && token == "" {
		var newUsers DB_management.User
		newUsers = DB_management.User{Username: githubUser.Login, Github_id: &githubUser.ID, Github_token: &ghresp.AccessToken, Email: &githubUser.Email}
		if DB_management.CreateUser(`INSERT INTO users (username, github_id, github_token, email, github_email, github_pseudo) VALUES($1, $2, $3, $4, $5, $6)`, newUsers.Username, *newUsers.Github_id, *newUsers.Github_token, githubUser.Email, githubUser.Email, newUsers.Username) != nil {
			log.Println("error while creating user with github")
			return DB_management.User{}, errors.New("Error while creating account please try again")
		} else {
			log.Println("user created and send correctly from github")
			DB_management.Users = append(DB_management.Users, newUsers)
		}
	} else if id == 0 && token != "" {
		var newUsers DB_management.User
		newUsers = DB_management.User{Github_id: &githubUser.ID, Github_token: &ghresp.AccessToken, Github_email: &githubUser.Email, Github_pseudo: &githubUser.Login}
		sqlStatement := `UPDATE users SET github_id = $1, github_token = $2, github_email = $3, github_pseudo = $4 WHERE token = $5`
		_, err = DB_management.DB.Exec(sqlStatement, *newUsers.Github_id, *newUsers.Github_token, *newUsers.Github_email, *newUsers.Github_pseudo, token)
		if err != nil {
			log.Println("err = ", err)
			log.Println("error while adding user with github")
			return DB_management.User{}, errors.New("Error while creating account please try again")
		} else {
			log.Println("user added and send correctly from github")
			for i, a := range DB_management.Users {
				if *a.Token == token {
					DB_management.Users[i].Github_id = newUsers.Github_id
					DB_management.Users[i].Github_token = newUsers.Github_token
					DB_management.Users[i].Github_email = newUsers.Github_email
					DB_management.Users[i].Github_pseudo = newUsers.Github_pseudo
				}
			}
		}
	} else if id != 0 && token != "" {
		log.Println("user already exist")
		return DB_management.User{}, errors.New("User already logged with github")
	}
	for _, a := range DB_management.Users {
		if a.Github_id != nil && *a.Github_id == githubUser.ID{
			return a, nil
		}
	}
	return DB_management.User{}, errors.New("Error while creating account please try again")
}

func Discord(code string, token string) (DB_management.User, error) {
	tokenEndpoint := "https://discord.com/api/oauth2/token"

	clientID := os.Getenv("DISCORD_ID")
	clientSecret := os.Getenv("DISCORD_SECRET")
	redirectURI := "https://google.com"
	authCode := code

	data := url.Values{}
	data.Set("client_id", clientID)
	data.Set("client_secret", clientSecret)
	data.Set("grant_type", "authorization_code")
	data.Set("code", authCode)
	data.Set("redirect_uri", redirectURI)

	resp, err := http.PostForm(tokenEndpoint, data)
	if err != nil {
		log.Println("Error exchanging authorization code:", err)
		return DB_management.User{}, errors.New("Error while creating account please try again")
	}
	defer resp.Body.Close()

	var tokenResponse map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&tokenResponse)
	if err != nil {
		log.Println("Error decoding token response:", err)
		return DB_management.User{}, errors.New("Error while creating account please try again")
	}

	accessToken, found := tokenResponse["access_token"].(string)
	if !found {
		log.Println("Access token not found in response")
		return DB_management.User{}, errors.New("Error while creating account please try again")
	}

	userInfoURL := "https://discord.com/api/users/@me"
	req, err := http.NewRequest("GET", userInfoURL, nil)
	if err != nil {
		log.Println("Error creating request for user info:", err)
		return DB_management.User{}, errors.New("Error while creating account please try again")
	}
	req.Header.Set("Authorization", "Bearer "+accessToken)

	client := &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		log.Println("Error getting user info:", err)
		return DB_management.User{}, errors.New("Error while creating account please try again")
	}
	defer resp.Body.Close()

	var userInfo map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&userInfo)
	if err != nil {
		log.Println("Error decoding user info response:", err)
		return DB_management.User{}, errors.New("Error while creating account please try again")
	}

	user_id, found := userInfo["id"].(string)
	if !found {
		log.Println("User ID not found in response")
		return DB_management.User{}, errors.New("Error while creating account please try again")
	}
	username, found := userInfo["username"].(string)
	if !found {
		log.Println("Username not found in response")
		return DB_management.User{}, errors.New("Error while creating account please try again")
	}
	sqlStatement := `SELECT id FROM users WHERE discord_id = $1`
	id, err := DB_management.GetStringFromDB(sqlStatement, user_id)
	if id == "" && token == "" {
		sqlStatement := `INSERT INTO users (username, discord_id, discord_token) VALUES($1, $2, $3)`
		_, err = DB_management.DB.Exec(sqlStatement, username, user_id, accessToken)
		if err != nil {
			log.Println("err = ", err)
			log.Println("error while creating user with discord")
			return DB_management.User{}, errors.New("Error while creating account please try again")
		} else {
			log.Println("err = ", err)
			var newUsers DB_management.User
			newUsers = DB_management.User{Username: username, Discord_id: &user_id, Discord_token: &accessToken}
			DB_management.Users = append(DB_management.Users, newUsers)
			log.Println("user created and send correctly from discord")
			log.Println("newUsers = ", newUsers)
		}
	} else if id == "" && token != "" {
		sqlStatement := `UPDATE users SET discord_id = $1, discord_token = $2 WHERE token = $3`
		_, err = DB_management.DB.Exec(sqlStatement, user_id, accessToken, token)
		if err != nil {
			log.Println("err = ", err)
			log.Println("error while adding user with discord")
			return DB_management.User{}, errors.New("Error while creating account please try again")
		} else {
			log.Println("user added and send correctly from discord")
			for i, a := range DB_management.Users {
				if a.Token != nil && *a.Token == token {
					DB_management.Users[i].Discord_id = &user_id
					DB_management.Users[i].Discord_token = &accessToken
				}
			}
		}
	} else if id != "" && token != "" {
		log.Println("user already exist")
		return DB_management.User{}, errors.New("User already logged with discord")
	}
	for _, a := range DB_management.Users {
		if a.Discord_id != nil && *a.Discord_id == user_id {
			return a, nil
		}
	}
	return DB_management.User{}, errors.New("Error while creating account please try again")
}
