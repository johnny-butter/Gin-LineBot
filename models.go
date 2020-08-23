package main

type GoKeyword struct {
	Id          int    `db:"id"`
	Keyword     string `db:"keyword"`
	ResponseCls string `db:"response_cls"`
}
