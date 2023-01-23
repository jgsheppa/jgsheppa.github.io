FROM golang:1.19.5-alpine3.17 AS builder
WORKDIR /

COPY go.mod ./
COPY go.sum ./

RUN go mod download
RUN go build -o jgsheppa

COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM scratch
WORKDIR /root/
COPY --from=builder /app ./

CMD ["./jgsheppa"]