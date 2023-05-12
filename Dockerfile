# use official Golang image
FROM golang:1.16.3-alpine3.13

# set working directory
WORKDIR /app

# Copy the source code
COPY . . 

# Download and install the dependencies
RUN go get -d -v ./...

# Build the Go app
RUN go build -o api .

#EXPOSE the port
EXPOSE 8080

# Run the executable
CMD ["./api"]

# when we're building the image using the Dockerfile, we'll never see the CMD in the terminal because the container layer will be another layer on top of the image to run this container 