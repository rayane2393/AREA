package call_API

import (
	// "math/rand"
	// "io/ioutil"
    "net/http"
	"net/smtp"
	"strconv"
	"encoding/json"
	// "github.com/gin-gonic/gin"
	"net/url"
	"errors"

	// "API/API_key"
	"API/DB_management"

	// "fmt"
    "log"
    "os"
	"strings"
	"io"
	"encoding/base64"
	"net"
)

// type Pokemon struct {
// 	Name string `json:"name"`
// 	Sprite map[string]string `json:"sprites"`
// }

type Spotify_token struct {
	Access_token string `json:"access_token"`
	Token_type string `json:"token_type"`
	Expiration_time string `json:"expires_in"`
}

type EmailStruct struct {
	UserID int `json:"user_id"`
	Subject string `json:"subject"`
	Content string `json:"content"`
}

type GithubStruct struct {
	Owner string `json:"owner"`
	RepoName string `json:"repo_name"`
	FileName string `json:"file_name"`
	Content string `json:"content"`
	GithubToken string `json:"github_token"`
	Email string `json:"email"`
	EmailSubject string `json:"email_subject"`
	GithubFile string `json:"github_file"`
}

type SpotifyStruct struct {
	ArtistsName string `json:"artists_name"`
	AlbumName string `json:"album_name"`
	EmailContent string `json:"email_content"`
	EmailSubject string `json:"email_subject"`
	GithubFile string `json:"github_file"`
}

type LolStruct struct {
	SummonerName string `json:"summoner_name"`
	EmailContent string `json:"email_content"`
	EmailSubject string `json:"email_subject"`
	GithubFile string `json:"github_file"`
}

type TftStruct struct {
	SummonerName string `json:"summoner_name"`
	EmailContent string `json:"email_content"`
	EmailSubject string `json:"email_subject"`
	GithubFile string `json:"github_file"`
}

type MeteoStruct struct {
	EmailContent string `json:"email_content"`
	EmailSubject string `json:"email_subject"`
	GithubFile string `json:"github_file"`
}

type TimeStruct struct {
	Title string `json:"title"`
	Content string `json:"content"`
}

type SteamStruct struct {
	EmailContent string `json:"email_content"`
	EmailSubject string `json:"email_subject"`
	GithubFile string `json:"github_file"`
}

type CallReactionStruct struct {
	AreaInfo DB_management.Area `json:"area_info"`
	Email EmailStruct `json:"email"`
	Spotify SpotifyStruct `json:"spotify"`
	Lol LolStruct `json:"lol"`
	Meteo MeteoStruct `json:"meteo"`
	Github GithubStruct `json:"github"`
	Tft TftStruct `json:"tft"`
	Time TimeStruct `json:"time"`
	Steam SteamStruct `json:"steam"`
}

// func Call_Pokemon (c *gin.Context) {
// 	if API_key.CheckApiKey(c) == false {
// 		return
// 	}
// 	random_number := rand.Intn(1292)
// 	var pokemon Pokemon

// 	response, err := http.Get("https://pokeapi.co/api/v2/pokemon/" + strconv.Itoa(random_number))
// 	if err != nil {
//         fmt.Print(err.Error())
//         os.Exit(1)
//     }

//     responseData, err := ioutil.ReadAll(response.Body)
//     if err != nil {
//         log.Fatal(err)
//     }
// 	json.Unmarshal(responseData, &pokemon)
// 	c.IndentedJSON(http.StatusOK, gin.H{"message": "pokemon summon successfully", "name": pokemon.Name, "sprite": pokemon.Sprite["front_default"]})
// }

func getResponseData(method string, link string, data io.Reader, header map[string]string) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, link, data)
	if err != nil {
		log.Println("err = ", err)
		return nil, errors.New("error while creating new request")
	}
	for key, value := range header {
		req.Header.Set(key, value)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("err = ", err)
		return resp, errors.New("error while getting response")
	}
	return resp, nil
}

func callAPI(method string, link string, data io.Reader, header map[string]string) ([]byte, error) {
	resp, err := getResponseData(method, link, data, header)
	if err != nil {
		log.Println("error = ", err)
		return nil, errors.New("error while getting response")
	} else if resp.StatusCode != 200 {
		log.Println("error = ", err)
		return nil, errors.New("error while getting response")
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("error = ", err)
		return nil, errors.New("error while reading response")
	}
	return bodyText, nil
}

