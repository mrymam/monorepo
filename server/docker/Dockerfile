FROM golang:1.19 as builder
RUN mkdir /build
COPY ./server /build/
WORKDIR /build
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -o server .

ADD https://github.com/benbjohnson/litestream/releases/download/v0.3.8/litestream-v0.3.8-linux-amd64-static.tar.gz /tmp/litestream.tar.gz
RUN tar -C /usr/local/bin -xzf /tmp/litestream.tar.gz

FROM debian:buster
COPY --from=builder /build/server /go/bin/server
COPY --from=builder /usr/local/bin/litestream /usr/local/bin/litestream

RUN mkdir /sqlite
COPY ./sqlite /sqlite
COPY ./server/docker/litestream.yml /usr/bin/litestream.yml
COPY ./server/docker/entrypoint.sh /scripts/entrypoint.sh

RUN apt update
RUN apt install -y sqlite3
RUN apt install -y ca-certificates && update-ca-certificates

ENV SQLITE_FILEPATH /sqlite/prod.sqlite3
ENV ENV prod

EXPOSE 9000

RUN chmod +x /scripts/entrypoint.sh

CMD ["/scripts/entrypoint.sh"]