# Using golang base image
FROM golang:1.12.0-alpine3.9

# We create an /app directory for source code
RUN mkdir /app

# Copying source code int /app directory
ADD . /app

RUN ls -latr

# work directory
WORKDIR /app

# run go build to compile the binary executable of our Go program
RUN go build -o main .

RUN ls -latr

# run the api
CMD ["/app/main"]

#use port 3000
EXPOSE 3000