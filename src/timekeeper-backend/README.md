# Timekeeper Backend

Timekeeper Backend is a simple API built with Go, using the Gin framework and GORM for SQLite. It provides endpoints to manage version of remote entry for module federation. The API is documented using Swagger.

## Features

- **Health Check**: Check the status of the server.
- **Manage Remotes**: Add and retrieve remote entries.
- **Swagger Documentation**: Interactive API documentation.

## Project Structure

```md
timekeeper-backend/
 ├── cmd/ 
 │ └── app/ 
 │   └── main.go 
 ├── docs/ 
 │ ├── docs.go 
 │ ├── swagger.json 
 │ └── swagger.yaml 
 ├── internal/ 
 │ ├── db/ 
 │ │ └── db.go 
 │ ├── health/ 
 │ │ ├── handler_test.go 
 │ │ └── handler.go 
 │ ├── models/ 
 │ │ └── model.go 
 │ └── remote/ 
 │   ├── handler_test.go 
 │   └── handler.go 
 ├── .gitignore 
 ├── go.mod 
 ├── go.sum 
 ├── Makefile 
 └── README.md 
```

## Getting Started

### Prerequisites

- Go 1.22 or later
- SQLite

### Installation

1. Clone the repository:

    ```sh
    git clone https://github.com/timekeeperjs/timekeeper-backend.git
    cd timekeeper-backend
    ```

2. Install dependencies

    ```sh
    go mod tidy
    ```

3. Generate Swagger documentation:

    ```sh
    make swag
    ```

## Running the Application

You can run the application using the provided run.sh script or the Makefile.

```sh
make run
```

The application will be running at `http://localhost:8080.`

## Swagger Documentation

Swagger documentation is available at `http://localhost:8080/swagger/index.html.`

## Makefile Commands

- `make run`: Run the application.
- `make build`: Build the application.
- `make clean`: Clean build artifacts.
- `make test`: Run tests.
- `make swag`: Generate Swagger documentation.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.
