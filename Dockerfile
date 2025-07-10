## Multistage docker file 
# Build stage
FROM golang:1.24-alpine3.22 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main main.go

# Run stage
FROM alpine:3.22
WORKDIR /app

# Copy the binary
COPY --from=builder /app/main .

# Copy the config file ...verify exact name
COPY app.env /app/app.env

# Copy any other required files...ike TLS certs
# COPY *.pem /app/

EXPOSE 8080
CMD ["/app/main"]




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