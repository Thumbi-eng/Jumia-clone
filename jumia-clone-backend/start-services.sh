#!/bin/bash

# Start all services in background
./services/user-service/bin/user-service &
./services/product-service/bin/product-service &
./services/order-service/bin/order-service &
./services/cart-service/bin/cart-service &
./api-gateway/bin/api-gateway &

echo "All services started!"
wait