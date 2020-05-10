package middleware

import "github.com/jmoiron/sqlx"
import "github.com/labstack/echo"
import "log"

type DatasourceConfig struct {
	DriverName     string
	DataSourceName string
}

func DataSourceMiddleware(config DatasourceConfig) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			db, err := sqlx.Open(config.DriverName, config.DataSourceName)
			log.Println("Datasource connected")
			if err != nil {
				return err
			}
			defer func() {
				db.Close()
				log.Println("Datasource closed")
			}()
			c.Set("dataSource", db)
			err = next(c)
			return err
		}
	}
}
