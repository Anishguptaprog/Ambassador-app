FROM golang:1.23.1

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# binary will be $(go env GOPATH)/bin/air
RUN curl -sSfL https://raw.githubusercontent.com/air-verse/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

CMD [ "air" ]
