package main

import (
	"errors"
	//"log"
	"strings"
	//"time"

	"github.com/badoux/checkmail"
	"gorm.io/gorm"
)

type Person struct {
	gorm.Model
	//ID       int    `json:"id"`
	Fullname string `json:"fullname"`
	Password string `gorm:"size:100" json:"password" `
	Email    string `json:"email"`
	Gender   string `json:"gender"`
	Book     *Book  `json:"book"` //`gorm:"ForeignKey:PersonID"` //`gorm:"ForeignKey:PersonID;AssociationForeignKey:bookID"`
	//BookID   uint   //`gorm:"Foreignkey:book_id"`
}

// REmove white spaces
func (p *Person) Prepare() {
	p.Fullname = strings.TrimSpace(p.Fullname)
	p.Email = strings.TrimSpace(p.Email)
	p.Gender = strings.TrimSpace(p.Gender)
}
func (p *Person) BeforeSave() error {
	hashedPassword, err := HashPassword(p.Password)
	if err != nil {
		return err
	}
	p.Password = string(hashedPassword)
	return nil
}

func (p *Person) SavePerson(db *gorm.DB) (*Person, error) {
	var err error
	err = db.Debug().Create(&p).Error
	if err != nil {
		return &Person{}, err
	}
	return p, nil
}

func (p *Person) GetUser(db *gorm.DB) (*Person, error) {
	account := &Person{}
	if err := db.Debug().Table("people").Where("email = ?", p.Email).First(account).Error; err != nil {
		return nil, err
	}
	return account, nil
}

func GetAllPersons(db *gorm.DB) (*[]Person, error) {
	people := []Person{}
	if err := db.Debug().Table("people").Find(&people).Error; err != nil {
		return &[]Person{}, err
	}
	return &people, nil
}

func GetPeopleAndBooks(db *gorm.DB) (*[]Person, error) {
	pnb := []Person{}
	if err := db.Debug().Preload("Book").
		Joins("INNER JOIN books ON books.person_id = people.id").Find(&pnb).Error; err != nil {
		return &[]Person{}, err
	}
	return &pnb, nil
}

func GetAllUsersWithoutBks(db *gorm.DB) (*[]Person, error) {
	people := []Person{}
	if err := db.Raw("select books.*, people.* from books right outer join people on books.person_id=people.id where title is null").Scan(&people).Error; err != nil {
		return &[]Person{}, err
	}
	return &people, nil
}

func GetAllUsersWithBks(db *gorm.DB) (*[]Person, error) {
	people := []Person{}
	if err := db.Raw("select books.*, people.* from books inner join people on books.person_id=people.id where title is not null").Scan(&people).Error; err != nil {
		return &[]Person{}, err
	}
	return &people, nil
}

func GetUserWithId(id int, db *gorm.DB) (*Person, error) {
	//params := mux.Vars(r)

	user := &Person{}
	if err := db.Debug().Table("people").Where("id= ?", id).First(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (p *Person) UpdateUser(id int, db *gorm.DB) (*Person, error) {
	// err := p.BeforeSave()
	// if err != nil {
	// 	log.Fatal(err)

	if err := db.Debug().Table("people").Where("id= ?", id).Updates(Person{
		Fullname: p.Fullname,
		Email:    p.Email,
		Password: p.Password,
		Gender:   p.Gender,
		Book:     p.Book}).Error; err != nil {
		return &Person{}, err
	}
	return p, nil

	// if db.Error != nil {
	// 	return &Person{}, db.Error
	// }
	// err = db.Debug().Model(&Person{}).Where("id =?", uid).Take(&p).Error
	// if err != nil {
	// 	return &Person{}, err
	// }
	// return p, nil
}

func (p *Person) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if p.Email == "" {
			return errors.New("required password")
		}
		if p.Fullname == "" {
			return errors.New("required fullname")
		}
		if p.Password == "" {
			return errors.New("required password")
		}
		if p.Gender == "" {
			return errors.New("required gender")
		}
		if err := checkmail.ValidateFormat(p.Email); err != nil {
			return errors.New("invalid email")
		}
		return nil
	case "login":
		if p.Password == "" {
			return errors.New("required password")
		}
		if p.Email == "" {
			return errors.New("required Email")

		}
		if err := checkmail.ValidateFormat(p.Email); err != nil {
			return errors.New("invalid email")
		}
		return nil

	default:
		if p.Email == "" {
			return errors.New("required email")
		}
		if p.Fullname == "" {
			return errors.New("required username")
		}
		if p.Password == "" {
			return errors.New("required password")
		}
		if p.Gender == "" {
			return errors.New("required gender")
		}
		if err := checkmail.ValidateFormat(p.Email); err != nil {
			return errors.New("invalid email")
		}
		return nil

	}

}
func (p *Person) FormatError(err string) error {
	if strings.Contains(err, "fullname") {
		return errors.New("nickname already taken")
	}
	if strings.Contains(err, "email") {
		return errors.New("email already taken")
	}
	if strings.Contains(err, "hashedPassword") {
		return errors.New("incorrect password")
	}
	return nil
}
