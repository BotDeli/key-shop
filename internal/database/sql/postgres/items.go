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
	Name        string `json:"name"`
	Description string `json:"description"`
	Count       string `json:"count"`
	Cost        string `json:"cost"`
}

var (
	ErrorAddItem         = errors.New("error add item")
	ErrorDeleteItem      = errors.New("error delete item")
	ErrorLimitCountItems = errors.New("error limit count items")
)

func (p Postgres) AddItem(login string, item Item) error {
	if limitedCountItems(p, login) {
		return ErrorLimitCountItems
	}
	query := "INSERT INTO items(name, description, count, cost, seller) VALUES($1, $2, $3, $4, $5)"
	return executeQueryAllParams(p, "AddItem", query, login, item, ErrorAddItem)
}

func limitedCountItems(p Postgres, login string) bool {
	query := "SELECT COUNT(*) FROM items WHERE seller = $1"
	var count int
	err := p.Database.QueryRow(query, login).Scan(&count)
	return err == nil && count >= 10
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
