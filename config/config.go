package config

import (
	"runtime"

	arg "github.com/alexflint/go-arg"
)

// Db config
type Db struct {
	Host               string `arg:"-H,help:postgres host"`
	User               string `arg:"-U,help:postgres user"`
	Password           string `arg:"-P,help:postgres user password"`
	DbName             string `arg:"-D,help:postgres db to connect to"`
	Port               uint16 `arg:"-p,help:postgres port"`
	MaxConnectionCount int    `arg:"-c,help:postgres max connection count"`
}

// Bench config for processing
type Bench struct {
	WorkersNum        int `arg:"-W,help:workers number"`
	OperationsPerTask int `arg:"-O,help:operations per worker"`
	RowsPerTask       int `arg:"-R,help:rows to copy (insert) per operation"`
	PayloadSize       int `arg:"-s,help:payload size in kb"`
}

// Config fot app
type Config struct {
	Db
	Bench
}

// Read flags from input
func Read() *Config {
	c := &Config{
		Db: Db{
			Host:               "localhost",
			User:               "appspector",
			Password:           "",
			DbName:             "appspector_event_history",
			Port:               5432,
			MaxConnectionCount: runtime.NumCPU() * 2,
		},
		Bench: Bench{
			WorkersNum:        10,
			OperationsPerTask: 100,
			RowsPerTask:       1000,
			PayloadSize:       2,
		},
	}
	arg.MustParse(c)
	return c
}
