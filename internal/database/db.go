package database

import (
	"errors"

	"crypto/sha256"

	"encoding/hex"

	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

func UserExist(email string) bool {
	_, err := Db.Query("select * from users where email = $1", email)
	if err != nil {
		return false
	}

	return true
}

func RegUser(name, lastname, email, password, phone, male, birthday string) error {
	h := sha256.New()
	h.Write([]byte(password))

	hash := hex.EncodeToString(h.Sum(nil))

	if !UserExist(email) {
		_, err := Db.Exec("insert into users (name, lastname, email, password) values ($1, $2, $3, $4)", name, lastname, email, hash)
		if err != nil {
			return err
		}

		return nil
	}

	return errors.New("User is already regist.")
}

func LoginUser(name, lastname, email, password string) (error, bool) {
	row := Db.QueryRow("select name, lastname, password from users where email = $1", email)

	h := sha256.New()
	h.Write([]byte(password))

	hash := hex.EncodeToString(h.Sum(nil))

	var nameDb string
	var lastnameDb string
	var passwordHash string

	err := row.Scan(&nameDb, &lastnameDb, &passwordHash)

	if err != nil {
		return err, false
	}

	if nameDb == name && lastnameDb == lastname && passwordHash == hash {
		return nil, true
	}

	return nil, false
}

func GetUsers() (map[int]map[string]string, error) {
	var Users = make(map[int]map[string]string)

	rows, err := Db.Query("select name, lastname, nickname from users")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {

		var (
			name     string
			nickname string
			lastname string
			id       int
		)

		if err := rows.Scan(&name, &lastname, &name, &nickname); err != nil {
			return nil, err
		}

		Users[id] = map[string]string{
			"name":     name,
			"lastname": lastname,
			"nickname": nickname,
		}

	}

	return Users, nil
}
