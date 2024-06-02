# Golang gRPC User Service with Search

This is a Golang gRPC service that provides functionalities for managing user details and includes a search capability.

## Objective

The main objectives of this service are:

- Simulate a database by maintaining a list of user details within a variable.
- Create gRPC endpoints for fetching user details based on a user ID and retrieving a list of user details based on a list of user IDs.
- Implement a search functionality to find user details based on specific criteria.

## Usage

### Prerequisites

- Golang should be installed on your machine.
- Protocol Buffers Compiler (protoc) should be installed.
- Install the required Go packages:

  ```bash
  go get -u google.golang.org/grpc
  go get -u github.com/golang/protobuf/protoc-gen-go


Clone this repository:
git clone <repository_url>

Run the server:
go run server/server.go

Run the client:
go run client/client.go
