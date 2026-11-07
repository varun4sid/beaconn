# Beaconn

I built an alternate client API for my college hostel's student dashboard with automated quality of life features and more.

Click here to view the [frontend repository](https://github.com/varun4sid/beaconn-frontend).

## Prerequisites

- Go 1.27.1 or later
- Make (optional, for the provided commands)

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/varun4sid/beaconn-backend.git
   cd beaconn-backend
   ```

2. Download Go dependencies:

   ```bash
   go mod download
   ```

3. Optionally create a `.env` file:

   ```env
   PORT=8080
   APP_ENV=development
   ALLOWED_ORIGIN=http://localhost:3000
   CLIENT_DOMAIN=http://localhost:3000
   ```

   `PORT` defaults to `8080`. The other values default to empty strings, except
   `APP_ENV`, which defaults to `no_env`.

## Running

Start the development server:

```bash
make run
```

The server is available at `http://localhost:8080` by default.

To build and launch the compiled binary:

```bash
make launch
```

## Other commands

```bash
make fmt    # Format Go source files
make vet    # Run go vet
make test   # Run tests
make build  # Format, vet, test, and build the binary
```
