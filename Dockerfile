# Use the official Golang image as the base image
FROM golang:1.22-alpine

# Metadata as labels
LABEL Author1="brianbaraza40@gmail.com"\
      Author2="hezronokwach@gmail.com"\
      Author3="annmaina614@gmail.com"\
      version="1.0"\
      description="ASCII Art Web Stylize"\
      source="https://github.com/Bantu-art/ascii-art-web-dockerize.git" 


# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod file
COPY go.mod ./

# Download all dependencies. Dependencies will be cached if the go.mod file is not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o main .

# Expose port 8090 to the outside world
EXPOSE 8060

# Run the executable
CMD ["./main"]






