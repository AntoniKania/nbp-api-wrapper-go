# NBP API Wrapper (golang edition)

This repository contains a Go version of an [existing project](https://github.com/AntoniKania/spring-course-exam) originally implemented as a part of university's Spring course that I attended. The purpose of this project was to explore and experiment with new technologies, while replicating the functionality of the original project.

## Features
- Exposes an HTTP endpoint to calculate the average currency rate from a specified time window.
- Utilizes the Gin framework for building the HTTP server.
- Saves request details, including timestamps and parameters, to a PostgreSQL database.
- Sends and process requests to one of [NBP Web API](http://api.nbp.pl/) endpoints.
- Unit tests implemented using [WireMock](https://hub.docker.com/r/wiremock/wiremock) (along with [Go WireMock Client](https://github.com/walkerus/go-wiremock)) and [Testcontainers](https://golang.testcontainers.org/) for isolated testing of external server behavior.
- Docker Compose configuration included for setting up the PostgreSQL database.

## Usage

### Database
The PostgreSQL database is managed using Docker Compose. To start the database container, run:
```
docker compose up -d
```

### Endpoint
To start the server and expose the endpoint, run the following command (database needs to be started first):
```
go run ./cmd/nbp-api-wrapper-go/main.go
```
The server will start listening on a specified port (default: 8080). You can access the endpoint at http://localhost:8080/nbp/{currency_code}/{start_date}/{end_date}.

### Unit Tests
Run the tests using:
```
go test ./...
```