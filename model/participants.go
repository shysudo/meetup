package model

import "time"

type Participant struct {
	Name          string    `json:"name"`
	Age           int64     `json:"age"`
	DOB           time.Time `json:"dob"`
	Profession    string    `json:"profession"`
	Locality      string    `json:"locality"`
	NumberOfGuest int64     `json:"number_of_guest"`
	Address       string    `json:"address"`
}

const (
	None = iota
	One
	Two
)

const (
	Employed = "Employed"
	Student = "Student"
)