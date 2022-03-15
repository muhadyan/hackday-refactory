package main

import (
	"be/db"
	"be/route"
)

func main() {
	db.Init()
	e := route.Init()

	e.Logger.Fatal(e.Start(":8080"))
}