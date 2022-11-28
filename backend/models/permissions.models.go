package models

type Permissions struct {
	Id             int    `json:"id"`
	Name           string `json:"name"`
	Active         bool   `json:"active"`
	IsPartConsult  bool   `json:"isPartConsult"`
	IsPartRegister bool   `json:"isPartRegister"`
	IsPartUpdate   bool   `json:"isPartUpdate"`
	IsPartDelete   bool   `json:"isPartDelete"`
	IsPartPrintOut bool   `json:"isPartPrintOut"`
}
