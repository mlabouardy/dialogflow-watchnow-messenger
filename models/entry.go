package models

type Entry struct {
	ID        string      `json:"id,omitempty"`
	Time      int         `json:"time,omitempty"`
	Messaging []Messaging `json:"messaging,omitempty"`
}