func CallCheckSpotifyRelease(areaID int, artist_id string) (SpotifyStruct, error) {
	type AlbumResponse struct {
		Albums struct {
			Items []struct {
				Artists []struct {
					Name string `json:"name"`
				} `json:"artists"`
				Name string `json:"name"`
			} `json:"items"`
		} `json:"albums"`
	}

	var token Spotify_token

	sqlStatement := `SELECT new_release FROM areas_spotify WHERE id = $1`
	var last_spotify_release string
	err := DB_management.DB.QueryRow(sqlStatement, areaID).Scan(&last_spotify_release)
	if err != nil {
		log.Println("error = ", err)
		log.Println("error while getting last spotify release")
		return SpotifyStruct{}, errors.New("error while getting last spotify release")
	}
	method := "POST"
	link := "https://accounts.spotify.com/api/token"
	data := `grant_type=client_credentials&client_id=` + os.Getenv("SPOTIFY_ID") + `&client_secret=`+  os.Getenv("SPOTIFY_SECRET")
	header := map[string]string{"Content-Type": "application/x-www-form-urlencoded"}
	bodyText, err := callAPI(method, link, strings.NewReader(data), header)
	if err != nil {
		log.Println("error = ", err)
		return SpotifyStruct{}, errors.New("error while getting token")
	}
	json.Unmarshal(bodyText, &token)

	method = "GET"
	link = "https://api.spotify.com/v1/browse/new-releases?country=FR&limit=1&offset=5"
	data = ""
	header = map[string]string{"Authorization": "Bearer " + token.Access_token}
	bodyText, err = callAPI(method, link, strings.NewReader(data), header)
	if err != nil {
		log.Println("error = ", err)
		return SpotifyStruct{}, errors.New("error while getting new releases")
	}
	var album AlbumResponse
	json.Unmarshal(bodyText, &album)
	// log.Printf("album = %s, artist = %s\n", album.Albums.Items[0].Name, album.Albums.Items[0].Artists[0].Name)
	if album.Albums.Items[0].Name + " by " + album.Albums.Items[0].Artists[0].Name == last_spotify_release {
		return SpotifyStruct{}, errors.New("no new album release")
	}
	sqlStatement = `UPDATE areas_spotify SET new_release = $1 WHERE id = $2`
	_, err = DB_management.DB.Exec(sqlStatement, album.Albums.Items[0].Name + " by " + album.Albums.Items[0].Artists[0].Name, areaID)
	if err != nil {
		log.Println("error = ", err)
		return SpotifyStruct{}, errors.New("error while updating last spotify release")
	}
	str_result := "New album release\n" + album.Albums.Items[0].Name + " by " + album.Albums.Items[0].Artists[0].Name + " is out now !"
	info_result := SpotifyStruct{ArtistsName: album.Albums.Items[0].Artists[0].Name, AlbumName: album.Albums.Items[0].Name, EmailContent: str_result, EmailSubject: "New album release", GithubFile: "New album release"}
	return info_result, nil
}

func CallCheckNewAlbumArtist (areaID int, artist_id string) (SpotifyStruct, error) {
	type AlbumArtistResponse struct {
		Items []struct {
			Artists []struct {
				Name string `json:"name"`
			}
			Name string `json:"name"`
		} `json:"items"`
	}

	var token Spotify_token

	sqlStatement := `SELECT new_album_name FROM areas_spotify WHERE id = $1 AND action_name = $2`
	var last_artist_release string
	err := DB_management.DB.QueryRow(sqlStatement, areaID, "new_album_artist").Scan(&last_artist_release)
	if err != nil {
		log.Println("error = ", err)
		log.Println("error while getting last album release artist")
		return SpotifyStruct{}, errors.New("error while getting last album release")
	}
	method := "POST"
	link := "https://accounts.spotify.com/api/token"
	data := `grant_type=client_credentials&client_id=` + os.Getenv("SPOTIFY_ID") + `&client_secret=`+  os.Getenv("SPOTIFY_SECRET")
	header := map[string]string{"Content-Type": "application/x-www-form-urlencoded"}
	bodyText, err := callAPI(method, link, strings.NewReader(data), header)
	if err != nil {
		log.Println("error = ", err)
		return SpotifyStruct{}, errors.New("error while getting token")
	}
	json.Unmarshal(bodyText, &token)
	method = "GET"
	link = "https://api.spotify.com/v1/artists/" + artist_id + "/albums/?limit=1"
	data = ""
	header = map[string]string{"Authorization": "Bearer " + token.Access_token}
	bodyText, err = callAPI(method, link, strings.NewReader(data), header)
	if err != nil {
		log.Println("error = ", err)
		return SpotifyStruct{}, errors.New("error while getting new releases")
	}
	var album AlbumArtistResponse
	json.Unmarshal(bodyText, &album)
	if album.Items[0].Name + " by " + album.Items[0].Artists[0].Name == last_artist_release {
		return SpotifyStruct{}, errors.New("no new album release")
	}
	sqlStatement = `UPDATE areas_spotify SET new_album_name = $1 WHERE action_name = $2 AND id = $3`
	_, err = DB_management.DB.Exec(sqlStatement, album.Items[0].Name + " by " + album.Items[0].Artists[0].Name, "new_album_artist", areaID)
	if err != nil {
		log.Println("error = ", err)
		return SpotifyStruct{}, errors.New("error while updating last spotify release")
	}
	str_result := "New album release\n" + album.Items[0].Name + " by " + album.Items[0].Artists[0].Name + " is out now !"
	info_res := SpotifyStruct{ArtistsName: album.Items[0].Artists[0].Name, AlbumName: album.Items[0].Name, EmailContent: str_result, EmailSubject: "New album release from " + album.Items[0].Artists[0].Name, GithubFile: "New album release from " + album.Items[0].Artists[0].Name}
	return info_res, nil
}

func SendMail(sended EmailStruct) error {
	var email string

	sqlStatement := `SELECT email FROM users WHERE id = $1`
	err := DB_management.DB.QueryRow(sqlStatement, sended.UserID).Scan(&email)
	if err != nil {
		log.Println("error while getting email")
		return errors.New("error while getting email")
	}

	auth := smtp.PlainAuth("", os.Getenv("EMAIL"), os.Getenv("EMAIL_PASSWORD"), "smtp.gmail.com")
	from := os.Getenv("EMAIL")
	to := []string{string(email)}
	msg := []byte("From: " + from + "\n" +
    	"To: " + to[0] + "\n" +
    	"Subject:" + sended.Subject + "\n" +
    	"\n" + sended.Content + "\n")
		// New album release
		// + album.Albums.Items[0].Name + " by " + album.Albums.Items[0].Artists[0].Name + " is out now !

	err = smtp.SendMail("smtp.gmail.com:587", auth, from, to, msg)
	if err != nil {
		log.Println("error = ", err)
		return errors.New("error while sending email")
	}
	return nil
}

