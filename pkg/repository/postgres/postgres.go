package postgres

import (
	"database/sql"
	"github.com/burov/task_databases/pkg/model"
)

type ContactsRepositoryPostgreSQL struct {
	db *sql.DB
}

func NewContactsRepositoryPostgreSQL(db *sql.DB) *ContactsRepositoryPostgreSQL {
	return &ContactsRepositoryPostgreSQL{
		db: db,
	}
}

func (r *ContactsRepositoryPostgreSQL) Save(contact model.Contact) (model.Contact, error) {
	query := "INSERT INTO contacts (first_name, last_name, phone, email) VALUES($1, $2, $3, $4) returning id;"
	err := r.db.QueryRow(query, contact.FirstName, contact.LastName, contact.Phone, contact.Email).Scan(&contact.ID)
	return contact, err
}

func (r *ContactsRepositoryPostgreSQL) ListAll() (contacts []model.Contact, err error) {
	query := "SELECT id, first_name, last_name, email, phone FROM contacts;"
	rows, err := r.db.Query(query)

	if err != nil {
		return contacts, err
	}

	defer rows.Close()
	for rows.Next() {
		var contact model.Contact
		err = rows.Scan(&contact.ID, &contact.FirstName, &contact.LastName, &contact.Email, &contact.Phone)
		if err != nil {
			return nil, err
		}
		contacts = append(contacts, contact)
	}

	return contacts, nil
}

func (r *ContactsRepositoryPostgreSQL) GetByID(id uint) (contact model.Contact, err error) {
	query := "SELECT id, first_name, last_name, email, phone FROM contacts WHERE id = $1;"
	row := r.db.QueryRow(query, id)
	err = row.Scan(&contact.ID, &contact.FirstName, &contact.LastName, &contact.Email, &contact.Phone)
	return contact, err
}

func (r *ContactsRepositoryPostgreSQL) GetByPhone(phone string) (contact model.Contact, err error) {
	query := "SELECT id, first_name, last_name, email, phone FROM contacts WHERE phone = $1;"
	row := r.db.QueryRow(query, phone)
	err = row.Scan(&contact.ID, &contact.FirstName, &contact.LastName, &contact.Email, &contact.Phone)
	return contact, err
}

func (r *ContactsRepositoryPostgreSQL) GetByEmail(email string) (contact model.Contact, err error) {
	query := "SELECT id, first_name, last_name, email, phone FROM contacts WHERE email = $1;"
	row := r.db.QueryRow(query, email)
	err = row.Scan(&contact.ID, &contact.FirstName, &contact.LastName, &contact.Email, &contact.Phone)
	return contact, err
}

func (r *ContactsRepositoryPostgreSQL) SearchByName(n string) (contacts []model.Contact, err error) {
	query := "SELECT id, first_name, last_name, email, phone FROM contacts WHERE first_name = $1 or last_name like $1;"
	rows, err := r.db.Query(query, n)

	if err != nil {
		return contacts, err
	}

	defer rows.Close()
	for rows.Next() {
		var contact model.Contact
		err = rows.Scan(&contact.ID, &contact.FirstName, &contact.LastName, &contact.Email, &contact.Phone)
		if err != nil {
			return nil, err
		}
		contacts = append(contacts, contact)
	}

	return contacts, nil
}

func (r *ContactsRepositoryPostgreSQL) Delete(id uint) error {
	query := "DELETE FROM contacts WHERE id = $1;"
	_, err := r.db.Exec(query, id)
	return err
}
