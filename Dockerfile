FROM golang:1.19.5-alpine3.17 AS builder
WORKDIR /

COPY go.mod ./
COPY go.sum ./

RUN go mod download
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM scratch
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app ./

CMD ["./app"]