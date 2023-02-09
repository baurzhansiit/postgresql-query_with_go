FROM golang:1.15-alpine

WORKDIR /app
COPY . /app/
RUN go get github.com/lib/pq
RUN go build -o myapp
RUN ls -la /app/myapp

CMD ["./myapp"]