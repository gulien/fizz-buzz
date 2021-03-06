# fizz-buzz REST server

[![Go Reference](https://pkg.go.dev/badge/github.com/gulien/fizz-buzz.svg)](https://pkg.go.dev/github.com/gulien/fizz-buzz)
[![Tests](https://github.com/gulien/fizz-buzz/actions/workflows/tests.yml/badge.svg)](https://github.com/gulien/fizz-buzz/actions/workflows/tests.yml)
[![Lint](https://github.com/gulien/fizz-buzz/actions/workflows/lint.yml/badge.svg)](https://github.com/gulien/fizz-buzz/actions/workflows/lint.yml)
[![codecov](https://codecov.io/gh/gulien/fizz-buzz/branch/master/graph/badge.svg?token=60U5BV3JM8)](https://codecov.io/gh/gulien/fizz-buzz)
[![Go Report Card](https://goreportcard.com/badge/github.com/gulien/fizz-buzz)](https://goreportcard.com/report/github.com/gulien/fizz-buzz)

A simple fizz-buzz REST server with statistics.

```
Usage of fizzbuzz:
      --port int      Set the port on which the fizz-buzz server should listen (default 80)
      --timeout int   Set the maximum duration in seconds before timing out execution of fizz-buzz (default 30)
```

## Endpoints
 
### GET `/`

Basic health check, always returns `200` if the server is running.

### GET `/api/v1/fizz-buzz`

This endpoint returns a JSON list of strings with numbers from 1 to `limit`, where: 
all multiples of `int1` are replaced by `str1`, all multiples of `int2` are replaced
by `str2`, all multiples of `int1` and `int2` are replaced by `str1str2`.

It accepts the following query parameters:

* `int1` - required, non-zero integer
* `int2` - required, non-zero integer
* `limit` - required, non-zero positive integer
* `str1` - a string
* `str2` - a string

<details>
    <summary>Example (200 OK)</summary>

`/api/v1/fizz-buzz?int1=2&int2=3&limit=10&str1=foo&str2=bar`

```json
["1","foo","bar","foo","5","foobar","7","foo","bar","foo"]
```
</details>

<details>
    <summary>Example (400 Bad Request)</summary>

`/api/v1/fizz-buzz?int1=2&int2=0&limit=10&str1=foo&str2=bar`

```json
{"message":"zero int1 and/or int2"}
```
</details>

<details>
    <summary>Example (503 Service Unavailable)</summary>

```json
{"message":"context deadline exceeded"}
```
</details>

### GET `/api/v1/stats`

This endpoint returns a JSON object with the parameters of the most frequent request and
the number of occurrences of this request.

<details>
    <summary>Example (200 OK)</summary>

`/api/v1/stats`

```json
{"count":10,"int1":"2","int2":"3","limit":"10","str1":"foo","str2":"bar"}
```
</details>
 
???? Current implementation of this endpoint relies on an in memory data source.
In other words, the statistics are not persisted between runs nor are they relevant
in a distributed environment.

However, one may provide its own data source implementation by implementing the
`stats.Statitistics` interface. Suitable data sources could be either No-SQL or SQL.
For the latter, the implementer will have to make sure the requests' parameters do not
lead to SQL injection (e.g., `str1` equals `DROP TABLE foo;`).

## Development

**Requirements**

* Go >= 1.16
* [golangci-lint](https://golangci-lint.run/) >= 1.39
* A linux-like terminal (ideally)

**Makefile commands**

* `make fmt` - Shortcut for `go fmt` and `go mod tidy`
* `make lint` - Runs linters
* `make tests` - Runs tests
* `make todos` - Shows TODOs
* `make godoc` - Runs a local webserver for godoc
* `make run` - Runs the application (`PORT` and `TIMEOUT` are available as variables, i.e, `make run PORT=80 TIMEOUT=30`)
* `make build` - Builds the application (`VERSION` is available as variable, i.e, `make build VERSION=foo`)