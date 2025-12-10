FROM golang:1.22

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o evote cmd/server/main.go

EXPOSE 8080

CMD ["./evote"]
