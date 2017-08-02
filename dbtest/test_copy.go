package dbtest

import (
	"log"
	"time"

	"github.com/jackc/pgx"
	"github.com/techery/payload-perfomance-tests/config"
)

// Copy test
func Copy(db *pgx.ConnPool, c *config.Config) {
	defer db.Close()
	payload := make([]byte, 1024*c.Bench.PayloadSize)

	finishChan := make(chan int)

	for count := 0; count < c.Bench.WorkersNum; count++ {
		go func(e chan int, c *config.Config) {

			for j := 0; j < c.Bench.OperationsPerTask; j++ {

				inputRows := [][]interface{}{}

				for i := 0; i < c.Bench.RowsPerTask; i++ {
					inputRows = append(inputRows, []interface{}{1, byte(10), time.Now(), payload})
				}

				copyCount, err := db.CopyFrom(pgx.Identifier{"session_event"}, []string{"session_id", "event_type", "event_timestamp", "binary_payload"}, pgx.CopyFromRows(inputRows))

				if err != nil {
					log.Fatal(err)
				}

				if copyCount != len(inputRows) {
					log.Fatal(len(inputRows), copyCount)
				}
			}
			e <- 1
		}(finishChan, c)
	}

	// wait for all workers to complere
	finishedWokers := 0
	finishLoop := false
	for {
		if finishLoop {
			break
		}
		select {
		case n := <-finishChan:
			finishedWokers += n
			if finishedWokers == c.Bench.WorkersNum {
				finishLoop = true
			}
		}
	}
}
