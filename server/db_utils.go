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

func NewDataHandler(host, port, password string) (db DataHandler) {
	db.pg, _ = gorm.Open(postgres.New(postgres.Config{
		DSN:                  "host=" + host + " user=postgres password=" + password + " dbname=postgres port=" + port + " sslmode=disable",
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

func (db DataHandler) NoteUpdate(id uint, title, content string) (*Note, error) {
	note := &Note{
		Id: id,
	}

	res := db.pg.First(&note)

	if res.Error != nil {
		return &Note{}, res.Error
	}

	note.Title = title
	note.Content = content

	res = db.pg.Save(&note)
	return note, res.Error
}

func (db DataHandler) GetNote(id uint) (*Note, error) {

	note := &Note{
		Id: id,
	}
	res := db.pg.Where(note).First(&note)

	return note, res.Error
}

func (db DataHandler) GetUserNotes(oid uint) ([]Note, error) {
	notes := make([]Note, 10)
	note := &Note{
		Owner: oid,
	}
	res := db.pg.Where(note).Find(&notes)

	return notes, res.Error
}

func getHash(str string) (sha string) {
	h := sha256.New()
	h.Write([]byte(str))
	b := h.Sum(nil)
	sha = base64.StdEncoding.EncodeToString(b)
	return
}
