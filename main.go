package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nyae44/estore/controllers"
	"github.com/nyae44/estore/database"
	"github.com/nyae44/estore/middleware"
	"github.com/nyae44/estore/routes"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	//app := controllers.NewApplication(database.ProductData(database.Client))
}
