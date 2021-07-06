package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func freak(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	dsn := "n9e2tq4wo6uhxslh:jmm9tpr4pjtkyazi@tcp(ao9moanwus0rjiex.cbetxkdyhwsb.us-east-1.rds.amazonaws.com:3306)/apc4exudm9mlh3ul"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	freak((err))
	println(db)
}
