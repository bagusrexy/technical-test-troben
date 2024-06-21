FROM golang:1.20.3

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

RUN apt-get update && apt-get install -y procps && rm -rf /var/lib/apt/lists/*

COPY . .

RUN go build -o technical-test-troben

EXPOSE 1800

CMD ["go","run","main.go"]