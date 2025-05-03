# receipt-processor

Solution to the Receipt Processor challenge.

## Endpoints

### `/receipts/process`

Body: A JSON-formatted `Receipt`.

Response:

- 200: A JSON-formatted body in the form of:

  ```json
  { "id": "062e96a5-3dcf-4a42-b565-06ea3bc3d2e9" }
  ```

  where `id` can be later used in `GET /receipts/{id}/points`.

- 400: If the receipt did not pass validation, or is otherwise malformed.

### `/receipts/{id}/points`

Body: None

Response:

- 200: A JSON-formatted body in the form of:

  ```json
  { "points": 23 }
  ```

  where `points` is the points awarded to the processed receipt, according to a set of predefined rules.

- 404: If the `id` did not come from a prior `/receipts/process` request.

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
