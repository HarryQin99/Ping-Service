FROM golang:1.16-alpine 

WORKDIR /app

COPY . /app

RUN go build -o /ping-service

CMD [ "/ping-service" ]