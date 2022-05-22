# Docker
## How to run:

docker build -f Dockerfile . -t go:server_client

docker run -p 9988:9988 go:server_client

## Communication Pattern
1. Server running inside docker container
2. Client connecting from outside container 