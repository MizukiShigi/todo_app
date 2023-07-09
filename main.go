package main

import (
	"udemy/app/controllers"
	"udemy/app/models"
	"fmt"
)

func main() {
	fmt.Println(models.Db)
	controllers.StartMainServer()
}