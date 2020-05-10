package main

import "github.com/labstack/echo"
import "github.com/labstack/echo/middleware"
import _ "github.com/mattn/go-sqlite3"
import "net/http"
import mymiddleware "github.com/saiki/golang-echo-sqlx-todo/server/middleware"

func main() {
	e := echo.New()

	// Debug mode
	e.Debug = true
	config := mymiddleware.DatasourceConfig{
		DriverName:     "sqlite3",
		DataSourceName: "./my.db",
	}

	e.Use(middleware.Logger())
	e.Use(mymiddleware.DataSourceMiddleware((config)))

	// Handler
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
