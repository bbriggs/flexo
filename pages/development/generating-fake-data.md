## Using Faker to Generate Fake Data

Once your instance has been deployed, you can use our `faker` utility to generate fake teams, categories, and events for testing.

### Running faker

**NOTE**: Faker is for local development only and assumes both your API and database are exposed on localhost.

Assuming a working Golang environment and the stack is deployed, preferrably by `docker-compose`, run `go run ./faker/main.go` from the root of the repo. If it runs without error, you will now have a randomly generated set of teams, categories, and events for testing your API.

**NOTE**: Because of the uniqueness constraints of team IDs, you cannot run faker twice on the same database without clearing the data first.
