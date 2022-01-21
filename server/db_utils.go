package main

import (
	"crypto/sha256"
	"encoding/base64"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	Id       uint   `gorm:"primaryKey;autoIncrement"`
	Username string `gorm:"uniqueIndex"`
	Password string
	Admin    bool
}

type Note struct {
	Id      uint `gorm:"primaryKey;autoIncrement"`
	Owner   uint `gorm:"index"`
	Title   string
	Content string
}

type DataHandler struct {
	pg *gorm.DB
}

func NewDataHandler() (db DataHandler) {
	db.pg, _ = gorm.Open(postgres.New(postgres.Config{
		DSN:                  "host=localhost user=postgres password=password dbname=postgres port=5432 sslmode=disable",
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	db.pg.AutoMigrate(&User{})
	db.pg.AutoMigrate(&Note{})

	return
}

func (db DataHandler) UserAuth(username string, password string) (*User, error) {
	phash := getHash(password)

	user := &User{
		Username: username,
		Password: phash,
	}
	res := db.pg.Where(user).First(user)
	if res.Error != nil {
		return &User{}, res.Error
	}

	return user, nil
}

func (db DataHandler) UserAdd(username string, password string) (*User, error) {
	phash := getHash(password)

	user := &User{
		Username: username,
		Password: phash,
	}
	res := db.pg.Create(&user)
	if res.Error != nil {
		return &User{}, res.Error
	}
	return user, nil
}

func (db DataHandler) NoteAdd(oid uint, title, content string) (*Note, error) {
	note := &Note{
		Owner:   oid,
		Title:   title,
		Content: content,
	}
	res := db.pg.Create(&note)

	return note, res.Error
}

func (db DataHandler) GetNote(id uint) (*Note, error) {

	note := &Note{
		Id: id,
	}
	res := db.pg.Where(note).First(&note)

	return note, res.Error
}

func (db DataHandler) GetUserNotes(oid uint) (*Note, error) {

	note := &Note{
		Owner: oid,
	}
	res := db.pg.Where(note).Find(&note)

	return note, res.Error
}

func getHash(str string) (sha string) {
	h := sha256.New()
	h.Write([]byte(str))
	b := h.Sum(nil)
	sha = base64.StdEncoding.EncodeToString(b)
	return
}
