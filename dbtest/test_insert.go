package dbtest

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx"
	"github.com/jackc/pgx/pgtype"
)

func insert(db *pgx.ConnPool) {

	var query = "INSERT INTO session_event (session_id, event_type, event_timestamp, binary_payload) VALUES ($1, $2, $3, $4)"

	payload := make([]byte, 1024*2)

	var err error
	_, err = db.Prepare("insert_event", query)
	if err != nil {
		log.Fatal(err)
	}

	finishChan := make(chan int)

	for count := 0; count < 10; count++ {
		go func(c chan int) {
			for batch := 0; batch < 100; batch++ {
				batchInsert := db.BeginBatch()
				for inserts := 0; inserts < 1000; inserts++ {
					batchInsert.Queue("insert_event",
						[]interface{}{1, byte(10), time.Now(), payload},
						[]pgtype.OID{pgtype.Int4OID, pgtype.Int2OID, pgtype.TimestampOID, pgtype.ByteaOID}, nil,
					)
				}
				err := batchInsert.Send(context.Background(), nil)
				if err != nil {
					log.Fatal(err)
				}
				err = batchInsert.Close()
				if err != nil {
					log.Fatal(err)
				}
			}
			c <- 1
		}(finishChan)
	}

	finishedGophers := 0
	finishLoop := false
	for {
		if finishLoop {
			break
		}
		select {
		case n := <-finishChan:
			finishedGophers += n
			if finishedGophers == 10 {
				finishLoop = true
			}
		}
	}
}
