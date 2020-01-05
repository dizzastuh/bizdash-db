package model

import (
    "log"
    "fmt"
    "os"
    "time"

    client "github.com/influxdata/influxdb1-client/v2"
    db "github.com/dizzastuh/bizdash-db/internal"
    . "github.com/nswekosk/fred_go_toolkit"
)

// time formatting must be done with  Mon Jan 2 15:04:05 MST 2006
const YYYY_MM_DD = "2006-01-02"

func InsertFredObs(ft *FredType, name string) {
    fmt.Printf("Inserting observations for %s\n", name)
    con := db.Connect()

    config := client.BatchPointsConfig {
        Database: os.Getenv("DB_NAME"),
    }

    bp, _ := client.NewBatchPoints(config)

    for i:= 1; i < len(ft.Observations); i++ {
        obs := ft.Observations[i]
        tags := map[string]string{
            "source": "fred",
            "name": name,
        }

        fields := map[string]interface{}{
            "value":   obs.Value,
        }

        timestamp, err := time.Parse(YYYY_MM_DD, obs.Date)
        evaluate(err)

        point, err := client.NewPoint(
            name,
            tags,
            fields,
            timestamp,
        )

        evaluate(err)
        bp.AddPoint(point)
    }

    err := con.Write(bp)
    con.Close()

    evaluate(err)
    fmt.Println("Done")
}

func FetchFred() {
    // TODO: retrieve the data
}

func evaluate(err error) {
    if err != nil {
        log.Fatal(err)
    }
}
