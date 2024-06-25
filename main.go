package main

import (
	"fmt"

	"local.package/golang_todo/app/controllers"
	"local.package/golang_todo/app/models"
)

func main() {
	fmt.Println(models.Db)
	go controllers.StartMainServer()

	for {

	}
}
