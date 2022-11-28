package models

type CollaboratorPermissions struct {
	Id             int  `json:"id"`
	IdCollaborator int  `json:"idCollaborator"`
	IdPermissions  int  `json:"idPermissions"`
	Consult        bool `json:"consult"`
	Register       bool `json:"register"`
	Update         bool `json:"update"`
	Delete         bool `json:"delete"`
	PrintOut       bool `json:"printOut"`
}
