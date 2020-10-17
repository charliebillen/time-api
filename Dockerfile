FROM golang:alpine AS build
RUN mkdir /build
ADD . /build
WORKDIR /build
RUN CGO_ENABLED=0 GOOS=linux go build -a -o webserver cmd/webserver/main.go

FROM golang:alpine
EXPOSE 8000/tcp
COPY --from=build /build/webserver .
ENTRYPOINT ["./webserver"]
