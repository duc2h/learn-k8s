# build binary
FROM golang:1.18.1 as builder

WORKDIR /

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o backend 


# run binary
FROM alpine:3.12 

WORKDIR /
# copy binary file (name: backend) from previous step
COPY --from=builder ./backend ./backend

CMD ["./backend"]