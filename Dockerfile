FROM golang:alpine3.12 AS build
RUN mkdir /build
ADD . /build
WORKDIR /build
RUN CGO_ENABLED=0 GOOS=linux go build -a -o webserver cmd/webserver/main.go

FROM alpine:3.12.0
EXPOSE 8000/tcp
COPY --from=build /build/webserver .
ENTRYPOINT ["./webserver"]