// func SendMail

func CallCheckLOLGameStart(summoner_name string, area_id int) (LolStruct, error) {
	sqlStatement := `SELECT summoner_id FROM areas_lol WHERE id = $1`
	var summoner_id string
	err := DB_management.DB.QueryRow(sqlStatement, area_id).Scan(&summoner_id)
	if err != nil {
		log.Println("error = ", err)
		log.Println("error while getting summoner id")
		return LolStruct{}, errors.New("error while getting summoner id")
	}
	method := "GET"
	link := `https://euw1.api.riotgames.com/lol/spectator/v4/active-games/by-summoner/` + summoner_id + "?api_key=" + os.Getenv("RIOT_API_KEY")
	data := ""
	header := map[string]string{}
	resp, err := getResponseData(method, link, strings.NewReader(data), header)
	if err != nil {
		log.Println("error = ", err)
		return LolStruct{}, errors.New("error while getting game info")
	}
	sqlStatement = `SELECT in_game FROM areas_lol WHERE id = $1`
	var in_game bool
	err = DB_management.DB.QueryRow(sqlStatement, area_id).Scan(&in_game)
	if err != nil {
		log.Println("error = ", err)
		log.Println("error while getting in_game")
		return LolStruct{}, errors.New("error while getting in_game")
	}
	if resp.StatusCode == 404 {
		if in_game == true {
			sqlStatement = `UPDATE areas_lol SET in_game = $1 WHERE id = $2`
			_, err = DB_management.DB.Exec(sqlStatement, false, area_id)
			if err != nil {
				log.Println("error = ", err)
				log.Println("error while updating in_game")
				return LolStruct{}, errors.New("error while updating in_game")
			}
		}
		return LolStruct{}, errors.New("no game found")
	}
	if resp.StatusCode != 200 {
		log.Println("error = ", err)
		return LolStruct{}, errors.New("error while getting game info")
	}
	if in_game == false {
		sqlStatement = `UPDATE areas_lol SET in_game = $1 WHERE id = $2`
		_, err = DB_management.DB.Exec(sqlStatement, true, area_id)
		if err != nil {
			log.Println("error = ", err)
			log.Println("error while updating in_game")
			return LolStruct{}, errors.New("error while updating in_game")
		}
	} else {
		return LolStruct{}, errors.New("game already found")
	}
	str_result := "A game has been found for " + summoner_name + " !"
	info_res := LolStruct{SummonerName: summoner_name, EmailContent: str_result, EmailSubject: "New Game for " + summoner_name, GithubFile: "New Game"}
	return info_res, nil
}

func CallCheckLolEloChange(summoner_name string, area_id int) (LolStruct, error) {
	list_tier := map[string]int {
		"IRON": 1,
		"BRONZE": 2,
		"SILVER": 3,
		"GOLD": 4,
		"PLATINUM": 5,
		"EMERALD": 6,
		"DIAMOND": 7,
		"MASTER": 8,
		"GRANDMASTER": 9,
		"CHALLENGER": 10,
	}

	list_rank := map[string]int {
		"IV": 1,
		"III": 2,
		"II": 3,
		"I": 4,
	}

	type LolRespBodyPlayerInfo struct {
		QueueType string `json:"queueType"`
		Tier string `json:"tier"`
		Rank string `json:"rank"`
	}

	type Summoner struct {
		ID string `json:"id"`
		SummonerRank string `json:"elo"`
	}

	sqlStatement := `SELECT elo, summoner_id FROM areas_lol WHERE id = $1`
	var summoner Summoner
	err := DB_management.DB.QueryRow(sqlStatement, area_id).Scan(&summoner.SummonerRank, &summoner.ID)
	if err != nil {
		log.Println("error = ", err)
		log.Println("error while getting summoner rank")
		return LolStruct{}, errors.New("error while getting summoner rank")
	}
	method := "GET"
	link := `https://euw1.api.riotgames.com/lol/league/v4/entries/by-summoner/` + summoner.ID + "?api_key=" + os.Getenv("RIOT_API_KEY")
	data := ""
	header := map[string]string{}
	bodyText, err := callAPI(method, link, strings.NewReader(data), header)
	if err != nil {
		log.Println("error = ", err)
		return LolStruct{}, errors.New("error while getting summoner info")
	}
	var summoner_info []LolRespBodyPlayerInfo
	json.Unmarshal(bodyText, &summoner_info)
	var summoner_rank string
	for _, info := range summoner_info {
		if info.QueueType == "RANKED_SOLO_5x5" {
			summoner_rank = info.Tier + " " + info.Rank
		}
	}
	if summoner_rank == summoner.SummonerRank {
		return LolStruct{}, errors.New("no rank change")
	}
	sqlStatement = `UPDATE areas_lol SET elo = $1 WHERE id = $2`
	_, err = DB_management.DB.Exec(sqlStatement, summoner_rank, area_id)
	if err != nil {
		log.Println("error = ", err)
		log.Println("error while updating summoner rank")
		return LolStruct{}, errors.New("error while updating summoner rank")
	}
	info_res := LolStruct{SummonerName: summoner_name, EmailSubject: "New Rank for " + summoner_name, GithubFile: "New Rank"}
	if list_tier[strings.Split(summoner_rank, " ")[0]] > list_tier[strings.Split(summoner.SummonerRank, " ")[0]] {
		str_result := summoner_name + " has been promoted to " + summoner_rank + " !"
		info_res.EmailContent = str_result
		return info_res, nil
	} else if list_rank[strings.Split(summoner_rank, " ")[1]] > list_rank[strings.Split(summoner.SummonerRank, " ")[1]] {
		str_result := summoner_name + " has been promoted to " + summoner_rank + " !"
		info_res.EmailContent = str_result
		return info_res, nil
	} else if list_tier[strings.Split(summoner_rank, " ")[0]] < list_tier[strings.Split(summoner.SummonerRank, " ")[0]] {
		str_result := summoner_name + " has been demoted to " + summoner_rank + " !"
		info_res.EmailContent = str_result
		return info_res, nil
	} else {
		str_result := summoner_name + " has been demoted to " + summoner_rank + " !"
		info_res.EmailContent = str_result
		return info_res, nil
	}
	return LolStruct{}, errors.New("no rank change")
}

