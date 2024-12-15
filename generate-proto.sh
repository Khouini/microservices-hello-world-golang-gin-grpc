#!/bin/bash

# Generate protobuf for service-a
cd service-a
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    pb/service_a.proto
cd ..

# Generate protobuf for service-b
cd service-b
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    pb/service_b.proto
cd ..