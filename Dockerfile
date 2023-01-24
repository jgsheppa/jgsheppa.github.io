FROM golang:1.19.5-alpine3.17 AS builder
WORKDIR /

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM gcr.io/distroless/base-debian11:nonroot
WORKDIR /

COPY --from=builder /app ./
COPY --from=builder /views ./views
COPY --from=builder /images ./images

EXPOSE 8080

CMD ["./app"]