func CallCheckMeteo(city string) (MeteoStruct, error) {
	type Location struct {
		Latitude string `json:"lat"`
		Longitude string `json:"lon"`
	}

	type HourlyRespBody struct {
		Temperature []float64 `json:"temperature_2m"`
		Precipitation []int `json:"precipitation_probability"`
	}

	type MeteoRespBody struct {
		Hourly HourlyRespBody `json:"hourly"`
	}

	var location []Location
	cityEncoded := url.QueryEscape(city)

	method := "GET"
	link := `https://nominatim.openstreetmap.org/search?q=` + cityEncoded + `&format=json`
	data := ""
	header := map[string]string{}
	bodyText, err := callAPI(method, link, strings.NewReader(data), header)
	if err != nil {
		log.Println("error = ", err)
		return MeteoStruct{}, errors.New("error while getting location")
	}
	json.Unmarshal(bodyText, &location)

	method = "GET"
	link = `https://api.open-meteo.com/v1/forecast?latitude=` + location[0].Latitude + `&longitude=` + location[0].Longitude + `&hourly=temperature_2m,precipitation_probability&timezone=Europe%2FLondon&forecast_days=1`
	data = ""
	header = map[string]string{}
	bodyText, err = callAPI(method, link, strings.NewReader(data), header)
	if err != nil {
		log.Println("error = ", err)
		return MeteoStruct{}, errors.New("error while getting meteo info")
	}
	var meteo_info MeteoRespBody
	json.Unmarshal(bodyText, &meteo_info)
	var str_result string
	for i := 0; i < 24; i++ {
		str_result += "Hour " + strconv.Itoa(i) + " : " + strconv.FormatFloat(meteo_info.Hourly.Temperature[i], 'f', 2, 64) + "°C, " + strconv.Itoa(meteo_info.Hourly.Precipitation[i]) + "% of rain\n"
	}
	info_res := MeteoStruct{EmailContent: str_result, EmailSubject: "Meteo for " + city, GithubFile: "Meteo for " + city}
	return info_res, nil
}

func CreateFileGithub(information GithubStruct) error {
	type GithubRespBody struct {
		Sha string `json:"sha"`
	}

	if information.Owner == "" || information.RepoName == "" || information.FileName == "" || information.Content == "" || information.GithubToken == "" || information.Email == "" {
		return errors.New("missing information")
	}
	base64Content := base64.StdEncoding.EncodeToString([]byte(information.Content))
	var path_owner string
	for _, c := range information.Owner {
		if c == ' ' {
			path_owner += "-"
		} else {
			path_owner += string(c)
		}
	}
	var file_name string
	for _, c := range information.FileName {
		if c == ' ' {
			file_name += "-"
		} else {
			file_name += string(c)
		}
	}
	method := "GET"
	link := `https://api.github.com/repos/` + path_owner + `/` + information.RepoName + `/contents/` + file_name + `.txt`
	data := ""
	header := map[string]string{"Authorization": "Bearer " + information.GithubToken, "Accept": "application/vnd.github+json", "X-GitHub-Api-Version": "2022-11-28"}
	result, err := callAPI(method, link, strings.NewReader(data), header)
	var github_info GithubRespBody
	json.Unmarshal(result, &github_info)
	method = "PUT"
	link = `https://api.github.com/repos/` + path_owner + `/` + information.RepoName + `/contents/` + file_name + `.txt`
	data = `{"message" : "create file", "content" : "` + base64Content + `", "sha" : "` + github_info.Sha + `", "branch" : "main"}`
	header = map[string]string{"Authorization": "Bearer " + information.GithubToken, "Accept": "application/vnd.github+json", "X-GitHub-Api-Version": "2022-11-28"}
	_, err = getResponseData(method, link, strings.NewReader(data), header)
	if err != nil {
		log.Println("error = ", err)
		return errors.New("error while creating file")
	}
	return nil
}

