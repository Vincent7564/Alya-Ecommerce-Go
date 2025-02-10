# Build stage
FROM golang:1.23 as builder

WORKDIR /app

# Copy and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o Alya-Ecommerce-Go

# Final stage
FROM alpine:latest

WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/app .

# Copy necessary files (e.g. templates, static assets)
COPY utils/ForgotPassword.html ./utils/

EXPOSE 8080

CMD [ "./app" ]
