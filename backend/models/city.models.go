package models

type City struct {
	Id      int    `json:"id"`
	IdState int    `json:"idState"`
	Active  bool   `json:"active"`
	Name    string `json:"name"`
}
