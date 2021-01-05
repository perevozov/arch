FROM golang:1.14.7
WORKDIR /go/src/arch-hw03
COPY app/ .
RUN CGO_ENABLED=0 GOOS=linux go build .

FROM alpine:latest
EXPOSE 8000
WORKDIR /app
COPY --from=0 /go/src/arch-hw03/arch-hw03 /app
COPY init.sql /app

ENTRYPOINT [ "/app/arch-hw03" ]