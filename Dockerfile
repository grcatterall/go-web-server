FROM golang:1.22-bullseye

# BUILDER
RUN apt-get update -y
RUN apt-get install -y iputils-ping

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . ./

RUN go build -o web-server ./cmd

EXPOSE 80

CMD ["./web-server"]
