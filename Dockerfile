FROM golang:alpine

WORKDIR /usr/src/notifier

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/notifier ./

EXPOSE 8080

CMD ["/usr/local/bin/notifier"]