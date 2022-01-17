# Flexo

![BING BONG](https://img.shields.io/badge/BING-BONG-red?style=plastic&logo=appveyor)
![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/SECCDC/flexo)
![Docker build status](https://img.shields.io/github/workflow/status/SECCDC/flexo/Deploy%20Docker?label=Docker%20build&style=plastic)
![License](https://img.shields.io/badge/License-MIT-brightgreen?style=plastic)
![GitHub contributors](https://img.shields.io/github/contributors/SECCDC/flexo?style=plastic)

## Quickstart

### Running with docker-compose
`SECRET=zbeul DB_USER=flexo docker-compose up`

> All 3 config variables default to `flexo`, but DB_USER must be specified in the command so that the health check can execute successfully.
> `SECRET` is the secret shared between the frontend and the backend. It defaults to "shared_secret"

> DBSSL sets the connection to the database's ssl mode options. It is set to `disable` by default.

## Testing

### Resetting the database
If you want a clean database, you have to stop the docker compose stack, remove the docker compose stack, then delete the docker volume with the DB data before starting the stack again.

```
➜  flexo  git:(main) ✗ docker volume list
DRIVER              VOLUME NAME
local               flexo_db-data

➜  flexo  git:(main) ✗ docker-compose rm -f
Going to remove flexo_flexo_1, flexo_db_1
Removing flexo_flexo_1 ... done
Removing flexo_db_1    ... done

➜  flexo  git:(main) ✗ docker volume rm flexo_db-data
flexo_db-data
```

### Building fresh code
`docker-compose up --build -d`

### Authentication and Authorization
The shared secret is used like a very basic JWT.
`Authorization` header must have a value of `Bearer $secret`.

`http -v --auth-type=jwt --auth="test" "localhost:8080/report/team/1"`

You need the [httpie-jwt-auth plugin](https://github.com/teracyhq/httpie-jwt-auth) to run this command.

### Adding mock data
From the `faker` directory, `go run ./main.go`

### Sending an event
`http --json post http://localhost:8080/event targets:='[1,2,3]' teams:='[1,2,3]' category:=1 description="test event"`

### Get teams
`http localhost:8080/teams`

### Get categories
`http localhost:8080/categories`

### Get events
`http localhost:8080/events`

### Get targets
`http localhost:8080/targets`

### Getting a team report
`http localhost:8080/report/team/$ID` where $ID is the team's ID.

### Getting all teams report
`http localhost:8080/report/teams` return a json list of each team with their associated report
