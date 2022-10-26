FROM golang:1.19 as builder
RUN mkdir /build
COPY ./server /build/
WORKDIR /build
RUN CGO_ENABLED=0 GO111MODULE=on go build -o server .

FROM debian:buster
COPY --from=builder /build/server /go/bin/server

RUN mkdir /sqlite
COPY ./sqlite /sqlite
COPY ./server/docker/litestream.yml /etc/litestream.yml
COPY ./server/docker/entrypoint.sh /etc/entrypoint.sh
WORKDIR /server

ADD https://github.com/benbjohnson/litestream/releases/download/v0.3.8/litestream-v0.3.8-linux-amd64-static.tar.gz /tmp/litestream.tar.gz
RUN tar -C /usr/local/bin -xzf /tmp/litestream.tar.gz

RUN apt update
RUN apt install -y sqlite3

ENV SQLITE_FILEPATH /server/sqlite/prod.sqlite3
ENV ENV prod

EXPOSE 9000

CMD [ "bash", "/etc/entrypoint.sh" ]
