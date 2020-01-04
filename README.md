# goapi.g8

Giter8 template to develop `API server / Microservices` in Go using Echo framework.

Latest release: [template-v0.4.r4](RELEASE-NOTES.md).

## Features

- Create new project from template with [go-giter8](https://github.com/btnguyen2k/go-giter8).
- API framework:
  - JSON-encoded.
  - Plugable filters, included built-in ones:
    - `AddPerfInfoFilter`
    - `LoggingFilter`
    - `AuthenticationFilter`
  - REST API gateway built on top [Echo Framework v4](https://echo.labstack.com).
  - gRPC API gateway (since [template-v0.4.r2](RELEASE-NOTES.md)).
- Samples:
  - [Trivial APIs](src/main/g8/src/samples/bootstrap.go)
  - [CRUD APIs with MongoDB](src/main/g8/src/samples_crud_mongodb/bootstrap.go)
  - [CRUD APIs with PostgreSQL](src/main/g8/src/samples_crud_pgsql/bootstrap.go)
  - [Filters](src/main/g8/src/samples_api_filters/bootstrap.go)
- Sample `.gitlab-ci.yaml` & `Dockerfile` to package application as Docker image
- Scaffolding (since [template-v0.4.r2](RELEASE-NOTES.md), require [go-giter8](https://github.com/btnguyen2k/go-giter8) `v0.4.0` or higher).

## Getting Started

### Install `go-giter8`

This a Giter8 template, so it is meant to be used in conjunction with a giter8 tool.
Since this is a template for Go application, it make sense to use [go-giter8](https://github.com/btnguyen2k/go-giter8).

See [go-giter8](https://github.com/btnguyen2k/go-giter8) website for installation guide.

### Create new project from template

```shell script
g8 new btnguyen2k/goapi.g8
```

and follow the instructions.

> Note: This template requires `go-giter8` version `0.3.2` or higher.

### Application configurations

Application configurations are loaded from `config/application.conf` file, in [HOCON format](https://github.com/lightbend/config/blob/master/HOCON.md).
Its content is human-readable and self-explained. So, this readme lists only key configurations:

**application info**
```hocon
app {
  # this section configures application's info such as name, version, description, etc.
}
```

**API configurations**
```hocon
api {
  # "api" section configure common API settings swuch as max request size or request timeout

  http {
    # this sub-section configures HTTP/Rest API gateway
  }

  grpc {
    # this sub-section configures gRPC API gateway
  }
}
```

**API endpoints**
```hocon
api.http.endpoints {
  # this sub-section defines API HTTP endpoints
}
```

### Implement APIs

**General guideline**

- Each API is an `itineris.IApiHandler` instance (e.g. `func(*ApiContext, *ApiAuth, *ApiParams) *ApiResult`)
- Register APIs with the global API router instance `goapi.ApiRouter`
- (Optional) Define API HTTP endpoints in configuration `api.http.endpoints` sub-section.

> If gRPC API gateway is enable, APIs are automatically available via gRPC.
> See [gRPC service definition file](src/main/g8/grpc/api_service.proto).

Sample APIs:
- [Trivial APIs](src/main/g8/src/samples/bootstrap.go)
- [CRUD APIs with MongoDB](src/main/g8/src/samples_crud_mongodb/bootstrap.go)
- [CRUD APIs with PostgreSQL](src/main/g8/src/samples_crud_pgsql/bootstrap.go)

**API filter**

API filter (instance of type `itineris.IApiFilter`) is plugable component that is used to intercept API call and do some pre-processing, intercept result and do some post-processing before returning to caller.

See [sample API filter](src/main/g8/src/samples_api_filters/bootstrap.go).

### Scaffolding

At project's root directory, run command

```shell script
$ g8 scaffold <scaffold-name>
```

to generate scaffolds. This template includes a few [scaffolds](src/main/scaffolds/).

> Scaffolding requires [go-giter8](https://github.com/btnguyen2k/go-giter8) `v0.4.0` or higher.


## LICENSE & COPYRIGHT

See [LICENSE.md](LICENSE.md) for details.

## Giter8 template

For information on giter8 templates, please see http://www.foundweekends.org/giter8/
