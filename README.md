# bakery

Bakery Simulator

Simulates a bakery production line

Comprised of two applications

1. API Admin interface for CRUD operations on Products and recipes

2. Game app that will be the back end for a game where players can fulfil orders to gain points

## Build

1. Build api: `go build ./cmd/api`

2. Build game: `go build ./cmd/game`

## Running the code locally

1. Run api: `go build ./cmd/api`

If you install VScode extension REST Client by Huachao Mao then you can run http requests against the code by clicking on 'Send Request' in a \*.http file inside the api-requests folder

Otherwise open the url (eg http://localhost:8080/products) in a browser

2. Build game: `go run ./cmd/game`

In terminal you will see a random order being created and fulfilled

## Unit tests

Run `go test ./...` for short output
Running `go test ./... -v` will show you more detail of passing tests

### Run a single test

Navigate to the folder containing the test file and run below , replacing TestName 
`go test -run TestName -v`

eg 
```
cd internal\products
go test -run TestProductControllerGetAll -v
```

### Run a single sub-test
Navigate to the folder containing the test file and run below, replacing the TestName and SubName
`go test -run TestName/SubName`

eg 
```
cd internal\products
go test -run TestProductControllerGetAll/zero -v
```
### Test Coverage

Create an output file with command below:
```
go test -coverprofile c.out ./...
```
then **open bash terminal** (don't use VSCode terminal in Windows) and enter command
```
go tool cover -html=c.out  
```
it should open a tab in your browser and you can use the drop down to  see color coded coverage for each file

## Tools

### Postman

The json file in `tools/postman` can be imported into Postman and contains sample client requests
