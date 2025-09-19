# pull the base image
FROM golang:latest AS builder
# create base working directory inside container

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o exporter .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/exporter .
EXPOSE 8080
ENTRYPOINT ["./exporter"]
CMD ["start"]