package main

import (
	"db/postgres"
	_ "db/postgres"
)

func main() {
	postgres.SqlSelect()
}
