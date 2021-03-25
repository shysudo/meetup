package update

import (
	"fmt"
	mc "github.com/shysudo/meetup/common"
	"github.com/shysudo/meetup/handlers/common"
	mm "github.com/shysudo/meetup/model"
	"log"
	"net/http"
	"sync"
)

var updatewg sync.WaitGroup

func updateRoutine(c chan mm.Participant, parti mm.Participant) {
	defer updatewg.Done()
	stmt, err := mc.DB.Prepare(`update participants set name = ?, age = ?, dob = ?, profession = ?, locality = ?, number_of_guest = ?, address = ? where name = ?`)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer stmt.Close()
	_, err = stmt.Exec(parti.Name, parti.Age, parti.DOB, parti.Profession, parti.Locality, parti.NumberOfGuest, parti.Address, parti.Name)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	c <- parti
}

func UpdateParticipantHandler(w http.ResponseWriter, r *http.Request) {
	participant, err := common.Readfromrequest(r)
	if err != nil {
		log.Println(err.Error())
		return
	}
	//Add participant
	p := make(chan mm.Participant, 30)
	fmt.Println(participant)
	updatewg.Add(1)
	go updateRoutine(p, participant)
	<-p
	updatewg.Wait()
	close(p)
}
