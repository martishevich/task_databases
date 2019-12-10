package model

type ContactsRepository interface {
	Save(Contact) (Contact, error)
	ListAll() ([]Contact, error)
	GetByID(uint) (Contact, error)
	GetByPhone(string) (Contact, error)
	GetByEmail(string) (Contact, error)
	SearchByName(string) ([]Contact, error)
	Delete(uint) error
}

type Contact struct {
	ID uint `json:"id"`

	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`

	Phone string `json:"phone"`
	Email string `json:"email"`
}
