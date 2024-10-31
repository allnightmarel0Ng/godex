# godex
[Hoogle](https://hoogle.haskell.org/)-like search engine for Go, literally Google

## Quickstart

### Configure

To configure the project create .env file with this variables:

| Variable Name           | Description                                                                 |
|-------------------------|-----------------------------------------------------------------------------|
| `GATEWAY_PORT`       | Integer value representing the port where HTTP requests should go.         |
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
GATEWAY_PORT=8080
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

**URL**: `/store`

**Method**: `POST`

**Content-Type**: `application/json`

**Request Body**:
  ```json
  {
    "link": "https://raw.githubusercontent.com/golang/go/master/src/sync/mutex.go"
  }
  ```
**Response**:
  - **200 OK**
  - **500 Internal Server Error**
  
    **Content-Type**: `application/json`
```json
{
  "code": 500,
  "message": "unexpected server error"
}
```
```json
{
  "code": 500,
  "message": "error reading request body"
}
```
  - **404 Not Found**

    **Content-Type**: `application/json`
```json
{
  "code": 404,
  "message": "unable to fetch the data from link: 404 Not Found"
}
```
  - **400 Bad Request**

    **Content-Type**: `application/json`
```json
{
  "code": 400,
  "message": "invalid link: not a .go file"  
}
``` 
```json
{
  "code": 400,
  "message": "invalid link: not in whitelist"
}
``` 
```json
{
  "code": 400,
  "message": "wrong content type: should be application/json"
}
```
```json
{
  "code": 400,
  "message": "error parsing JSON"
}
```
**Example**:
```bash
curl -X POST http://localhost:8080/store \
     -H "Content-Type: application/json" \
     -d '{
           "link": "https://raw.githubusercontent.com/golang/go/master/src/bufio/bufio.go"
         }'
```

### Find
Provide the signature to find all the functions with such signature in database. On success, responses with an array of functions

**URL**: `/find`

**Method**: `GET`

**Content-Type**: `application/json`

**Request Body**:
```json
{
  "signature": "(string)(int, error)"
}
```

**Response**:
  - **200 OK**

    **Content-Type**: `application/json`
```json
[
  {
      "name": "WriteString",
      "signature": "(string)(int,error)",
      "comment": "WriteString writes a string.\nIt returns the number of bytes written.\nIf the count is less than len(s), it also returns an error explaining\nwhy the write is short.\n",
      "file": {
          "name": "bufio.go",
          "package": {
              "name": "bufio",
              "link": "https://raw.githubusercontent.com/golang/go/master/src/bufio/bufio.go"
          }
      }
  },
  {
      "name": "WriteString",
      "signature": "(string)(int,error)",
      "comment": "WriteString implements [io.StringWriter] so that we can call [io.WriteString]\non a pp (through state), for efficiency.\n",
      "file": {
          "name": "print.go",
          "package": {
              "name": "fmt",
              "link": "https://raw.githubusercontent.com/golang/go/master/src/fmt/print.go"
          }
      }
  },
  {
      "name": "WriteString",
      "signature": "(string)(int,error)",
      "comment": "NoComment",
      "file": {
          "name": "io.go",
          "package": {
              "name": "io",
              "link": "https://raw.githubusercontent.com/golang/go/master/src/io/io.go"
          }
      }
  },
  {
      "name": "Atoi",
      "signature": "(string)(int,error)",
      "comment": "Atoi is equivalent to ParseInt(s, 10, 0), converted to type int.\n",
      "file": {
          "name": "atoi.go",
          "package": {
              "name": "strconv",
              "link": "https://raw.githubusercontent.com/golang/go/master/src/strconv/atoi.go"
          }
      }
  }
]
```
  - **500 Internal Server Error**

    **Content-Type**: `application/json`
```json
{
  "code": 500,
  "message": "error reading request body"
}
```
  - **400 Bad Request**

    **Content-Type**: `application/json`
```json
{
  "code": 400,
  "message": "wrong content type: should be application/json"
}
```
```json
{
  "code": 400,
  "message": "error parsing JSON"
}
```
  - **404 Not Found**

    **Content-Type**: `application/json`
```json
{
  "code": 404,
  "message": "error finding signature"
}
```

  
## Architecture

![](img/architecture.png)