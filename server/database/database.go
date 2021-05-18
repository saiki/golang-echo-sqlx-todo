package database

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var initDDL = `
DROP TABLE IF EXISTS todo;
CREATE TABLE todo (
  id INTEGER NOT NULL PRIMARY KEY,
  task text DEFAULT '',
  finished boolean default 0
)`

type DatasourceConfig struct {
	DriverName     string
	DataSourceName string
}

var Config DatasourceConfig

func InitDatabase() {
	db, err := Open()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	db.MustExec(initDDL)
	return
}

func Open() (*sqlx.DB, error) {
	return sqlx.Open(Config.DriverName, Config.DataSourceName)
}
