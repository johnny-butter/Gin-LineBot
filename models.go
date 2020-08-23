package main

import (
	"log"

	"github.com/gobuffalo/pop"
)

func GetDBConnect() (db *pop.Connection) {
	db, err := pop.Connect("development")
	if err != nil {
		log.Panic(err)
	}

	return
}

type GoKeyword struct {
	Id          int    `db:"id"`
	Keyword     string `db:"keyword"`
	ResponseCls string `db:"response_cls"`
}
