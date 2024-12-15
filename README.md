# microservices-hello-world Microservices Project

## Prerequisites
- Go 1.21+
- Protocol Buffers compiler
- gRPC

## Setup
1. Install dependencies:
```bash
go mod tidy
```

2. Generate Protobuf:
```bash
./generate-proto.sh
```

3. Run Services:
```bash
./run-services.sh
```

## Endpoints
- Service A: http://localhost:8080/service-a?name=Alice
- Service B: http://localhost:8080/service-b?name=Bob
