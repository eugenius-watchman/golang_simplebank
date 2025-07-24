# Build stage
FROM golang:1.24-alpine3.22 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /app/main main.go

# Install migrate (fixed extraction path)
RUN apk add curl tar
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.18.3/migrate.linux-amd64.tar.gz | tar xvz -C /app

# Run stage
FROM alpine:3.22
WORKDIR /app

# Copy binaries
COPY --from=builder /app/main .
COPY --from=builder /app/migrate ./migrate 

# Copy config and migrations
COPY app.env .
COPY start.sh .
COPY wait-for.sh .
COPY db/migration ./migration

EXPOSE 8080
CMD ["/app/main"]

## Multistage docker file 
# # Build stage
# FROM golang:1.24-alpine3.22 AS builder
# WORKDIR /app
# COPY go.mod go.sum ./
# RUN go mod download
# COPY . .
# RUN go build -o /app/main main.go 
# # RUN go build -o main main.go

# RUN apk add curl
# RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.18.3/migrate.linux-amd64.tar.gz | tar xvz

# # Run stage
# FROM alpine:3.22
# WORKDIR /app

# # Copy the binary
# COPY --from=builder /app/main .

# COPY --from=builder /app/migrate.linux-amd64 ./migrate

# # Copy the config file ...verify exact name
# COPY app.env /app/app.env

# COPY db/migration ./migration

# # Copy any other required files...ike TLS certs
# # COPY *.pem /app/

# EXPOSE 8080
# CMD ["/app/main"]
# ENTRYPOINT ["/app/start.sh"]




# FROM golang:1.24-alpine3.22 AS builder
# WORKDIR /app
# COPY . .
# RUN go build -o main main.go

# # Run stage
# FROM alpine:3.22
# WORKDIR /app
# COPY --from=builder /app/main .
# COPY app.env .

# EXPOSE 8080
# CMD [ "/app/main" ]