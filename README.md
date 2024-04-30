# Go TCP Game Server
A TCP server that can handle multiple clients, and facilitate commands over the wire as an array of bytes in Big Endian format.

---
## Getting Started
TODO: Add instructions for releases/download from source

### Prerequisites

### Building
To build from source use:
```bash
go build -o go-tcp-game-server
```

### Running tests
To run all the tests use:
```bash
go test ./...
```

To run a specific test use:
```bash
# See go help testflag for detailed testing flags, this will run the tests in verbose mode and only once
go test -run testName folderPath --verbose -count 1
```

### Running the server
To run the server use:
```bash
go run main.go
```

To run the server from the binary use:
```bash
./go-tcp-game-server -ip 127.0.0.1 -port 8006
```