func CallCheckLolLevelChange(summoner_name string, area_id int) (LolStruct, error) {
	type Summoner struct {
		ID string `json:"id"`
		SummonerLevel int `json:"summonerLevel"`
	}

	sqlStatement := `SELECT summoner_level, summoner_id FROM areas_lol WHERE id = $1`
	var summoner Summoner
	err := DB_management.DB.QueryRow(sqlStatement, area_id).Scan(&summoner.SummonerLevel, &summoner.ID)
	if err != nil {
		log.Println("error = ", err)
		log.Println("error while getting summoner level")
		return LolStruct{}, errors.New("error while getting summoner level")
	}
	method := "GET"
	link := `https://euw1.api.riotgames.com/lol/summoner/v4/summoners/` + summoner.ID + "?api_key=" + os.Getenv("RIOT_API_KEY")
	data := ""
	header := map[string]string{}
	resp, err := callAPI(method, link, strings.NewReader(data), header)
	if err != nil {
		log.Println("error = ", err)
		return LolStruct{}, errors.New("error while getting summoner info")
	}
	var summoner_info Summoner
	json.Unmarshal(resp, &summoner_info)
	if summoner_info.SummonerLevel == summoner.SummonerLevel {
		return LolStruct{}, errors.New("no level change")
	}
	sqlStatement = `UPDATE areas_lol SET summoner_level = $1 WHERE id = $2`
	_, err = DB_management.DB.Exec(sqlStatement, summoner_info.SummonerLevel, area_id)
	if err != nil {
		log.Println("error = ", err)
		log.Println("error while updating summoner level")
		return LolStruct{}, errors.New("error while updating summoner level")
	}
	info_res := LolStruct{SummonerName: summoner_name, EmailSubject: "New Level for " + summoner_name, GithubFile: "New Level"}
	str_result := summoner_name + " is now level " + strconv.Itoa(summoner_info.SummonerLevel) + " !"
	info_res.EmailContent = str_result
	return info_res, nil
}

func CallCheckTftEloChange(summoner_name string, area_id int) (TftStruct, error) {
	list_tier := map[string]int {
		"IRON": 1,
		"BRONZE": 2,
		"SILVER": 3,
		"GOLD": 4,
		"PLATINUM": 5,
		"EMERALD": 6,
		"DIAMOND": 7,
		"MASTER": 8,
		"GRANDMASTER": 9,
		"CHALLENGER": 10,
	}

	list_rank := map[string]int {
		"IV": 1,
		"III": 2,
		"II": 3,
		"I": 4,
	}

	type TftRespBodyPlayerInfo struct {
		QueueType string `json:"queueType"`
		Tier string `json:"tier"`
		Rank string `json:"rank"`
	}

	type Summoner struct {
		ID string `json:"id"`
		SummonerRank string `json:"elo"`
	}

	sqlStatement := `SELECT elo, summoner_id FROM areas_tft WHERE id = $1`
	var summoner Summoner
	err := DB_management.DB.QueryRow(sqlStatement, area_id).Scan(&summoner.SummonerRank, &summoner.ID)
	if err != nil {
		log.Println("error = ", err)
		log.Println("error while getting summoner rank")
		return TftStruct{}, errors.New("error while getting summoner rank")
	}
	method := "GET"
	link := `https://euw1.api.riotgames.com/tft/league/v1/entries/by-summoner/` + summoner.ID + "?api_key=" + os.Getenv("RIOT_API_KEY")
	data := ""
	header := map[string]string{}
	bodyText, err := callAPI(method, link, strings.NewReader(data), header)
	if err != nil {
		log.Println("error = ", err)
		return TftStruct{}, errors.New("error while getting summoner info")
	}
	var summoner_info []TftRespBodyPlayerInfo
	json.Unmarshal(bodyText, &summoner_info)
	var summoner_rank string
	for _, info := range summoner_info {
		if info.QueueType == "RANKED_TFT" {
			summoner_rank = info.Tier + " " + info.Rank
		}
	}
	if summoner_rank == summoner.SummonerRank {
		return TftStruct{}, errors.New("no rank change")
	}
	sqlStatement = `UPDATE areas_tft SET elo = $1 WHERE id = $2`
	_, err = DB_management.DB.Exec(sqlStatement, summoner_rank, area_id)
	if err != nil {
		log.Println("error = ", err)
		log.Println("error while updating summoner rank")
		return TftStruct{}, errors.New("error while updating summoner rank")
	}
	info_res := TftStruct{SummonerName: summoner_name, EmailSubject: "New Rank for " + summoner_name, GithubFile: "New Rank"}
	if list_tier[strings.Split(summoner_rank, " ")[0]] > list_tier[strings.Split(summoner.SummonerRank, " ")[0]] {
		str_result := summoner_name + " has been promoted to " + summoner_rank + " !"
		info_res.EmailContent = str_result
		return info_res, nil
	} else if list_rank[strings.Split(summoner_rank, " ")[1]] > list_rank[strings.Split(summoner.SummonerRank, " ")[1]] {
		str_result := summoner_name + " has been promoted to " + summoner_rank + " !"
		info_res.EmailContent = str_result
		return info_res, nil
	} else if list_tier[strings.Split(summoner_rank, " ")[0]] < list_tier[strings.Split(summoner.SummonerRank, " ")[0]] {
		str_result := summoner_name + " has been demoted to " + summoner_rank + " !"
		info_res.EmailContent = str_result
		return info_res, nil
	} else {
		str_result := summoner_name + " has been demoted to " + summoner_rank + " !"
		info_res.EmailContent = str_result
		return info_res, nil
	}
	return TftStruct{}, errors.New("no rank change")
}

