# bizdash-db

InfluxDB - BizDash Database


## CLI Use

Run the instance with the specific config file.

`influxd -config influxdb.conf`

You must pass your credentials via the cli when accessing the service.

`influx -username <username> -password "<password>"`

^ Note that you need quotes around the password

## Basic Commands

- `SHOW DATABASES`
   - lists existing databases
- `USE <DB-NAME>`
   - specifies the database you wish to interact with

## Data

This is a time-series database.  Formatting the data is important for consumption.

Influx data consists of:

- `time`
    - The timestamp associated with the data point
- `measurement`
    - Like an SQL table, where the primary index is always `time`.
- `tags`
    - indexed values
- `fields`
   - non-indexed values

There can be millions of measurements and no up-front defined schema.

Null values simply aren't stored.

`tags`:
- `source`
   - defines the source from which the datapoint was taken
- `

`fields`:
- `unit`
   - defines the unit of measure for associated datapoints
   