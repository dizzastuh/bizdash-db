package model

import (
    "log"
    "fmt"
    "time"

    . "github.com/dizzastuh/bizdash-db/internal"
    . "github.com/nswekosk/fred_go_toolkit"
)

// time formatting must be done with  Mon Jan 2 15:04:05 MST 2006
const YYYY_MM_DD = "2006-01-02"

func InsertFredObs(obs Observation, name string) {
    fmt.Printf("Inserting observations for %s\n", name)
    client := model.Connect()
    defer client.Close()

    config := Config{
        Database: os.Getenv("INFLUX_TABLE"),
    }

    bp, _ := client.NewBatchPoints(config)

    for i:= 1; i < len(srs.Observations); i++ {
        obs := srs.Observations[i]

        // build tags
        tags := map[string]string{"source": "fred"}

        // build fields
        fields := map[string]interface{}{
            "value":   obs.Value,
        }

        // TODO: parse the end time into a Timestamp
        timestamp := time.Parse(YYYY_MM_DD, obs.Date)

        point, err := client.NewPoint(
            name,
            tags,
            fields,
            obs.Date,
        )

        if err != nil {
            log.Fatal(err)
        }


        bp.AddPoint(point)
    }

    err = client.Write(bp)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Done")
}

func FetchFred() {
    con := model.Connect()

    // TODO: retrieve the data
}
