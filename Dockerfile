FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/mypackage/myapp/
COPY . .

RUN go get -d -v

RUN go test
RUN go build -o /app/cmd/golangtester
FROM scratch

COPY --from=builder /app/cmd/golangtester /golangtester


EXPOSE 8080
ENTRYPOINT ["/golangtester"]