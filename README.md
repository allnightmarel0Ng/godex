# godex
[Hoogle](https://hoogle.haskell.org/)-like search engine for Go, literally Google

## Quickstart

### Configure

To configure the project create .env file with this variables:

| Variable Name           | Description                                                                 |
|-------------------------|-----------------------------------------------------------------------------|
| `ENTRYPOINT_PORT`       | Integer value representing the port where HTTP requests should go.         |
| `PARSER_PORT`           | Port of the gRPC server implemented in the parser microservice.             |
| `CONTAINER_PORT`        | Port of the gRPC server implemented in the container microservice.          |
| `KAFKA_BROKER`          | Apache Kafka port.                                                          |
| `ZOOKEEPER_PORT`        | Zookeeper port.                                                             |
| `POSTGRES_NAME`         | Postgres database name.                                                     |
| `POSTGRES_USER`         | Postgres username.                                                          |
| `POSTGRES_PASSWORD`     | Postgres password for the user.                                             |
| `POSTGRES_PORT`         | Postgres port.                                                              |
| `POSTGRES_PORT_INCREMENTED` | Incremented version of the Postgres port for controlling the database outside of the image. |
| `WHITE_LIST`            | List of domains from which files are allowed to be downloaded.              |

.env example:
```
ENTRYPOINT_PORT=8080
PARSER_PORT=5000
CONTAINER_PORT=5001
KAFKA_BROKER=9092
ZOOKEEPER_PORT=2181
POSTGRES_NAME=godex
POSTGRES_USER=admin
POSTGRES_PASSWORD=admin
POSTGRES_PORT=5432
POSTGRES_PORT_INCREMENTED=5433
WHITE_LIST=["raw.githubusercontent.com"]
```

### Build

```shell
docker compose --env-file .env -f deployments/docker-compose.yml build
```

### Run

```shell
docker compose --env-file .env -f deployments/docker-compose.yml up -d
```

## API

### Store
Provide the link to raw .go file (hostname must be from `WHITE_LIST`) and parser will get all the functions for it to store in the database. Database is empty from the start, so you can store the functions you want to

- **URL**: `/store`
- **Method**: `POST`
- **Request Body**:
  ```json
  {
    "link": "https://raw.githubusercontent.com/golang/go/master/src/sync/mutex.go"
  }

### Find
Provide the signature to find all the functions with such signature in database. On success, responses with an array of functions

- **URL**: `/find`
- **Method**: `GET`
- **Request Body**:
  ```json
  {
    "signature": "(int)string"
  }
  
## Architecture

![](img/architecture.png)