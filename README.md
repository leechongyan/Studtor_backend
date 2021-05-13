# Studtor_backend

To instantiate the database:
```
sqlite3 ./build/studtor.db
# then, run database scripts under scripts
```

To build the program:
```
go build -o ./build/main ./cmd/main.go
```

To run the program:
```
./build/main
```

## Routes
Web server exposes three routes (for testing purposes):
- GET : localhost:8080/ping - example GET request endpoint
- GET : localhost:8080/students/:id - endpoint to retrieve student from db by ID
- GET : localhost:8080/students/create/:name - endpoint to insert student with name into db
