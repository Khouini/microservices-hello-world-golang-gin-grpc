#!/bin/bash

# Start Service A
gnome-terminal -- bash -c "cd service-a && go run main.go; bash"

# Start Service B 
gnome-terminal -- bash -c "cd service-b && go run main.go; bash"

# Start API Gateway
gnome-terminal -- bash -c "cd api-gateway && go run main.go; bash"
