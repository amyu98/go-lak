FROM golang:1.19.1-alpine3.16 as builder

COPY go.mod go.sum /go/src/github.com/amyu98/go-lak/
WORKDIR /go/src/github.com/amyu98/go-lak
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/lak github.com/amyu98/go-lak

FROM alpine

#RUN apk add --no-cache ca-certificates && update-ca-certificates
COPY --from=builder /go/src/github.com/amyu98/go-lak/build/lak /usr/bin/lak

WORKDIR  /go/src/github.com/amyu98/go-lak

EXPOSE 8080 8080
