package postgres

import (
	"database/sql"
	"errors"
	_ "github.com/lib/pq"
	"key-shop/pkg/errorHandle"
	"log"
)

var (
	path           = "keyShop/internal/database/sql/postgres/"
	errDontConnect = errors.New("connection to storage is not established")
)

//go:generate go run github.com/vektra/mockery/v2@v2.32.0 --name=DB
type DB interface {
	Close() error
	Exec(query string, args ...any) (sql.Result, error)
	Query(query string, args ...any) (*sql.Rows, error)
	QueryRow(query string, args ...any) *sql.Row
}

type Postgres struct {
	Database DB
}

//go:generate go run github.com/vektra/mockery/v2@v2.32.0 --name=Storage
type Storage interface {
	Disconnect()
	Authorization
	PageDisplay
	ItemsDisplay
}

func MustNewStorage(dataSourceName string) Storage {
	database, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Fatal(errorHandle.ErrorFormatString(path, "storage.go", "NewStorage", err))
	}
	if err = database.Ping(); err != nil {
		log.Fatal(errDontConnect)
	}
	storage := Postgres{Database: database}
	return storage
}

func (p Postgres) Disconnect() {
	err := p.Database.Close()
	if err != nil {
		log.Println(errorHandle.ErrorFormatString(path, "storage.go", "Disconnect", err))
	}
}
