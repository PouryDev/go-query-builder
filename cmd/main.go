package main

import (
	"fmt"
	"learn/internal/db"
)

func main() {
	qb := db.QueryBuilder{}
	query := qb.Table("products").
		Select().
		Where("price", ">", db.IntValue(2000)).
		Build()

	fmt.Println(query)
}
