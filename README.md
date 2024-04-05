# go-template
This repository contains the code for our fully featured go-template, from which we can start new Go projects.

## Features
- **Logging**: Structured logging with [slog](https://pkg.go.dev/golang.org/x/exp/slog)
- **Configuration**: Configuration management with [koanf](https://pkg.go.dev/github.com/knadh/koanf)
- **HTTP Server**: HTTP server with [echo](https://pkg.go.dev/github.com/labstack/echo/v4)
- **Database**: DB ORM with [ent](https://entgo.io)

## Requirements
- [Go >= 1.21.5](https://go.dev/doc/install)
- [swagger-cli](https://www.npmjs.com/package/swagger-cli)
- [ent](https://entgo.io/docs/getting-started)
- [ogen](https://ogen.dev/docs/intro)
- [wire](https://github.com/google/wire)

## Getting Started
1. Clone the repository
2. Install the dependencies
3. Run the application

## Codegen
A detailed explanation of the codegen process can be found [here](docs/en/codegen.md).

## Devcontainers
This repository contains a devcontainer configuration for Visual Studio Code. The devcontainer is pre-configured with all the necessary tools and extensions to start developing right away. A detailed explanation of the devcontainer can be found [here](docs/en/devcontainer.md).