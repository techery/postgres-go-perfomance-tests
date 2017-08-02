package config

import (
	"log"

	"github.com/jackc/pgx"
)

var db *pgx.ConnPool
var err error

// InitDatabase postgres db connection
func InitDatabase(dbHost string, dbUser string, dbPass string, dbName string, dbPort uint16, maxConnectionsInPool int) *pgx.ConnPool {
	var config pgx.ConnPoolConfig

	config.Host = dbHost
	config.User = dbUser
	config.Password = dbPass
	config.Database = dbName
	config.Port = dbPort

	config.MaxConnections = maxConnectionsInPool

	connPool, err := pgx.NewConnPool(config)
	if err != nil {
		log.Fatal(err)
	}
	return connPool
}
