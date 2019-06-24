FROM golang:alpine AS builder
RUN apk update && apk add --no-cache git
WORKDIR /api
COPY . .
RUN CGO_ENABLED=0 go build -ldflags "-s -w"  -o /dist/api
FROM scratch AS final
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /dist/api dist/api
EXPOSE 8080
ENTRYPOINT ["dist/api"]
