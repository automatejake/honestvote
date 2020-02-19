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
CMD ./main --tcp 7002 --http 7003 --role producer --collection-prefix a_ --registry true --institution-name "West Chester University" --private-key "3077020101042061a9cec9f6502df62d8190c008ef29f485142e8e03de993e469c250966e574ada00a06082a8648ce3d030107a1440342000420f6ae9be26dfde8b50f550bfb273ad77d1012a9c427f4e5ea761faa108ab0b69a042448b15e09c67075cba02931c2ae602b9125afad8f0480f83d24c55d3bc5" --public-key "3059301306072a8648ce3d020106082a8648ce3d0301070342000420f6ae9be26dfde8b50f550bfb273ad77d1012a9c427f4e5ea761faa108ab0b69a042448b15e09c67075cba02931c2ae602b9125afad8f0480f83d24c55d3bc5" --database-host "mongo" --hostname "portainer.honestvote.io"
# CMD while true; do sleep 1000; done
