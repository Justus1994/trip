FROM golang:alpine AS builder
RUN apk update && apk add --no-cache git # install git to download go modules
WORKDIR /api

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
# Run build 
RUN CGO_ENABLED=0 go build -ldflags "-s -w"  -o /dist/api
FROM scratch AS final
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /dist/api dist/api
EXPOSE 8080
ENTRYPOINT ["dist/api"]
