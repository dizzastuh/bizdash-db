package internal

import (
    "log"
    "os"

    client "github.com/influxdata/influxdb1-client/v2"
)

func Connect() client.Client {
    username := os.Getenv("INFLUX_USERNAME")
    password := os.Getenv("INFLUX_PASSWORD")
    dbhost   := os.Getenv("DB_HOST")

    conf := client.HTTPConfig {
        Addr:     dbhost,
        Username: username,
        Password: password,
    }

    con, err := client.NewHTTPClient(conf)

    if err != nil {
        log.Fatal(err)
    }

    return con
}
