# even-odd
![GitHub](https://img.shields.io/github/license/eisenwinter/even-odd)


A microservice returning even or odd numbers.

## Run in docker

```
docker run -e EVEN_ODD_PORT=8080 -p 8080:8080 docker.pkg.github.com/eisenwinter/even-odd/even-odd
```

Navigate to http://localhost:8080/even

## API

The service consists of 2 methods

### GET /even

Returns a json containing a even number.

```
{
"value": 163388078
}
```

### GET /odd

Returns a json containg a odd number.

```
{
"value": 3700939635
}
```