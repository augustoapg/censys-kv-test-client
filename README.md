# Censys KV Store Test Client

A test client service built with Go for testing the Censys KV Store service.

This service provides the following REST API endpoints to test the key-value store functionality:

- `POST /test_deletion` - Tests key-value pair deletion functionality
- `POST /test_overwrite` - Tests key-value pair overwrite functionality

The service is designed to verify that the KV store service works correctly by performing end-to-end tests that involve creating, retrieving, updating, and deleting key-value pairs.

## Testing Instructions

### Prerequisites

- [Docker](https://docs.docker.com/engine/install/)
- [Docker Compose](https://docs.docker.com/compose/install/)
- A running KV store service (either on host or in container)

### Running this service directly (outside Docker)

1. Run the service:
```bash
# NOTE: The default port is 8081. You can change it by setting the PORT environment variable.
# The default KV store URL is http://localhost:8080. You can change it by setting the KV_STORE_URL environment variable.
go mod tidy
go run main.go # Service will be available at http://localhost:8081 (or the port you set)
```

2. Test endpoints:
```bash
# Test deletion functionality
curl -X POST http://localhost:8081/test_deletion

# Test overwrite functionality
curl -X POST http://localhost:8081/test_overwrite
```

### Running in Docker

Since this and the KV store service are meant to run on containers, they need a docker network to communicate with each other.

0. Create Docker network, clone and run the KV store service in Docker:
```bash
docker network create censys-network
git clone git@github.com:augustoapg/censys-kv-store.git
cd censys-kv-store
docker-compose up -d
```

1. Run this test client service in Docker:
```bash
docker-compose up -d  # Test client will be available at http://localhost:8081
```

2. Call test client endpoints:
```bash
# Verifies that deleting a key-value pair works
curl -X POST http://localhost:8081/test_deletion

# Verifies that overwriting a key-value pair works
curl -X POST http://localhost:8081/test_overwrite
```

## Expected Responses

### Successful Deletion Test
```json
{
  "message": "Deletion test successful"
}
```

### Successful Overwrite Test
```json
{
  "message": "Overwrite test successful"
}
```

## Tearing down

To stop the services, run:
```bash
docker-compose down
```

To remove the network, run:
```bash
docker network rm censys-network
```
