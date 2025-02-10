FROM golang:1.23 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o app

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/app .

COPY utils/ForgotPassword.html ./utils/


EXPOSE 8080

CMD ["./app"]