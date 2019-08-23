# use golang base image
FROM golang:1.12.0-alpine3.9

# create an /app directory for source code
RUN mkdir /app

# copy source code int /app directory
ADD . /app

# specify work directory
WORKDIR /app

# run go build to compile the binary executable
RUN go build -o main .

# run the api
CMD ["/app/main"]

#use port 8083
EXPOSE 8083