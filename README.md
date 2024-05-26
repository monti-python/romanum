# Introduction
Romanum is a ready-to-run REST API to convert a range of numbers to Roman numerals.

# Run the app
## Linux
1. Make sure Go 1.22 is available in your PATH
```bash
go version
# go version go1.22.3 linux/amd64
```
2. Build the app 
```bash
go test -v ./...
go build -o romanum ./cmd/server
```
3. Run it
```bash
./romanum
# Server is running on port 8080...
```
## Docker
Romanum can also be run as a container

```bash
docker build -t romanum .
docker run --rm -p 8080:8080 romanum
```

# Progress
- [x] Ensure the client can submit a range of numbers (e.g., 10-50).
- [x] Handle numbers within the range of 1 to 3999; any numbers outside this range should return an error message.
- [x] The server must accurately convert all numbers in the specified range to Roman numerals.
- [x] Collect and return all results in ascending order.
- [ ] Provide the OpenAPI specification.
- [x] Wrap the service into a Docker container.
- [x] Create an integration test suite to automate the evaluation of the functionality.

# Improvements
- Logging
- Choosing port
- Custom verbosity
- Provide thoughts on what improvements could be made if you had more time. Also, provide information about what is missing to have this API production ready.
