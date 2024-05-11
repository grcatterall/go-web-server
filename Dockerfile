FROM golang:1.22-bullseye

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . ./

RUN go build -o web-server ./cmd

EXPOSE 80

CMD ["./web-server"]
