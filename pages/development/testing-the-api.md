## Testing the API

### Tools

We strongly prefer using [`httpie`]() for API testing but you can use whatever you want. Instructions for API calls in this doc will be written for HTTPie.

Additionally, it is recommended you use the [JWT Auth plugin](https://github.com/teracyhq/httpie-jwt-auth) for HTTPie

### Authentication and Authorization

The shared secret is used like a very basic JWT.
`Authorization` header must have a value of `Bearer $secret`.

Example:
```
http -v --auth-type=jwt --auth="test" "localhost:8080/report/team/1"
```

You need the [httpie-jwt-auth plugin](https://github.com/teracyhq/httpie-jwt-auth) to run this command.

### Basic tasks

#### POST event
`http --json post http://localhost:8080/event targets:='[1,2,3]' teams:='[1,2,3]' category:=1 description="test event"`

#### GET teams
`http localhost:8080/teams`

#### GET categories
`http localhost:8080/categories`

#### GET events
`http localhost:8080/events`

#### GET targets
`http localhost:8080/targets`

#### GET single team report
`http localhost:8080/report/team/$ID` where $ID is the team's ID.

#### GET all teams reports
`http localhost:8080/report/teams` return list of reports, one per team
