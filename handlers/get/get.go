package get

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/go-sql-driver/mysql"
	mc "github.com/shysudo/meetup/common"
	mm "github.com/shysudo/meetup/model"
	"net/http"
	"sync"
)

var getwg sync.WaitGroup

func getRoutine(c chan []mm.Participant ) {
	defer getwg.Done()
	rows, err := mc.DB.Query(`select name, age, dob, profession, locality, number_of_guest, address from participants`)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	var participant mm.Participant
	var parti []mm.Participant
	var name sql.NullString
	var age sql.NullInt64
	var dateofbirthday mysql.NullTime
	var locality sql.NullString
	var numberofguest sql.NullInt64
	var address sql.NullString
	var profession sql.NullString
	for rows.Next() {
		if err = rows.Scan(&name, &age, &dateofbirthday, &profession, &locality, &numberofguest, &address); err != nil {
			fmt.Println(err.Error())
			return
		}
		if name.Valid {
			participant.Name = name.String
		}
		if age.Valid {
			participant.Age = age.Int64
		}
		if dateofbirthday.Valid {
			participant.DOB = dateofbirthday.Time
		}
		if locality.Valid {
			participant.Locality = locality.String
		}
		if numberofguest.Valid {
			participant.NumberOfGuest = numberofguest.Int64
		}
		if address.Valid {
			participant.Address = address.String
		}
		if profession.Valid {
			participant.Profession =  profession.String
		}
		parti = append(parti, participant)
	}
	c <- parti
	return
}

func GetParticipantListHandler(w http.ResponseWriter, r *http.Request) {
	//get participant
	var participants []mm.Participant
	getwg.Add(1)
	c := make(chan []mm.Participant, 30)
	go getRoutine(c)
	participants = <-c
	getwg.Wait()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(participants)
}
