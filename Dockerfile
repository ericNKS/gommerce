FROM golang:1.23

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN go build -v -o /usr/local/bin/api ./cmd/api
# RUN go build -v -o /usr/local/bin/worker ./cmd/workers

EXPOSE 3000

CMD ["/usr/local/bin/api"]
