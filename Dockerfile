FROM golang:1.20.2-alpine3.16

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

COPY . ./

RUN go build -o /crypto-fetch-price

EXPOSE 3000

CMD ["/crypto-fetch-price"]