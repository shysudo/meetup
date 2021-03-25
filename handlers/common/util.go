package common

import (
	"encoding/json"
	"fmt"
	mm "github.com/shysudo/meetup/model"
	"log"
	"net/http"
)

func Readfromrequest(r *http.Request) (mm.Participant, error) {
	var participant mm.Participant
	if err := json.NewDecoder(r.Body).Decode(&participant); err != nil {
		log.Println(err.Error())
		return participant, err
	}
	return participant, nil
}

func ValidateParticipant(participant mm.Participant) ([]string) {
		var messages = []string{}
		if participant.NumberOfGuest != mm.None || participant.NumberOfGuest != mm.One || participant.NumberOfGuest != mm.Two {
			fmt.Println("Number of guest should be less than equal to 2")
			messages = append(messages, "Number of guest should be less than equal to 2")
		}
		if participant.Profession != mm.Student || participant.Profession != mm.Employed {
			fmt.Println("Not a valid Profession")
			messages = append(messages, "Not a valid Profession")
		}
	return messages
}
