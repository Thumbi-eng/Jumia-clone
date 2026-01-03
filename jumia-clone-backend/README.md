# Jumia Clone Backend

This repository contains the backend for the Jumia clone project, built using Go and gRPC microservices. The architecture is designed to be modular, with separate services handling different aspects of the application.

## Project Structure

- **services/**: Contains all microservices.
  - **user-service/**: Manages user-related operations.
    - **cmd/**: Entry point for the user service.
    - **internal/**: Contains handler, repository, and service logic.
    - **proto/**: gRPC definitions for the user service.
  - **product-service/**: Manages product-related operations.
    - **cmd/**: Entry point for the product service.
    - **internal/**: Contains handler, repository, and service logic.
    - **proto/**: gRPC definitions for the product service.
  - **order-service/**: Manages order-related operations.
    - **cmd/**: Entry point for the order service.
    - **internal/**: Contains handler, repository, and service logic.
    - **proto/**: gRPC definitions for the order service.
  - **cart-service/**: Manages cart-related operations.
    - **cmd/**: Entry point for the cart service.
    - **internal/**: Contains handler, repository, and service logic.
    - **proto/**: gRPC definitions for the cart service.

- **api-gateway/**: Acts as a single entry point for clients, routing requests to the appropriate services.
  - **cmd/**: Entry point for the API gateway.
  - **internal/**: Contains handler and middleware logic.

- **shared/**: Contains shared proto definitions and common logic used across services.

- **go.work**: Manages multiple modules in the workspace.

## Getting Started

1. **Clone the repository**:
   ```
   git clone https://github.com/Thumbi-eng/Jumia-clone.git
   cd jumia-clone-backend
   ```

2. **Install dependencies**:
   Each service has its own `go.mod` file. Navigate to each service directory and run:
   ```
   go mod tidy
   ```

3. **Define Proto Files**:
   Ensure that the proto files in each service's `proto` directory are defined correctly to specify the gRPC services and message types.

4. **Implement Service Logic**:
   Implement the business logic in the `internal/service` directories and the request handling in the `internal/handler` directories for each service.

5. **Run the Services**:
   Each service can be run independently by executing the `main.go` file in the `cmd` directory of each service.

6. **Set Up API Gateway**:
   Configure the API gateway to route requests to the appropriate services.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any enhancements or bug fixes.

