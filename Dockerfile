FROM golang:1.16-alpine
WORKDIR /workspace
COPY . .
EXPOSE 8080
CMD ["main.go"]