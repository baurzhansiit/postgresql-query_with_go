FROM golang:1.20-alpine

RUN apk update && apk add postgresql

WORKDIR /app
COPY . /app/

RUN go build -o myapp

RUN adduser app   -D -u 1001 && \
    chown app:app ./myapp

# USER app

CMD ./myapp