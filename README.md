# Kinfool

Worst-possible KISS RESTful API meta-framework on Gin, Golang.

Beware that this project is still in heavy-development, using as production is not recommended at all.

## Features

- Based on Gin
- Delivered as ready-to-develop template
- Kinfool's ``kn`` cli command for creating routes, controllers, services and custom middlewares

## Project Structure

- ``/kinfool.go``: Main Application
- ``/internal/controllers/*.go``: Controllers
- ``/internal/middlewares/*.go``: Middlewares
- ``/internal/routes/*.{http_method}.go``: Routes
- ``/internal/services/*.go``: Services
- ``/internal/utilities/*.go``: Utilities

## Quickstart

``git clone https://github.com/kinnnine/kinfool`` or fork via github

To using ``kn``, you need to ``go build .`` inside ``kinfool/kn`` directory then export into System's ``PATH``.

To build ``kinfool``, just use ``go build .`` or ``go build . -o appname`` for custom executable name.

## Environment Variables (dotenv)

```
### Kinfool configuration (required settings)
LISTEN=0.0.0.0
PORT=8080
TRUSTED_PROXIES=0.0.0.0

### Below here, you can customize and add env according to your application needs

DB_HOST=localhost
DB_PORT=5432
and so on ...
```