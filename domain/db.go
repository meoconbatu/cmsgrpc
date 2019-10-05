package domain

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var store = newDB()

// PgStore is an exported type
type PgStore struct {
	DB *sql.DB
}

func newDB() *PgStore {
	db, err := sql.Open("postgres", "user=cms dbname=cms sslmode=disable")
	if err != nil {
		panic(err)
	}
	return &PgStore{DB: db}
}

// GetPages return all pages from pages table
func GetPages() ([]*Page, error) {
	rows, err := store.DB.Query("SELECT * FROM pages")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	pages := []*Page{}
	for rows.Next() {
		var p Page
		err = rows.Scan(&p.Id, &p.Title, &p.Content)
		if err != nil {
			return nil, err
		}
		pages = append(pages, &p)
	}
	return pages, nil
}

// GetPage return a page
func GetPage(id string) (*Page, error) {
	var p Page
	err := store.DB.QueryRow("SELECT * FROM pages WHERE id = $1", id).Scan(&p.Id, &p.Title, &p.Content)
	return &p, err
}

// CreatePage insert a new page into pages table
func CreatePage(p *Page) (int, error) {
	var id int
	err := store.DB.QueryRow("INSERT INTO pages(title, content) VALUES($1, $2) RETURNING id", p.
		Title, p.Content).Scan(&id)
	return id, err
}
