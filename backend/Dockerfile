FROM golang:1.19.2

WORKDIR /app
ADD . /app
RUN go build

ENTRYPOINT /app/giftexchangeinator
