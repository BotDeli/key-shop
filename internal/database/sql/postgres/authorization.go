package postgres

import (
	"errors"
	"key-shop/pkg/errorHandle"
	"key-shop/pkg/hasher"
	"log"
)

type User struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

// if errorHandle is nil => authorization successful

type Authorization interface {
	Registration(User) error
	Login(User) error
}

var (
	ErrorUserAlreadyRegistered      = errors.New("user already registered")
	ErrorShortLoginOrPassword       = errors.New("short login or password")
	ErrorUserDontExist              = errors.New("user dont registered")
	ErrorDontCorrectLoginOrPassword = errors.New("dont correct login or password")
	ErrorRegistration               = errors.New("error registration")
	ErrorGetHashedPassword          = errors.New("error get hash password")
)

func (p Postgres) Registration(user User) error {
	if shortLoginOrPassword(user) {
		return ErrorShortLoginOrPassword
	}
	if ExistsLogin(p, user.Login) {
		return ErrorUserAlreadyRegistered
	}
	err := insertRow(p.Database, user.Login, user.Password)
	return err
}

func ExistsLogin(p Postgres, login string) bool {
	password, err := p.GetHashedPassword(login)
	return err == nil && password != ""
}

func (p Postgres) GetHashedPassword(login string) (string, error) {
	query := `SELECT password FROM users WHERE login = $1`
	var password string
	err := p.Database.QueryRow(query, login).Scan(&password)
	if err != nil {
		log.Println(errorHandle.ErrorFormat(path, "authorization.go", "GetHashedPassword", err))
		err = ErrorGetHashedPassword
	}
	return password, err
}

func shortLoginOrPassword(user User) bool {
	return len(user.Login) < 6 || len(user.Password) < 6
}

func insertRow(database DB, login string, password string) error {
	query := `INSERT INTO users (login, password) VALUES ($1, $2)`
	hashedPassword := hasher.Hashing(password)
	if _, err := database.Exec(query, login, hashedPassword); err != nil {
		log.Println(errorHandle.ErrorFormat(path, "authorization.go", "insertRow", err))
		return ErrorRegistration
	}
	return nil
}

func (p Postgres) Login(user User) error {
	if shortLoginOrPassword(user) {
		return ErrorShortLoginOrPassword
	}
	if !ExistsLogin(p, user.Login) {
		return ErrorUserDontExist
	}
	return authentication(p, user)
}

func authentication(p Postgres, user User) error {
	expectedHashedPassword, err := p.GetHashedPassword(user.Login)
	if err == nil {
		hashedPassword := hasher.Hashing(user.Password)
		if hashedPassword != expectedHashedPassword {
			return ErrorDontCorrectLoginOrPassword
		}
	}
	return err
}
