# build binary
FROM golang:1.18.1 as builder

WORKDIR /backend

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod ./
RUN go mod download && go mod verify

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o backend ./

# run
FROM alpine:3.12 

WORKDIR /
COPY --from=builder /backend ./

CMD ["./backend"]