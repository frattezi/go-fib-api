FROM golang:alpine

RUN mkdir /api
COPY ./api /api
WORKDIR /api

RUN go build -o main .
CMD ["/api/main"]