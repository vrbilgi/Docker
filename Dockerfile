FROM golang:1.18-alpine

WORKDIR /app
COPY client.go ./
COPY server.go ./
RUN go build client.go
RUN go build server.go
EXPOSE 5000
CMD ["./server" ] 

