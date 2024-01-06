package Email

import (
	"errors"
	"API/call_API"
	// "log"
)

type reaction struct {
	name string `json:"name"`
	reaction func(call_API.EmailStruct) error `json:"reaction"`
}

var listReaction = []reaction {{name: "send_mail", reaction: call_API.SendMail}}

func CheckReactionExecution(reaction call_API.CallReactionStruct) error {
	for _, a := range listReaction {
		if a.name == reaction.AreaInfo.Reaction_name {
			// log.Println("reaction email = ", reaction.Email)
			return a.reaction(reaction.Email)
		}
	}
	return errors.New("unknown reaction")
}