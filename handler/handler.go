package handler

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type DB *gorm.DB

type Handler struct {
	db *DB
}

func New(db *DB) *Handler {
	return &Handler{
		db: db,
	}
}

// user model

type User struct {
	ID int `json:"id"`
	Name  string `json:"name"`
	Token string `json:"token"`
}


func (db DB) CreateUser(name string, token string) error {
	return db.Create(&User{
		Name:  name,
		Token: token,
	}).Error
}

func (db DB) GetUser(token string) (User, error) {
	u := User{}
	err := db.Where("token = ?", token).Find(&u).Error

	if err != nil {
		fmt.Println(err)
	}

	return u, err
}

func (db DB) UpdateUser(name string, token string) error {

	return db.Model(&User{}).Where("token = ?", token).Update(&User{
		Name: name,
	}).Error
}

func (db DB) GetUserID(token string) (User, error) {

	u := User{}
	err := db.Where("token = ?", token).Find(&u).Error

	if err != nil {
		fmt.Println(err)
	}

	return u, err
}


// character model

type Character struct {
	ID   string `json:"characterID"`
	Name string `json:"name"`
}

func (db DB) GetCharacter(characterID string) (Character, error) {
	c := Character{}
	err := db.Where("id = ?", characterID).Find(&c).Error

	if err != nil {
		fmt.Println(err)
	}

	return c, err
}


// user_character model

type UserChara struct {
	UserID int `json:"userID"`
	CharacterID string `json:"characterID"`
}

func (ucl *UserChara) TableName() string {
	return "user_character"
}

type UserCharas []UserChara

func (db DB) GetUserCharacter(userID int) (UserCharas, error) {
	uc := UserCharas{}
	err := db.Where("user_id = ?", userID).Find(&uc).Error

	if err != nil {
		fmt.Println(err)
	}
	return uc, err
}

func (db DB) CreateUserCharacter(userID int, characterID string) error {
	return db.Create(&UserChara{
		UserID: userID,
		CharacterID: characterID,
	}).Error
}


// gacha model

type GachaRate struct {
	Rate        int    `json:"rate"`
	CharacterID string `json:"characterID"`
}

type GachaRates []GachaRate

func (db DB) GetGachaRate() (GachaRates, error) {
	gr := GachaRates{}
	err := db.Find(&gr).Error

	if err != nil {
		fmt.Println(err)
	}

	return gr, err
}


