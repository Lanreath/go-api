FROM golang:1.21.5 AS builder
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . ./
RUN go build -o /api
EXPOSE 5000
CMD ["/api"]