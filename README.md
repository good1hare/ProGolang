# ProGolang

## Requirements:
* docker >= 17.12.0+
* docker-compose

## Installation
Need start Postgres, run this command `docker-compose up -d`


Need to create config.yaml in folder configs with data

And start the server.

```sh
go run main.go
```

## Access to postgres:
* `localhost:5432`
* **Username:** postgres (as a default)
* **Password:** changeme (as a default)

## Access to PgAdmin:
* **URL:** `http://localhost:5050`
* **Username:** pgadmin4@pgadmin.org (as a default)
* **Password:** admin (as a default)

## Add a new server in PgAdmin:
* **Host name/address** `postgres`
* **Port** `5432`
* **Username** as `POSTGRES_USER`, by default: `postgres`
* **Password** as `POSTGRES_PASSWORD`, by default `changeme`

## Logging

There are no easy way to configure pgadmin log verbosity, and it can be overwhelming at times. It is possible to disable pgadmin logging on the container level.

Add the following to `pgadmin` service in the `docker-compose.yml`:

```
logging:
  driver: "none"
```