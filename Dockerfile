FROM golang:1.14.6-alpine3.12
RUN mkdir /p2p-oracle
COPY . /p2p-oracle
WORKDIR /p2p-oracle

RUN apk add git
RUN apk add --update make
RUN go mod download

RUN make build

CMD ["./main"]