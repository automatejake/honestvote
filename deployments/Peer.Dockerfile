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
# CMD ["./main --tcp 7002 --http 7003 --role producer --collection-prefix a_ --registry true --institution-name \"West Chester University\" \
# --private-key \"307702010104201d9618c792ab7c1c44a333968e08a42aba7795000fdd15782b8bfd5b26d542afa00a06082a8648ce3d030107a1440342000474bb2f88fd66196ee5f5b8441ef57bc71576d25d8da1a1614ba12d230b7460d4ccfcf2736a1710688ce27538a18d7e5d0b382c11b287f87119ebc5e71e4c7e32\" \
# --public-key \"3059301306072a8648ce3d020106082a8648ce3d0301070342000474bb2f88fd66196ee5f5b8441ef57bc71576d25d8da1a1614ba12d230b7460d4ccfcf2736a1710688ce27538a18d7e5d0b382c11b287f87119ebc5e71e4c7e32\" --database-host mongo"] 
CMD while true; do sleep 1000; done
