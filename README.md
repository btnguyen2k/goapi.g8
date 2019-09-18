# go_echo-microservices-seed.g8

Giter8 template to develop microservices in Go using Echo framework.

Latest release: [template-v0.4.r1](RELEASE-NOTES.md).

## Features

- Create new project from template with [go-giter8](https://github.com/btnguyen2k/go-giter8).
- API framework:
  - Built on top [Echo Framework v4](https://echo.labstack.com).
  - JSON-encoded (request/response) RESTful APIs.
  - Plugable filters, included built-in ones:
    - `AddPerfInfoFilter`
    - `LoggingFilter`
    - `AuthenticationFilter`
- Samples:
  - [Trivial APIs](src/main/g8/src/samples)
  - [CRUD APIs with MongoDB](src/main/g8/src/samples_crud_mongodb)
  - [CRUD APIs with PostgreSQL](src/main/g8/src/samples_crud_pgsql)
  - [Filters](src/main/g8/src/samples_api_filters)
- Sample `.gitlab-ci.yaml` & `Dockerfile` to package application as Docker image

## Getting Started

### Install `go-giter8`

This a Giter8 template, so it is meant to be used in conjunction with a giter8 tool.
Since this is a template for Go application, it make sense to use [go-giter8](https://github.com/btnguyen2k/go-giter8) command line tool.

Installing `go-giter8` is as simple as:

```
go get github.com/btnguyen2k/go-giter8/g8
```

### Create new project from template

```
g8 new btnguyen2k/go_echo-microservices-seed.g8
```

and follow the instructions.


## LICENSE & COPYRIGHT

See [LICENSE.md](LICENSE.md) for details.

## Giter8 template

For information on giter8 templates, please see http://www.foundweekends.org/giter8/
