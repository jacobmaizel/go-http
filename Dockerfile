FROM golang:1.23-alpine

WORKDIR /app

RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum ./
# COPY ./cmd ./
RUN go mod download
COPY . .

# RUN go build -o ./tmp/main ./cmd/main.go

CMD ["air", "-c", ".air.toml"]
