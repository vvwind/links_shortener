package db

import (
	"context"

	"fmt"
	"github.com/jackc/pgx/v5"
	_ "github.com/lib/pq"
	"os"
)

type DbData struct {
	conn *pgx.Conn
}

func Init() *DbData {
	return &DbData{}
}

func (d *DbData) StartDB() error {
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	conn, err := pgx.Connect(context.Background(), "postgres://postgres:postgres@localhost:5435/postgres")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return err
	}
	d.conn = conn

	return nil
}
func (d *DbData) WriteOriginal(f_link string, s_link string) error {
	_, err := d.conn.Exec(context.Background(), "INSERT INTO original_links(original_link,new_link) VALUES ($1, $2);", f_link, s_link)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return err
	}
	return nil

}
func (d *DbData) WriteNew(f_link string, s_link string) error {
	_, err := d.conn.Exec(context.Background(), "INSERT INTO new_links(new_link, original_link) VALUES ($1, $2);", f_link, s_link)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return err
	}
	return nil
}

func (d *DbData) Reveal(s_link string) (string, error) {
	var orig string
	err := d.conn.QueryRow(context.Background(), "select original_link from new_links where new_link=$1", s_link).Scan(&orig)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return "", err
	}

	return orig, nil
}
