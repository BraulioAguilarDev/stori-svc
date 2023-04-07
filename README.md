# Stori

Reading data from CSV files then it's sent the summary of transactions by email.

# About project
- We are using a modular application.
- Dependency-injection, more info [here](https://github.com/alecthomas/inject).
- We are using [GORM](https://gorm.io/docs/) library.

# Layout
![Drag Racing](doc/stori.png)
# Environment variables
Read environment variables with [viper](github.com/spf13/viper).

* `DSN_DB`    string  // Gorm source string. Example: `postgres://postgres:postgres@localhost:5432/dbname?sslmode=disable`
* `SG_KEY`    string  // Api key for Sendgrid platform
* `SG_SENDER` string  // Email addres as Sender identity. Default: `dev@braulio.tech` 
* `CSV_FILE`  string  // CSV file path

# Requirements

* [docker](https://www.docker.com/)
* [docker-compose](https://docs.docker.com/compose/)

# Test code

```shell
#  Build docker image "stori:latest"
$ make docker

# Create and run the containers
$ make dc-up

# Print logs
$ docker logs -f service_stori
```

# Evidence
![Drag Racing](doc/email-example.png)
