package domain

type (
	Recipient struct {
		Name         string
		PhoneNumbers []string
		Comment      string
		Tags         []string
	}

	RecipientList []*Recipient
)
