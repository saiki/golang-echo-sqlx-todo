package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/mattn/go-sqlite3"

	"github.com/saiki/golang-echo-sqlx-todo/server/api"
	"github.com/saiki/golang-echo-sqlx-todo/server/database"
	mymiddleware "github.com/saiki/golang-echo-sqlx-todo/server/middleware"
)

func init() {
	database.Config = database.DatasourceConfig{
		DriverName:     "sqlite3",
		DataSourceName: "./my.db",
	}
	database.InitDatabase()
}

func main() {
	port := flag.Int("port", 8080, "http port.")
	flag.Parse()
	e := echo.New()
	// Debug mode
	e.Debug = true
	e.Use(middleware.Logger())
	e.Use(mymiddleware.DataSourceMiddleware())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	// Handler
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/api/todo", api.ListTodo)
	e.POST("/api/todo", api.CreateTodo)
	e.PUT("/api/todo/:id", api.UpdateTodo)
	e.DELETE("/api/todo/:id", api.DeleteTodo)

	errC := make(chan error)

	go func() {
		if err := e.Start(fmt.Sprintf(":%d", *port)); err != nil {
			errC <- err
		}
	}()

	quitC := make(chan os.Signal)
	signal.Notify(quitC, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err := <-errC:
		panic(err)
	case <-quitC:
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := e.Shutdown(shutdownCtx); err != nil {
			errC <- err
		}
	}

}
