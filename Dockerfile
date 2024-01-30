FROM golang:1.21.6

WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

CMD ["go", "run", "main.go"]
