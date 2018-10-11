package main

import (
	"crypto/sha1"
	"encoding/base64"
	"io"
	"log"

	_ "github.com/go-sql-driver/mysql" // importing mysql driver
	"github.com/jmoiron/sqlx"
)

const salt = "justAnotherSalt"

var db *sqlx.DB

func setupDB() {
	var err error
	db, err = sqlx.Connect("mysql", "root:root1234@tcp(localhost:3306)/kacago?parseTime=true")
	// log.Println("------>", db.DB)
	if err != nil {
		log.Printf("[setupDB] error:", err)
	}
}

func isValidUser(user, pass string) bool {
	h := sha1.New()
	io.WriteString(h, pass+salt)
	passStr := base64.URLEncoding.EncodeToString(h.Sum(nil))

	var resp struct {
		ID   int64  `db:"id"`
		User string `db:"user"`
		Pass string `db:"pass"`
	}
	err := db.Get(&resp, "select id, user, pass from user where user = ?", user)
	if err != nil {
		log.Println("[isValidUser] error:", err)
		return false
	}

	if passStr == resp.Pass {
		return true
	}
	return false
}
