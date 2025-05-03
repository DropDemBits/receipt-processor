FROM golang:1.23.8
WORKDIR /

# Server HTTP port
EXPOSE 8080
ENV PORT=8080

# Copy source files
COPY controllers ./controllers
COPY processor ./processor
COPY model ./model
COPY routes ./routes
COPY main.go ./
# Copy module snapshot
COPY go.mod go.sum ./

ENV GIN_MODE=debug
RUN go build

ENTRYPOINT ["./receipt-processor"]