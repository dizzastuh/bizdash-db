package internal

import (
    "log"
    "net/url"
    "os"

    client "github.com/influxdata/influxdb1-client"
)

func Connect() *client {
    username := os.Getenv("INFLUX_USERNAME")
    password := os.Getenv("INFLUX_PASSWORD")
    dbURL    := os.Getenv("DB_HOST")

    host, err := url.Parse(dbURL)
    if err != nil {
        log.Fatal(err)
    }

    conf := client.Config{
        URL:      *host,
        Username: username,
        Password: password,
    }

    con, err := client.NewHttpClient(conf)

    if err != nil {
        log.Fatal(err)
    }

    return con
}
