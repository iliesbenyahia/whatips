FROM golang:1.24.4 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o whatips .

FROM alpine AS run 

COPY --from=builder /app/whatips .

CMD ["./whatips"]