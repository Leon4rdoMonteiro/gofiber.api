FROM golang:1.16.3-alpine AS builder

WORKDIR /usr/src

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . . 

RUN go build -o ./out/app . 

EXPOSE 8000

CMD [ "./app" ]