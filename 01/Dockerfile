FROM golang:1.7.3
WORKDIR /go/src/hello-arch
COPY app/main.go .
RUN CGO_ENABLED=0 GOOS=linux go build .

FROM alpine:latest
EXPOSE 8000
COPY --from=0 /go/src/hello-arch/hello-arch /usr/local/bin
ENTRYPOINT [ "/usr/local/bin/hello-arch" ]