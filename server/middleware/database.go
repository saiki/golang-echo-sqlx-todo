package middleware

import (
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	_ "github.com/mattn/go-sqlite3"

	"github.com/saiki/golang-echo-sqlx-todo/server/database"
)

type CustomContext struct {
	echo.Context
	DB *sqlx.DB
}

func DataSourceMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			ctx := &CustomContext{
				Context: c,
			}
			db, err := database.Open()
			log.Println("Datasource connected")
			if err != nil {
				return err
			}
			defer func() {
				db.Close()
				log.Println("Datasource closed")
			}()
			ctx.DB = db
			err = next(ctx)
			return err
		}
	}
}
