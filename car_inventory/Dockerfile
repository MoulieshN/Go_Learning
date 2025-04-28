FROM golang:1.24 as builder

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go mod vendor

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod=vendor -o userapp

FROM alpine:3.20

WORKDIR /app

COPY --from=builder /app/userapp /app/userapp

EXPOSE 3015

CMD ["./userapp"]
