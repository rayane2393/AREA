package Discord

import (
	"errors"
	"API/call_API"
	"API/DB_management"
	"log"
)

type reaction struct {
	name string `json:"name"`
	reaction func(string, string) error `json:"reaction"`
}

var listReaction = []reaction {{name: "send_message", reaction: call_API.CallDiscordBot}}

func CheckReactionExecution(reaction call_API.CallReactionStruct) error {
	for _, a := range listReaction {
		if a.name == reaction.AreaInfo.Reaction_name {
			sqlStatement := `SELECT discord_id FROM users WHERE id = $1`
			var discord_id string
			err := DB_management.DB.QueryRow(sqlStatement, reaction.AreaInfo.UserID).Scan(&discord_id)
			if err != nil {
				log.Println("error while getting discord id")
				return errors.New("error while getting discord id")
			}
			return a.reaction(discord_id, reaction.Email.Content)
		}
	}
	return errors.New("unknown reaction")
}