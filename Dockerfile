FROM golang:alpine AS builder
RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/mypackage/myapp/
COPY . .
RUN go get -d -v
RUN CGO_ENABLED=0 GOOS=linux go build -o /go/bin/main

FROM scratch
COPY --from=builder /go/bin/main /main
COPY templates/ /templates
ENTRYPOINT ["/main"]
