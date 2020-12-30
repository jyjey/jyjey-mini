package main

import (
	"jyjey-mini/mysql"
	"jyjey-mini/router"
)

func main() {
	defer mysql.DB.Close()
	router.InitJYJEYRouter()
}
