package postgres

import (
	"database/sql"
	"errors"
	"key-shop/pkg/errorHandle"
	"log"
)

var (
	ErrorGetCountItems = errors.New("error get count items")
	ErrorGetAllItems   = errors.New("error get all items")
	ErrorGetMyRows     = errors.New("error get my rows")
)

const (
	countItemsPerPage = 25
)

type Page struct {
	Items [][]string `json:"items"`
}

type PageDisplay interface {
	GetCountPages() (int, error)
	GetPageAllItems(numberPage int) (Page, error)
	GetPageMyItems(login string) (Page, error)
}

func (p Postgres) GetCountPages() (int, error) {
	allItems, err := getCountItems(p)
	if err != nil {
		return 0, err
	}
	allPages := allItems / countItemsPerPage
	if allItems%countItemsPerPage != 0 {
		allPages++
	}
	return allPages, nil
}

func getCountItems(p Postgres) (int, error) {
	query := "SELECT COUNT(*) FROM items"
	var count int
	err := p.Database.QueryRow(query).Scan(&count)
	if err != nil {
		log.Println(errorHandle.ErrorFormat(path, "pageItems", "getCountItems", err))
		err = ErrorGetCountItems
	}
	return count, err
}

func (p Postgres) GetPageAllItems(numberPage int) (Page, error) {
	page := Page{}
	rows, err := getAllItems(p, numberPage)
	if err != nil {
		return page, err
	}
	for rows.Next() {
		var (
			name        string
			description string
			count       string
			cost        string
			seller      string
		)
		err := rows.Scan(&name, &description, &count, &cost, &seller)
		if err != nil {
			log.Println(errorHandle.ErrorFormat(path, "pageItems", "GetPage", err))
		}
		page.Items = append(page.Items, []string{
			name,
			count,
			cost,
			seller,
		})
	}
	return page, err
}

func getAllItems(p Postgres, numberPage int) (*sql.Rows, error) {
	skipItems := (numberPage - 1) * countItemsPerPage
	query := "SELECT * FROM items OFFSET $1 LIMIT $2"
	rows, err := p.Database.Query(query, skipItems, countItemsPerPage)
	if err != nil {
		log.Println(errorHandle.ErrorFormat(path, "pageItems", "GetPage", err))
		err = ErrorGetAllItems
	}
	return rows, err
}

func (p Postgres) GetPageMyItems(login string) (Page, error) {
	page := Page{}
	rows, err := getMyRows(p, login)
	if err != nil {
		return page, err
	}

	for rows.Next() {
		var (
			name        string
			description string
			count       string
			cost        string
			seller      string
		)
		err := rows.Scan(&name, &description, &count, &cost, &seller)
		if err != nil {
			log.Println(errorHandle.ErrorFormat(path, "items.go", "GetPageMyItems", err))
		}
		page.Items = append(page.Items, []string{
			name,
			count,
			cost,
		})
	}
	return page, nil
}

func getMyRows(p Postgres, login string) (*sql.Rows, error) {
	query := "SELECT * FROM items WHERE seller = $1"
	rows, err := p.Database.Query(query, login)
	if err != nil {
		log.Println(errorHandle.ErrorFormat(path, "items.go", "getMyRows", err))
		err = ErrorGetMyRows
	}
	return rows, err
}
