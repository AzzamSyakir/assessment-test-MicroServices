#start from golang base image
FROM golang:1.22-alpine

#install dependencies
RUN apk update
RUN apk add --no-cache git build-base
RUN go install github.com/air-verse/air@latest


#setup workdir
RUN mkdir -p /app
ENV WORKDIR /app
WORKDIR $WORKDIR


# copy source from host to docker container
COPY ./src/role-service/ ./src/role-service/
COPY grpc/ grpc/
COPY role.air.toml .
COPY .env .
COPY go.mod .
COPY go.sum .
COPY Makefile .

# install dependencies project
WORKDIR $WORKDIR/src/role-service/
RUN go mod download

# run the Go app.
WORKDIR $WORKDIR
CMD ["air", "-c", "role.air.toml"]
