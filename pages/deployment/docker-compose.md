#### docker-compose

This is the recommended method for local deployment and testing and is by far the easiest.

Prerequisites: `docker` and `docker-compose` are installed and on a relatively recent version.

##### Building and rebuilding with docker-compose
- `docker-compose up --build -d` to build the image from source and run flexo and the postgres database in the background. 
  - Flexo will be available on `localhost:8080`
  - The Postgresql DB is exposed on `localhost:5432`

##### Stopping
- `docker-compose stop`

##### Destroying the stack
- `docker-compose rm -f` to remove the containers

- `docker volume rm flexo_db-data` to remove DB data.

**NOTE**: Because the data volume persists after destroying the stack, rebuilding the stack without destroying the data volume means the stack will come back with the exact same data. This may or may not be what you want or intend when rebuilding. Additionally, running `faker` twice on an initialized and populated database _will_ fail due to unique field constraints for team IDs.

