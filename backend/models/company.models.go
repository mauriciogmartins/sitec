package models

import "sitec/enuns"

type Company struct {
	Id            int               `json:"id"`
	CorporateName string            `json:"corporateName"`
	FantasyName   string            `json:"fantasyName"`
	Cnpj          string            `json:"cnpj"`
	Ie            string            `json:"ie"`
	Im            string            `json:"im"`
	Soon          string            `json:"soon"`
	Address       string            `json:"address"`
	Number        int               `json:"number"`
	District      string            `json:"district"`
	City          int               `json:"city"`
	Cep           string            `json:"cep"`
	State         int               `json:"state"`
	Contacts      []CompanyContacts `json:"contacts"`
}

type CompanyContacts struct {
	Type   enuns.ContactsType `json:"type"`
	Data   string             `json:"data"`
	Public bool               `json:"public"`
}
