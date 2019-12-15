# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from the latest golang base image
FROM golang:latest

# Add Maintainer Info
LABEL maintainer="Jacob Neubaum <jacob@bizylife.com>"

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Get dependencies
#RUN go build -o main .
RUN go get github.com/jneubaum/honestvote/core/core-http/src
RUN go get github.com/jneubaum/honestvote/core/core-database/src
RUN go get github.com/jneubaum/honestvote/core/core-consensus/src
RUN go get github.com/jneubaum/honestvote/core/core-crypto/src

RUN go get github.com/joho/godotenv
RUN go get go.mongodb.org/mongo-driver/bson
RUN go get go.mongodb.org/mongo-driver/mongo
RUN go get go.mongodb.org/mongo-driver/mongo/options

# Build main app
WORKDIR /app/core/core-p2p/src
RUN go build -o main

# Expose port 7000 and 7001 to host machine
EXPOSE 7000
EXPOSE 7001

# Command to run the executable
CMD ["./main"] 
# CMD while true; do sleep 1000; done
