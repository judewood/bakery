# bakery

Bakery Simulator

Simulates a bakery production line

Comprised of two applications

1. API Admin interface for CRUD operations on Products and recipes

2. Game app that will be the back end for a game where players can fulfil orders to gain points

## Build

1. Build api: `go build ./cmd/api`

2. Build game: `go build ./cmd/game`

## Run

1. Run api: `go build ./cmd/api`

Open http://localhost:8080/products in a browser to see available products

2. Build game: `go run ./cmd/game`

In terminal you will see a random order being created and fulfilled
