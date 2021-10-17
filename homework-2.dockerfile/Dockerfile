FROM golang:1.17.2-alpine3.14

COPY src/main.go main.go
RUN go build main.go

EXPOSE 80

ENTRYPOINT [ "./main" ]