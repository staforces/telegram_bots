package main

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	"log"
)

var db *sql.DB

func insertdb(tprices int) error {
	dsn := fmt.Sprintf("host=%s user=%s dbname=%s port=%s sslmode=disable", cfg.PgHost, cfg.PgUser, cfg.PgBase, cfg.PgPort)
	data := `INSERT INTO tprice(wowtoken) VALUES($1,$2);`
	var err error
	db, err = sql.Open("pgx", dsn)
	if err != nil {
		log.Fatal(err)
		return err
	}
	if _, err = db.Exec(data, tprices); err != nil {
		return err
	}

	return nil
}
