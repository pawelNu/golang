# Project webserwis-net-http

-   [Project webserwis-net-http](#project-webserwis-net-http)
    -   [Getting Started](#getting-started)
    -   [MakeFile](#makefile)
    -   [REST API docs](#rest-api-docs)
    -   [Issues](#issues)

One Paragraph of project description goes here

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

## MakeFile

Run build make command with tests

```bash
make all
```

Build the application

```bash
make build
```

Run the application

```bash
make run
```

Live reload the application:

```bash
make watch
```

Run the test suite:

```bash
make test
```

Clean up binary from the last build:

```bash
make clean
```

## REST API docs

1. https://github.com/swaggo/swag
2. https://github.com/swaggo/http-swagger

## Issues

1. unknown field RightDelim in struct literal of type "github.com/swaggo/swag".Spec

    Solution

    ```bash
    go get -u github.com/swaggo/swag
    ```
