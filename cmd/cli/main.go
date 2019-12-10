package main

import (
	"fmt"
	"log"

	"github.com/burov/task_databases/pkg/model"
	"github.com/burov/task_databases/pkg/repository/memory"
)

const (
	CommandSave = iota + 1
	CommandListAll
	CommandGetByID
	CommandGetByPhone
	CommandGetByEmail
	CommandSearchByName
	CommandDelete
)

func main() {
	repository := memory.NewContactsRepositoryInMemory()

	for {
		fmt.Print(menu)
		var command int
		if _, err := fmt.Scanf("%d", &command); err != nil {
			log.Println(err)
		}

		switch command {
		case CommandSave:
			if err := Save(repository); err != nil {
				log.Println(err)
			}
		case CommandListAll:
			if err := ListAll(repository); err != nil {
				log.Println(err)
			}
		case CommandGetByID:
			if err := GetByID(repository); err != nil {
				log.Println(err)
			}
		case CommandGetByPhone:
			if err := GetByPhone(repository); err != nil {
				log.Println(err)
			}
		case CommandGetByEmail:
			if err := GetByEmail(repository); err != nil {
				log.Println(err)
			}
		case CommandSearchByName:
			if err := SearchByName(repository); err != nil {
				log.Println(err)
			}
		case CommandDelete:
			if err := Delete(repository); err != nil {
				log.Println(err)
			}
		default:
			log.Printf("command not foumd for value %d\n", command)
		}

	}

}

func ListAll (rep model.ContactsRepository) error {
	return nil
}

func GetByID (rep model.ContactsRepository) error {
	return nil
}

func GetByPhone (rep model.ContactsRepository) error {
	return nil
}

func GetByEmail (rep model.ContactsRepository) error {
	return nil
}

func SearchByName (rep model.ContactsRepository) error {
	return nil
}

func Delete (rep model.ContactsRepository) error {
	return nil
}

func Save (rep model.ContactsRepository) error {
	return nil
}


const menu = `
Please enter operation number:
  * 1 - Save
  * 2 - ListAll
  * 3 - GetByID
  * 4 - GetByPhone
  * 6 - GetByEmail
  * 7 - SearchByName
  * 8 - Delete 
  * Control + C - to exit 
`
