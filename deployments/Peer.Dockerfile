# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from the latest golang base image
FROM golang:latest

# Add Maintainer Info
LABEL maintainer="Jacob Neubaum <jacob@bizylife.com>"

# Set the Current Working Directory inside the container
# WORKDIR /usr/local/go/src/github.com/jneubaum/honestvote
WORKDIR /go/src/github.com/jneubaum/honestvote
# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Get dependencies



#RUN go build -o main .
# WORKDIR /usr/local/go/src/github.com/jneubaum/honestvote/build
# RUN ./install-binaries.sh
# RUN go build -o main

# Expose port 7000 and 7001 to host machine
EXPOSE 7000
EXPOSE 7001

# Command to run the executable
# WORKDIR /usr/local/go/src/github.com/jneubaum/honestvote/scripts

WORKDIR /go/src/github.com/jneubaum/honestvote/build

RUN chmod +x install-binaries.sh
RUN ./install-binaries.sh
RUN go build -o main
CMD ./main --tcp 7002 --http 7003 --role producer --collection-prefix a_ --registry true --institution-name "Honestvote" --database-host "mongo" --hostname "registry.honestvote.io"
# CMD while true; do sleep 1000; done