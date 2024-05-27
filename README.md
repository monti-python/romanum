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

4. Make sure the service is up
```bash
curl "localhost:8080/convert?start=3&end=5"
# [{"number":3,"roman":"III"},{"number":4,"roman":"IV"},{"number":5,"roman":"V"}]
```


## Docker
Romanum can also be run as a container

```bash
docker build -t romanum .
docker run --rm -d -p 8080:8080 --name romanum romanum
curl "localhost:8080/convert?start=3&end=5"
# [{"number":3,"roman":"III"},{"number":4,"roman":"IV"},{"number":5,"roman":"V"}]
docker stop romanum
```

# Progress
- [x] Ensure the client can submit a range of numbers (e.g., 10-50).
- [x] Handle numbers within the range of 1 to 3999; any numbers outside this range should return an error message.
- [x] The server must accurately convert all numbers in the specified range to Roman numerals.
- [x] Collect and return all results in ascending order.
- [x] Provide the OpenAPI specification.
- [x] Wrap the service into a Docker container.
- [x] Create an integration test suite to automate the evaluation of the functionality.

# Improvements
- Parameterize hostname and port for the service
- Memoize `ToRoman` conversion function (or turn it into a lookup table)
- Improve logging formatting and introduce various levels of verbosity
- Improve error handling and introduce more specific error codes and messages

# Production Readiness
- Reliability and Scalability: Deploy the app in a highly available, auto-scaling environment (e.g. Kubernetes, AWS ECS...)
- Security: Enforce HTTPS and introduce rate limiting to prevent service abuse or DDoS attacks
- Authentication: Introduce authentication mechanisms if the API is not intended for public use
- Observability: Capture logs and monitor key metrics to keep track of the API's health and performance
- Change Management: Set up CICD pipelines to automate building, testing and deployment
- Docs: Leverage openAPI specs to generate documentation automatically (e.g. using Swagger)
