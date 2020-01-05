package internal

import (
    "log"
    "os"
    "fmt"

    client "github.com/influxdata/influxdb1-client/v2"
)

func Connect() client.Client {
    username := os.Getenv("DB_USER")
    password := os.Getenv("DB_PW")
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
