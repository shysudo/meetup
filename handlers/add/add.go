package add

import (
	"encoding/json"
	"fmt"
	mc "github.com/shysudo/meetup/common"
	"github.com/shysudo/meetup/handlers/common"
	mm "github.com/shysudo/meetup/model"
	"log"
	"net/http"
	"sync"
)

var addwg sync.WaitGroup

func addRoutine(c chan mm.Participant, parti mm.Participant) {
	defer addwg.Done()
	stmt, err := mc.DB.Prepare(`insert into participants(name, age, dob, profession, locality, number_of_guest, address) values (?,?,?,?,?,?,?)`)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer stmt.Close()
	_, err = stmt.Exec(parti.Name, parti.Age, parti.DOB, parti.Profession, parti.Locality, parti.NumberOfGuest, parti.Address)
	if err != nil {
		fmt.Println(err.Error())
	}
	c <- parti
}

func RegisterParticipantHandler(w http.ResponseWriter, r *http.Request) {
	participant, err := common.Readfromrequest(r)
	if err != nil {
		log.Println(err.Error())
		return
	}
	//Add participant
	messages := common.ValidateParticipant(participant)
	if len(messages) > 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(messages)
		return
	}
	p := make(chan mm.Participant, 30)
	fmt.Println(participant)
	addwg.Add(1)
	go addRoutine(p, participant)
	<-p
	addwg.Wait()
	close(p)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(participant)
}