func CallCheckTftWinChange(summoner_name string, area_id int) (TftStruct, error) {
	type Summoner struct {
		Win int `json:"wins"`
	}
	sqlStatement := `SELECT win, summoner_id FROM areas_tft WHERE id = $1`
	var win int
	var summoner_id string
	err := DB_management.DB.QueryRow(sqlStatement, area_id).Scan(&win, &summoner_id)
	if err != nil {
		log.Println("error = ", err)
		log.Println("error while getting summoner win")
		return TftStruct{}, errors.New("error while getting summoner win")
	}
	method := "GET"
	link := `https://euw1.api.riotgames.com/tft/league/v1/entries/by-summoner/` + summoner_id + "?api_key=" + os.Getenv("RIOT_API_KEY")
	data := ""
	header := map[string]string{}
	resp, err := callAPI(method, link, strings.NewReader(data), header)
	if err != nil {
		log.Println("error = ", err)
		return TftStruct{}, errors.New("error while getting summoner info")
	}
	var summoner_info Summoner
	json.Unmarshal(resp, &summoner_info)
	if summoner_info.Win == win {
		return TftStruct{}, errors.New("no win change")
	}
	sqlStatement = `UPDATE areas_tft SET win = $1 WHERE id = $2`
	_, err = DB_management.DB.Exec(sqlStatement, summoner_info.Win, area_id)
	if err != nil {
		log.Println("error = ", err)
		log.Println("error while updating summoner win")
		return TftStruct{}, errors.New("error while updating summoner win")
	}
	info_res := TftStruct{SummonerName: summoner_name, EmailSubject: "New Win for " + summoner_name, GithubFile: "New Win"}
	str_result := summoner_name + " has now " + strconv.Itoa(summoner_info.Win) + " win !"
	info_res.EmailContent = str_result
	return info_res, nil
}

func CallCheckTftLoseChange(summoner_name string, area_id int) (TftStruct, error) {
	type Summoner struct {
		Lose int `json:"losses"`
	}
	sqlStatement := `SELECT lose, summoner_id FROM areas_tft WHERE id = $1`
	var lose int
	var summoner_id string
	err := DB_management.DB.QueryRow(sqlStatement, area_id).Scan(&lose, &summoner_id)
	if err != nil {
		log.Println("error = ", err)
		log.Println("error while getting summoner lose")
		return TftStruct{}, errors.New("error while getting summoner lose")
	}
	method := "GET"
	link := `https://euw1.api.riotgames.com/tft/league/v1/entries/by-summoner/` + summoner_id + "?api_key=" + os.Getenv("RIOT_API_KEY")
	data := ""
	header := map[string]string{}
	resp, err := callAPI(method, link, strings.NewReader(data), header)
	if err != nil {
		log.Println("error = ", err)
		return TftStruct{}, errors.New("error while getting summoner info")
	}
	var summoner_info Summoner
	json.Unmarshal(resp, &summoner_info)
	if summoner_info.Lose == lose {
		return TftStruct{}, errors.New("no lose change")
	}
	sqlStatement = `UPDATE areas_tft SET lose = $1 WHERE id = $2`
	_, err = DB_management.DB.Exec(sqlStatement, summoner_info.Lose, area_id)
	if err != nil {
		log.Println("error = ", err)
		log.Println("error while updating summoner lose")
		return TftStruct{}, errors.New("error while updating summoner lose")
	}
	info_res := TftStruct{SummonerName: summoner_name, EmailSubject: "New lose for " + summoner_name, GithubFile: "New lose"}
	str_result := summoner_name + " has now " + strconv.Itoa(summoner_info.Lose) + " lose !"
	info_res.EmailContent = str_result
	return info_res, nil
}

func CallCheckGithubPush(owner string, repo_name string, token string, area_id int) (GithubStruct, error) {
	sqlStatement := `SELECT nbr_push FROM areas_github WHERE id = $1`
	var nbr_push int
	err := DB_management.DB.QueryRow(sqlStatement, area_id).Scan(&nbr_push)
	if err != nil {
		log.Println("error = ", err)
		nbr_push = 0
	}
	method := "GET"
	link := `https://api.github.com/repos/` + owner + `/` + repo_name + `/activity`
	data := ""
	header := map[string]string{"Authorization": "Bearer " + token, "Accept": "application/vnd.github+json", "X-GitHub-Api-Version": "2022-11-28"}
	resp, err := callAPI(method, link, strings.NewReader(data), header)
	if err != nil {
		log.Println("error = ", err)
		return GithubStruct{}, errors.New("error while getting github info")
	}
	var github_info []map[string]interface{}
	json.Unmarshal(resp, &github_info)
	if len(github_info) == nbr_push {
		return GithubStruct{}, errors.New("no push")
	}
	sqlStatement = `UPDATE areas_github SET nbr_push = $1 WHERE id = $2`
	_, err = DB_management.DB.Exec(sqlStatement, len(github_info), area_id)
	if err != nil {
		log.Println("error = ", err)
		log.Println("error while updating nbr push")
		return GithubStruct{}, errors.New("error while updating nbr push")
	}
	info_res := GithubStruct{Owner: owner, RepoName: repo_name, GithubToken: token, EmailSubject: "New push", GithubFile: "New push"}
	str_result := "New push in repo" + repo_name + "\n"
	info_res.Content = str_result
	return info_res, nil
}

