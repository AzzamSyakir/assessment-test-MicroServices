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
COPY ./src/account-employee-service/ ./src/account-employee-service/
COPY grpc/ grpc/
COPY account.air.toml .
COPY .env .
COPY go.mod .
COPY go.sum .
COPY Makefile .

# install dependencies project
WORKDIR $WORKDIR/src/account-employee-service/
RUN go mod download

# run the Go app.
WORKDIR $WORKDIR
CMD ["air", "-c", "account.air.toml"]
