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
WORKDIR /app/build
RUN ./install-binaries.sh

# Build main app
WORKDIR /app/core/core-p2p/p2p
RUN go build -o main

# Expose port 7000 and 7001 to host machine
EXPOSE 7000
EXPOSE 7001

# Command to run the executable
CMD ["./main"] 
# CMD while true; do sleep 1000; done
