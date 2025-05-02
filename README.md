# receipt-processor

Solution to the Receipt Processor challenge.

## Running

```shell
go mod tidy
go run main.go
```

Or using Docker:

```shell
docker build -t receipt-processor:latest .
docker run -p 8080:8080 -it receipt-processor:latest
```
