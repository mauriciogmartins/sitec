package models

type State struct {
	Id       int    `json:"id"`
	IdIbge   int    `json:"idIbge"`
	State    string `json:"state"`
	Initials string `json:"initials"`
}
