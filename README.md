# Hermes

## Quickstart

### Prepare the Database
1. Start a database: `docker run -d --name flexo-db -it -e MYSQL_ROOT_PASSWORD=<password> -p 127.0.0.1:3306:3306 mysql:5.7.14`
2. `pip install -r requirements.txt && python3 ./seed-db.py` (or use a venv if you're fancy)
3. Run migrations: `go run ./main.go migrate --dbPass <password>`

### Start the server
`go run ./main.go run --dbPass <password>`

### Running with docker-compose
`docker-compose up`

DB password should be changed (in the docker-compose.yml file) before running this.

## Testing
Hermes only has 3 routes:
- GET /healthz
- GET /api/v1/products
- POST /api/v1/order

To send an order using [httpie](https://httpie.io/):
```
http --json post http://localhost:8080/api/v1/order IDs:='[1,2,3]'
```
