package models

import "sitec/enuns"

type Collaboratos struct {
	Id         int                     `json:"id"`
	Active     bool                    `json:"active"`
	Name       string                  `json:"name"`
	Cpf        string                  `json:"cpf"`
	Photograph string                  `json:"photograph"`
	Address    string                  `json:"address"`
	Number     int                     `json:"number"`
	District   string                  `json:"district"`
	City       int                     `json:"city"`
	Cep        string                  `json:"cep"`
	State      int                     `json:"state"`
	Contacts   []CollaboratorsContacts `json:"contacts"`
	User       string                  `json:"user"`
	Password   string                  `json:"password"`
}

type CollaboratorsContacts struct {
	Type enuns.ContactsType `json:"type"`
	Data string             `json:"data"`
}