func CallCheckCSUpdate(area_id int) (SteamStruct, error) {
	type SteamRespBody struct {
		Result struct {
			App struct {
				Version int `json:"version"`
			} `json:"app"`
		} `json:"result"`
	}

	method := "GET"
	link := `https://api.steampowered.com/ICSGOServers_730/GetGameServersStatus/v1/?key=` + os.Getenv("STEAM_API_KEY")
	data := ""
	header := map[string]string{}
	resp, err := callAPI(method, link, strings.NewReader(data), header)
	if err != nil {
		log.Println("error = ", err)
		return SteamStruct{}, errors.New("error while getting steam info")
	}
	var steam_info SteamRespBody
	json.Unmarshal(resp, &steam_info)
	if steam_info.Result.App.Version == 0 {
		return SteamStruct{}, errors.New("no update")
	}
	sqlStatement := `SELECT version FROM areas_steam WHERE id = $1`
	var last_version (*int)
	err = DB_management.DB.QueryRow(sqlStatement, area_id).Scan(&last_version)
	if err != nil {
		log.Println("error = ", err)
		log.Println("error while getting last version")
		return SteamStruct{}, errors.New("error while getting last version")
	}
	if last_version == nil {
		sqlStatement = `UPDATE areas_steam SET version = $1 WHERE id = $2`
		_, err = DB_management.DB.Exec(sqlStatement, steam_info.Result.App.Version, area_id)
		if err != nil {
			log.Println("error = ", err)
			log.Println("error while updating version")
			return SteamStruct{}, errors.New("error while updating version")
		}
		return SteamStruct{}, errors.New("no update")
	}
	if steam_info.Result.App.Version == *last_version {
		return SteamStruct{}, errors.New("no update")
	}
	sqlStatement = `UPDATE areas_steam SET version = $1 WHERE id = $2`
	_, err = DB_management.DB.Exec(sqlStatement, steam_info.Result.App.Version, area_id)
	if err != nil {
		log.Println("error = ", err)
		log.Println("error while updating version")
		return SteamStruct{}, errors.New("error while updating version")
	}
	info_res := SteamStruct{EmailSubject: "New CS:2 update", GithubFile: "New CS:2 update", EmailContent: "New CS:GO update\nhttps://blog.counter-strike.net/index.php/category/updates/\n"}
	return info_res, nil
}

func CallCheckTFUpdate(area_id int) (SteamStruct, error) {
	type SteamRespBody struct {
		Result struct {
			Version int `json:"active_version"`
		} `json:"result"`
	}

	method := "GET"
	link := `https://api.steampowered.com/IGCVersion_440/GetServerVersion/v1/?key=` + os.Getenv("STEAM_API_KEY")
	data := ""
	header := map[string]string{}
	resp, err := callAPI(method, link, strings.NewReader(data), header)
	if err != nil {
		log.Println("error = ", err)
		return SteamStruct{}, errors.New("error while getting steam info")
	}
	var steam_info SteamRespBody
	json.Unmarshal(resp, &steam_info)
	if steam_info.Result.Version == 0 {
		return SteamStruct{}, errors.New("no update")
	}
	sqlStatement := `SELECT version FROM areas_steam WHERE id = $1`
	var last_version (*int)
	err = DB_management.DB.QueryRow(sqlStatement, area_id).Scan(&last_version)
	if err != nil {
		log.Println("error = ", err)
		log.Println("error while getting last version")
		return SteamStruct{}, errors.New("error while getting last version")
	}
	if last_version == nil {
		sqlStatement = `UPDATE areas_steam SET version = $1 WHERE id = $2`
		_, err = DB_management.DB.Exec(sqlStatement, steam_info.Result.Version, area_id)
		if err != nil {
			log.Println("error = ", err)
			log.Println("error while updating version")
			return SteamStruct{}, errors.New("error while updating version")
		}
		return SteamStruct{}, errors.New("no update")
	}
	if steam_info.Result.Version == *last_version {
		return SteamStruct{}, errors.New("no update")
	}
	sqlStatement = `UPDATE areas_steam SET version = $1 WHERE id = $2`
	_, err = DB_management.DB.Exec(sqlStatement, steam_info.Result.Version, area_id)
	if err != nil {
		log.Println("error = ", err)
		log.Println("error while updating version")
		return SteamStruct{}, errors.New("error while updating version")
	}
	info_res := SteamStruct{EmailSubject: "New TF2 update", GithubFile: "New TF2 update", EmailContent: "An update has been released to Team Fortress 2\nhttps://www.teamfortress.com/?tab=updates\n"}
	return info_res, nil
}

