# Stori

Sending the transactions summary by email

# About project
- We're using a modular application.
- Dependency-injection, more info [here](https://github.com/alecthomas/inject).
- [GORM](https://gorm.io/docs/) ORM library for saving the data.
- [Goose](https://pressly.github.io/goose/) library to manage the migrations.

# Layout
![Drag Racing](doc/assets/layout.png)
# Environment variables
Read environment variables with [viper](github.com/spf13/viper).

* `DSN_DB`    string  // Gorm source string. Example: `postgres://postgres:postgres@localhost:5432/dbname?sslmode=disable`
* `SG_KEY`    string  // Api key for Sendgrid platform
* `SG_SENDER` string  // Email address as Sender identity.
* `CSV_FILE`  string  // CSV file path

Note: `SG_KEY` and `SG_SENDER`  are available only for this challenge.

# Requirements

* [docker](https://www.docker.com/)
* [docker-compose](https://docs.docker.com/compose/)

# Run application using Docker

```shell
# Clone repo
$ git@github.com:BraulioAguilarDev/stori-svc.git

# Go to stori directory
$ cd stori

# Building docker image "stori:latest" with following command
$ make docker

# Create and run the containers
$ make dc-up

# Print logs
$ docker logs -f app.stori.dev
```

# Evidence

Email info

![Drag Racing](doc/assets/email-example.png)

Logs info

![Drag Racing](doc/assets/logs.png)