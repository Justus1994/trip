FROM golang:alpine AS builder
RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/hb/api
COPY . .
RUN go get -d github.com/gorilla/mux && go get -u github.com/go-redis/redis && go get -u github.com/google/uuid
RUN CGO_ENABLED=0 go build -ldflags "-s -w"  -o /go/bin/api
FROM scratch AS final
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/bin/api /go/bin/api
EXPOSE 8080
ENTRYPOINT ["/go/bin/api"]
