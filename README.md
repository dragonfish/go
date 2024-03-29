# dragonfish/go

This repository is a collection of reusable packages written in Go. Each package serves a unique purpose and can be used independently or together to build robust applications.

## Packages

### Connect Interceptors

Located in `pkg/connect-interceptors`, this package provides interceptors for handling panic situations in your Go applications built with [connect](https://connect.build).

### Connect Middlewares

Located in `pkg/connect-interceptors`, this package provides middlewares for handling Firebase authentication in your Go applications built with [connect](https://connect.build).

### Logger

Located in `pkg/logger`, this package provides a logging interface with implementations for different logging libraries. It currently supports ZapLogger and LogrusLogger.

## Building the Project

This project uses Bazel as its build system. To build the project, you need to have Bazel installed on your machine. Once installed, you can build the project by running the following command in the root directory of the project:

`bazel build //...`
