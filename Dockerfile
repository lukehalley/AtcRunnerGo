FROM golang:1.19

WORKDIR /app

COPY main.go go.mod go.sum ./

ADD internal ./internal

RUN go mod download

RUN go build -o /atc-runner

CMD [ "/atc-runner" ]

