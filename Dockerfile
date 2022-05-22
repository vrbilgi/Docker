FROM golang:1.18-alpine

WORKDIR /app
COPY client.go ./
RUN go build client.go
#EXPOSE command is just indication or documentation 
#purpose. Which may not be the right port where server is listening
EXPOSE 5000
CMD ["./server" ] 

