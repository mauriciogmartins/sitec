package enuns

type ContactsType int

const (
	PhoneMain ContactsType = iota + 1
	PhoneSecondary
	EmailMain
	EmailSecondary
	WhatsappMain
	WhatsappSecondary
	Website
	Facebook
	Instagram
)

func (c ContactsType) ContactsTypeToString() string {
	return [...]string{"PhoneMain", "PhoneSecondary", "EmailMain", "EmailSecondary", "WhatsappMain", "WhatsappSecondary", "Website", "Facebook", "Instagram"}[c-1]
}
