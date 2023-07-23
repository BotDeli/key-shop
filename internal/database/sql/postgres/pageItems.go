package postgres

import (
	"errors"
	"key-shop/pkg/errorHandle"
	"log"
)

var (
	ErrorGetCountItems = errors.New("error get count items")
	ErrorGetPage       = errors.New("error get page")
)

type Page struct {
	Items [][]string `json:"items"`
}

type PageDisplay interface {
	GetCountPages() (int, error)
	GetPage(numberPage int) (Page, error)
}

func (p Postgres) GetCountPages() (int, error) {
	allItems, err := getCountItems(p)
	if err != nil {
		return 0, err
	}
	allPages := allItems % 20
	if allItems%20 != 0 {
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

func (p Postgres) GetPage(numberPage int) (Page, error) {
	skipItems := (numberPage - 1) * 20
	limitItems := 20
	page := Page{}
	query := "SELECT * FROM items OFFSET $1 LIMIT $2"

	rows, err := p.Database.Query(query, skipItems, limitItems)
	if err != nil {
		log.Println(errorHandle.ErrorFormat(path, "pageItems", "GetPage", err))
		err = ErrorGetPage
	} else {
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
	}

	return page, err
}
