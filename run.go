package main

import (
	"fmt"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/techery/payload-perfomance-tests/config"
	"github.com/techery/payload-perfomance-tests/dbtest"
)

func main() {

	start := time.Now()
	c := config.Read()

	fmt.Println("Using config:")
	spew.Dump(c)
	db := config.InitDatabase(c.Db.Host, c.Db.User, c.Db.Password, c.Db.DbName, c.Db.Port, c.Db.MaxConnectionCount)
	dbtest.Copy(db, c)
	elapsed := time.Since(start)

	fmt.Println(" ")
	fmt.Printf("%s", elapsed)
	fmt.Println(" ")
}
