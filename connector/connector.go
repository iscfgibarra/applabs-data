package connector

import (
	"database/sql"
	"fmt"
	"github.com/iscfgibarra/applabs-data/drivers"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"sync"
)

var (
	once sync.Once
)

var SqlCnn *SqlConnector

type SqlConnector struct {
	Db             *sql.DB
	dataSourceName string
	Driver         drivers.Driver
}

func InitSqlConnector(dataSourceName string, driver drivers.Driver) {
	if SqlCnn == nil {
		SqlCnn = newSqlConnector(dataSourceName, driver)
		_ = SqlCnn.GetConnection()
	}
}

func newSqlConnector(dataSourceName string, driver drivers.Driver) *SqlConnector {
	return &SqlConnector{
		dataSourceName: dataSourceName,
		Driver:         driver,
	}
}

func (sc *SqlConnector) GetConnection() *sql.DB {
	once.Do(func() {
		var err error
		sc.Db, err = sql.Open(string(sc.Driver), sc.dataSourceName)
		if err != nil {
			log.Fatalf("can't open DB %v", err)
		}

		if err = sc.Db.Ping(); err != nil {
			log.Fatalf("can't ping DB %v", err)
		}

		fmt.Println("Connect to " + string(sc.Driver))
	})

	return sc.Db
}
