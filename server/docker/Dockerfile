FROM golang:1.19 as builder
RUN mkdir /build
COPY ./server /build/
WORKDIR /build
RUN CGO_ENABLED=0 GO111MODULE=on go build -o server .

FROM scratch
COPY --from=builder /build/server /go/bin/server

RUN mkdir /sqlite
COPY ./sqlite /sqlite
COPY ./server/docker/litestream.yml /etc/litestream.yml
WORKDIR /server

ADD https://github.com/benbjohnson/litestream/releases/download/v0.3.8/litestream-v0.3.8-linux-amd64-static.tar.gz /tmp/litestream.tar.gz
RUN tar -C /usr/local/bin -xzf /tmp/litestream.tar.gz

RUN apt-get install -y sqlite3

EXPOSE 9000
