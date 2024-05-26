FROM golang:1.22
# Copy app code into the image
WORKDIR /app
COPY . .
# Run tests and build app binary
RUN go test -v ./... && go build -o romanum ./cmd/server
# Expose the port the app runs on
EXPOSE 8080
# Start app by default
CMD ["./romanum"]
