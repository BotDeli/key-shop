package postgres

import (
	"errors"
	"key-shop/pkg/errorHandle"
	"log"
)

type ItemsDisplay interface {
	AddItem(login string, item Item) error
	DeleteItem(login string, item Item) error
	ExistsItem(item Item) bool
}

type Item struct {
	Name        string
	Description string
	Count       string
	Cost        string
}

var (
	ErrorAddItem    = errors.New("error add item")
	ErrorDeleteItem = errors.New("error delete item")
)

func (p Postgres) AddItem(login string, item Item) error {
	query := "INSERT INTO items(name, description, count, cost, seller) VALUES($1, $2, $3, $4, $5)"
	return executeQueryAllParams(p, "AddItem", query, login, item, ErrorAddItem)
}

func (p Postgres) DeleteItem(login string, item Item) error {
	query := "DELETE FROM items WHERE name = $1 AND description = $2 AND count = $3 AND cost = $4 AND seller = $5"
	return executeQueryAllParams(p, "DeleteItem", query, login, item, ErrorDeleteItem)
}

func (p Postgres) ExistsItem(item Item) bool {
	seller, err := finderSeller(p, item)
	return err == nil && seller != ""
}

func executeQueryAllParams(p Postgres, funcResponse, query, login string, item Item, funcErr error) error {
	_, err := p.Database.Exec(query,
		item.Name,
		item.Description,
		item.Count,
		item.Cost,
		login,
	)
	if err != nil {
		log.Println(errorHandle.ErrorFormat(path, "items.go", funcResponse, err))
		err = funcErr
	}
	return err
}

func finderSeller(p Postgres, item Item) (string, error) {
	query := "SELECT seller FROM items WHERE name = $1 AND description = $2 AND count = $3 AND cost = $4"
	var seller string
	err := p.Database.QueryRow(query,
		item.Name,
		item.Description,
		item.Count,
		item.Cost,
	).Scan(&seller)
	return seller, err
}