func CallCheckDotaUpdate(area_id int) (SteamStruct, error) {
	type SteamRespBody struct {
		Result struct {
			Version int `json:"active_version"`
		} `json:"result"`
	}

	method := "GET"
	link := `https://api.steampowered.com/IGCVersion_570/GetServerVersion/v1/?key=` + os.Getenv("STEAM_API_KEY")
	data := ""
	header := map[string]string{}
	resp, err := callAPI(method, link, strings.NewReader(data), header)
	if err != nil {
		log.Println("error = ", err)
		return SteamStruct{}, errors.New("error while getting steam info")
	}
	var steam_info SteamRespBody
	json.Unmarshal(resp, &steam_info)
	if steam_info.Result.Version == 0 {
		return SteamStruct{}, errors.New("no update")
	}
	sqlStatement := `SELECT version FROM areas_steam WHERE id = $1`
	var last_version (*int)
	err = DB_management.DB.QueryRow(sqlStatement, area_id).Scan(&last_version)
	if err != nil {
		log.Println("error = ", err)
		log.Println("error while getting last version")
		return SteamStruct{}, errors.New("error while getting last version")
	}
	if last_version == nil {
		sqlStatement = `UPDATE areas_steam SET version = $1 WHERE id = $2`
		_, err = DB_management.DB.Exec(sqlStatement, steam_info.Result.Version, area_id)
		if err != nil {
			log.Println("error = ", err)
			log.Println("error while updating version")
			return SteamStruct{}, errors.New("error while updating version")
		}
		return SteamStruct{}, errors.New("no update")
	}
	if steam_info.Result.Version == *last_version {
		return SteamStruct{}, errors.New("no update")
	}
	sqlStatement = `UPDATE areas_steam SET version = $1 WHERE id = $2`
	_, err = DB_management.DB.Exec(sqlStatement, steam_info.Result.Version, area_id)
	if err != nil {
		log.Println("error = ", err)
		log.Println("error while updating version")
		return SteamStruct{}, errors.New("error while updating version")
	}
	info_res := SteamStruct{EmailSubject: "New Dota 2 update", GithubFile: "New Dota 2 update", EmailContent: "An update has been released to Dota 2\nhttps://www.dota2.com/news/updates/\n"}
	return info_res, nil
}

func CallCheckCSMonthlyPlayerCount(area_id int) (SteamStruct, error) {
	type SteamRespBody struct {
		Result struct {
			PlayerCount int `json:"players"`
		} `json:"result"`
	}

	method := "GET"
	link := `https://api.steampowered.com/ISteamUserStats/GetNumberOfCurrentPlayers/v1/?appid=730`
	data := ""
	header := map[string]string{}
	resp, err := callAPI(method, link, strings.NewReader(data), header)
	if err != nil {
		log.Println("error = ", err)
		return SteamStruct{}, errors.New("error while getting steam info")
	}
	var steam_info SteamRespBody
	json.Unmarshal(resp, &steam_info)
	if steam_info.Result.PlayerCount == 0 {
		return SteamStruct{}, errors.New("no player count")
	}
	sqlStatement := `SELECT player_count FROM areas_steam WHERE id = $1`
	var last_player_count (*int)
	err = DB_management.DB.QueryRow(sqlStatement, area_id).Scan(&last_player_count)
	if err != nil {
		log.Println("error = ", err)
		log.Println("error while getting last player count")
		return SteamStruct{}, errors.New("error while getting last player count")
	}
	if last_player_count == nil {
		sqlStatement = `UPDATE areas_steam SET player_count = $1 WHERE id = $2`
		_, err = DB_management.DB.Exec(sqlStatement, steam_info.Result.PlayerCount, area_id)
		if err != nil {
			log.Println("error = ", err)
			log.Println("error while updating player count")
			return SteamStruct{}, errors.New("error while updating player count")
		}
		return SteamStruct{}, errors.New("no player count")
	}
	if steam_info.Result.PlayerCount == *last_player_count {
		return SteamStruct{}, errors.New("no player count")
	}
	sqlStatement = `UPDATE areas_steam SET player_count = $1 WHERE id = $2`
	_, err = DB_management.DB.Exec(sqlStatement, steam_info.Result.PlayerCount, area_id)
	if err != nil {
		log.Println("error = ", err)
		log.Println("error while updating player count")
		return SteamStruct{}, errors.New("error while updating player count")
	}
	info_res := SteamStruct{EmailSubject: "This month CS:2 player count", GithubFile: "This month CS:2 player count", EmailContent: "CS:2 had " + strconv.Itoa(steam_info.Result.PlayerCount) + " players this month.\n"}
	return info_res, nil
}

func CreateRepoGithub(information GithubStruct) error {
	if information.Owner == "" || information.GithubToken == "" || information.FileName == "" {
		return errors.New("missing information")
	}
	var repo_name string
	log.Println("information.GithubFile = ", information.FileName)
	for _, c := range information.FileName {
		if c == ' ' || c == '.' || c == ':' {
			repo_name += "-"
		} else {
			repo_name += string(c)
		}
	}
	method := "POST"
	link := `https://api.github.com/user/repos`
	log.Println("link = ", link)
	data := `{"name" : "` + repo_name + `", "private": true}`
	log.Println("data = ", data)
	header := map[string]string{"Authorization": "Bearer " + information.GithubToken, "Accept": "application/vnd.github+json", "X-GitHub-Api-Version": "2022-11-28"}
	log.Println("header = ", header)
	resulte, err := getResponseData(method, link, strings.NewReader(data), header)
	log.Println("resulte = ", resulte)
	if err != nil {
		log.Println("error = ", err)
		return errors.New("error while creating repo")
	}
	return nil
}

func CallDiscordBot(user_id string, content string) error {
	type Message struct {
		Command string `json:"command"`
		UserId string `json:"user_id"`
		Content string `json:"content"`
	}
	conn, err := net.Dial("tcp", "localhost:55555")
    if err != nil {
        log.Println("error = ", err)
		return errors.New("error while connecting to discord bot")
    }
    defer conn.Close()

    msg := Message{
		Command: "private_message",
		UserId: user_id,
        Content: content,
    }

    data, err := json.Marshal(msg)
    if err != nil {
        log.Println("error = ", err)
		return errors.New("error while marshalling")
    }

    conn.Write(data)
	return nil
}