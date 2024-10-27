# gRPC Restaurant Service (Go)

This project implements a gRPC service for managing restaurants and their menu items in Go.
It supports operations like adding a restaurant, retrieving all restaurants,
finding a restaurant by name, and getting menu items by category.

## Project Structure

``` css
Go instance lab2
│
├── proto
│   └── restaurant.proto
├── restaurant-service
│   └── proto
│       ├── restaurant.pb.go
│       └── restaurant_grpc.pb.go
```

## Overview

This Go project provides a gRPC service for restaurant management. The service allows clients to:

- **AddRestaurant**: Adds a new restaurant.
- **GetAllRestaurants**: Retrieves all restaurants.
- **GetRestaurantByName**: Finds a restaurant by its name.
- **GetMenuItemsByCategory**: Retrieves menu items by category.

### Prerequisites

- Go 1.19+
- Protocol Buffers Compiler (`protoc`)
- gRPC Go libraries

### Setup Instructions

1. **Clone the repository**:

   ```bash
   git clone https://github.com/kheder-hassoun/grpc_homwork_Go_Lang.git
   ```

2. **Install Go dependencies**:

   Run the following command to install necessary gRPC and Protobuf libraries:

   ```bash
   go mod tidy
   ```

3. **Generate Go code from `.proto`**:

   Use `protoc` to generate the Go files from the `.proto` files. Ensure that you have the Go `protoc` plugins installed (`protoc-gen-go` and `protoc-gen-go-grpc`).

   ```bash
   protoc --go_out=. --go_opt=paths=source_relative \
          --go-grpc_out=. --go-grpc_opt=paths=source_relative \
          proto/restaurant.proto
   ```

   This will generate two Go files:

   - `restaurant.pb.go`: Contains the message definitions for the service.
   - `restaurant_grpc.pb.go`: Contains the gRPC server and client code.

4. **Build the project**:

   To build the Go project, run:

   ```bash
   go build
   ```

### Running the Server

1. **Start the gRPC server**:

   Run the server by executing the following command:

   ```bash
   go run restaurant-service/server.go
   ```

### Project Components

- **proto/**: This directory contains the `restaurant.proto` file, which defines the gRPC service and the message formats.
- **restaurant-service/proto/**: Contains the generated Go files for Protobuf and gRPC:
  - `restaurant.pb.go`: The generated file that contains Protobuf message definitions.
  - `restaurant_grpc.pb.go`: The generated file that contains the gRPC service and client code.
