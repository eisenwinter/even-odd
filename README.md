# even-odd
![GitHub](https://img.shields.io/github/license/eisenwinter/even-odd)


A microservice returning even or odd numbers.

## Run in docker

```
docker run -p 8080:8080 -p 8081:8081 docker.pkg.github.com/eisenwinter/even-odd/even-odd
```

Navigate to http://localhost:8080/even

## REST API

The service consists of 2 methods

### GET /rest/even

Returns a json containing a even number.

```
{
"value": 163388078
}
```

### GET /rest/odd

Returns a json containg a odd number.

```
{
"value": 3700939635
}
```


### Swagger

```
basePath: /
definitions:
  main.numberResponse:
    properties:
      value:
        type: integer
    type: object
host: localhost:8080
info:
  contact: {}
  description: This api supplied even or odd numbers.
  license:
    name: MIT
    url: https://github.com/eisenwinter/even-odd/blob/master/LICENSE
  title: even-odd API
  version: "1.0"
paths:
  /rest/even:
    get:
      consumes:
      - application/json
      description: Returns a even number
      operationId: even
      produces:
      - application/json
      responses:
        "200":
          description: OK
          schema:
            $ref: '#/definitions/main.numberResponse'
      summary: Returns a even number
  /rest/odd:
    get:
      consumes:
      - application/json
      description: Returns a odd number
      operationId: odd
      produces:
      - application/json
      responses:
        "200":
          description: OK
          schema:
            $ref: '#/definitions/main.numberResponse'
      summary: Returns a odd number
swagger: "2.0"
```

## GRPC API

Default port is 8081

### proto File

```
syntax = "proto3";
import "google/protobuf/empty.proto";

message NumberResponse {
    int64 value = 1;
  }
  

service EvenOddService {
    rpc Even(google.protobuf.Empty) returns (NumberResponse);
    rpc Odd(google.protobuf.Empty)  returns (NumberResponse);
}
```