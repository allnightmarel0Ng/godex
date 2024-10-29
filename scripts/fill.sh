#!/bin/bash

curl -X POST http://localhost:8080/store \
     -H "Content-Type: application/json" \
     -d '{
           "link": "https://raw.githubusercontent.com/golang/go/master/src/net/http/http.go"
         }'
curl -X POST http://localhost:8080/store \
     -H "Content-Type: application/json" \
     -d '{
           "link": "https://raw.githubusercontent.com/golang/go/master/src/bufio/bufio.go"
         }'
curl -X POST http://localhost:8080/store \
     -H "Content-Type: application/json" \
     -d '{
           "link": "https://raw.githubusercontent.com/golang/go/master/src/compress/bzip2/bzip2.go"
         }'
curl -X POST http://localhost:8080/store \
     -H "Content-Type: application/json" \
     -d '{
           "link": "https://raw.githubusercontent.com/golang/go/master/src/compress/flate/deflate.go"
         }'
curl -X POST http://localhost:8080/store \
     -H "Content-Type: application/json" \
     -d '{
           "link": "https://raw.githubusercontent.com/golang/go/master/src/compress/lzw/writer.go"
         }'
curl -X POST http://localhost:8080/store \
     -H "Content-Type: application/json" \
     -d '{
           "link": "https://raw.githubusercontent.com/golang/go/master/src/compress/zlib/reader.go"
         }'
curl -X POST http://localhost:8080/store \
     -H "Content-Type: application/json" \
     -d '{
           "link": "https://raw.githubusercontent.com/golang/go/master/src/container/heap/heap.go"
         }'
curl -X POST http://localhost:8080/store \
     -H "Content-Type: application/json" \
     -d '{
           "link": "https://raw.githubusercontent.com/golang/go/master/src/fmt/print.go"
         }'
curl -X POST http://localhost:8080/store \
     -H "Content-Type: application/json" \
     -d '{
           "link": "https://raw.githubusercontent.com/golang/go/master/src/os/os.go"
         }'
curl -X POST http://localhost:8080/store \
     -H "Content-Type: application/json" \
     -d '{
           "link": "https://raw.githubusercontent.com/golang/go/master/src/io/io.go"
         }'
curl -X POST http://localhost:8080/store \
     -H "Content-Type: application/json" \
     -d '{
           "link": "https://raw.githubusercontent.com/golang/go/master/src/strings/strings.go"
         }'
curl -X POST http://localhost:8080/store \
     -H "Content-Type: application/json" \
     -d '{
           "link": "https://raw.githubusercontent.com/golang/go/master/src/time/time.go"
         }'
curl -X POST http://localhost:8080/store \
     -H "Content-Type: application/json" \
     -d '{
           "link": "https://raw.githubusercontent.com/golang/go/master/src/sync/mutex.go"
         }'
curl -X POST http://localhost:8080/store \
     -H "Content-Type: application/json" \
     -d '{
           "link": "https://raw.githubusercontent.com/golang/go/master/src/math/abs.go"
         }'
curl -X POST http://localhost:8080/store \
     -H "Content-Type: application/json" \
     -d '{
           "link": "https://raw.githubusercontent.com/golang/go/master/src/strconv/atoi.go"
         }'