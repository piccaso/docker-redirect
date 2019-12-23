FROM golang:alpine AS builder
WORKDIR /src
COPY src/* .
RUN go build server.go
RUN ls -lah

FROM alpine
COPY --chown=0:0 --from=builder /src/server /server
USER 65534
ENTRYPOINT ["/server"]
EXPOSE 9000