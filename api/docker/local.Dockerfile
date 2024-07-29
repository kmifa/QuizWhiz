FROM golang:1.22-alpine

RUN apk --no-cache add ca-certificates curl

#Download cosmtrek/air
RUN go install github.com/cosmtrek/air@v1.43.0

# execute air binary
WORKDIR /go/bin
CMD ["air"]
