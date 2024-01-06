package DB_management

import (
	"fmt"
	"log"
	"os"
	"database/sql"
	"github.com/joho/godotenv"
	"errors"
	_ "github.com/lib/pq"
)

type User struct {
	Username string `json:"username"`
	Token (*string) `json:"token, omitempt"`
	Password (*string) `json:"password, omitempt"`
	Email (*string) `json:"email, omitempt"`
	Phone (*string) `json:"phone, omitempt"`
	Github_id (*int) `json:"github_id, omitempt"`
	Github_token (*string) `json:"github_token, omitempt"`
	Google_id (*int) `json:"google_id, omitempt"`
	Google_token (*string) `json:"google_token, omitempt"`
	Github_email (*string) `json:"github_email, omitempt"`
	Github_pseudo (*string) `json:"github_pseudo, omitempt"`
	Discord_id (*string) `json:"discord_id, omitempt"`
	Discord_token (*string) `json:"discord_token, omitempt"`
}

type Area struct {
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
	City string `json:"city"`
	Artist_ID string `json:"artist_id"`
	Title string `json:"title"`
	Content string `json:"content"`
	Interval string `json:"interval"`
	Repo_name string `json:"repo_name"`
}

var Users = []User {
}

var DB = OpenDB()

func ConnectToDb() {
	// sqlStatement := `INSERT INTO users (username, email, hashed_password)
	// VALUES('Rayane', 'rayane.bullshit@gmail.com', 'E')`
	// _, err = db.Exec(sqlStatement)
	info, err := DB.Query("SELECT username, email, password, token, phone, github_id, github_token, google_id, google_token, github_email, github_pseudo, discord_id, discord_token FROM users")
	if err != nil {
		log.Println("error while getting DB infos")
		panic(err)
	}
	b := User{}
	for info.Next() {
		err = info.Scan(&b.Username, &b.Email, &b.Password, &b.Token, &b.Phone, &b.Github_id, &b.Github_token, &b.Google_id, &b.Google_token, &b.Github_email, &b.Github_pseudo, &b.Discord_id, &b.Discord_token)
		if err != nil {
			log.Println("error while getting user")
			log.Println("error = ", err)
			panic(err)
		}
		// fmt.Println(b)
		Users = append(Users, b)
	}
}

func OpenDB() *sql.DB {
	err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
    "password=%s dbname=%s sslmode=disable",
	os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), "back-end", os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
  	db, err := sql.Open("postgres", psqlInfo)
  	if err != nil {
		log.Println("error while connecting to DB")
  	  	panic(err)
  	}
	fmt.Println("DB successfully connected!")
	return db
}

// func CloseDB(db *sql.DB) {
// 	defer db.Close()

//   	err := db.Ping()
//   	if err != nil {
// 	  	log.Println("error while closing DB")
//   	  	panic(err)
// 	}
// 	fmt.Println("DB successfully closed!")
// }

func GetStringFromDB(sqlStatement string, args ...interface{}) (string, error) {
	var result string

	err := DB.QueryRow(sqlStatement, args...).Scan(&result)
	if err != nil {
		log.Println("error while getting string from DB")
		return "", err
	}
	return result, nil
}

func GetIntFromDB(sqlStatement string, args ...interface{}) (int, error) {
	var result int

	err := DB.QueryRow(sqlStatement, args...).Scan(&result)
	if err != nil {
		log.Println("error = ", err)
		log.Println("error while getting int from DB")
		return 0, err
	}
	return result, nil
}

func GetAreaWithId(id int, db_name string) (Area, error) {
	var area Area

	sqlStatement := `SELECT id, user_id, name, color, reaction, toggled, time_execute, action_name, reaction_name FROM areas_` + db_name + ` WHERE id = $1`
	err := DB.QueryRow(sqlStatement, id).Scan(&area.ID, &area.UserID, &area.Name, &area.Color, &area.Reaction, &area.Toggled, &area.Time_execute, &area.Action_name, &area.Reaction_name)
	if err != nil {
		log.Println("error = ", err)
		log.Println("error while getting area with id")
		return area, err
	}
	return area, nil
}

func VerifyValidServices(area Area) error {
	var services string

	sqlStatement := `SELECT services FROM action WHERE name = $1`
	err := DB.QueryRow(sqlStatement, area.Action_name).Scan(&services)
	if err != nil {
		log.Println("wrong action name")
		log.Println("error = ", err)
		return errors.New("wrong action name")
	}
	sqlStatement = `SELECT services FROM reactions WHERE name = $1`
	err = DB.QueryRow(sqlStatement, area.Reaction_name).Scan(&services)
	if err != nil {
		log.Println("error while getting services")
		return errors.New("error while getting services")
	}
	return nil
}

func CreateUser(sqlStatement string, args ...interface{}) (error) {
	// github_id := 0
	_, err := DB.Exec(sqlStatement, args...)
	if err != nil {
		log.Println("error while creating user")
		log.Println("error = ", err)
		return err
	}
	return nil
}