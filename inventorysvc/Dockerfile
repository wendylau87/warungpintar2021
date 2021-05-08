## We specify the base image we need for our
## go application
FROM golang:1.15.12-alpine3.13
## We create an /app directory within our
## image that will hold our application source
## files
RUN mkdir /app
## We copy everything in the root directory
## into our /app directory
ADD . /app
## We specify that we now wish to execute
## any further commands inside our /app
## directory
WORKDIR /app
## we run go build to compile the binary
## executable of our Go program

RUN apk update && apk add --no-cache  \
        bash              \
        build-base        \
        coreutils         \
        gcc               \
        git               \
        make              \
        musl-dev          \
        openssl-dev       \
        openssl           \
        libsasl           \
        libgss-dev        \
        rpm               \
        lz4-dev           \
        zlib-dev          \
        ca-certificates   \
        wget          &&  \
    cd $(mktemp -d) && \
    wget -nv -O cyrus-sasl-2.1.27.tar.gz https://github.com/cyrusimap/cyrus-sasl/releases/download/cyrus-sasl-2.1.27/cyrus-sasl-2.1.27.tar.gz && \
    tar -xz --strip-components=1 -f cyrus-sasl-2.1.27.tar.gz && \
    rm -f cyrus-sasl-2.1.27.tar.gz && \
    ./configure --prefix=/usr --disable-sample --disable-obsolete_cram_attr --disable-obsolete_digest_attr --enable-static --disable-shared \
        --disable-checkapop --disable-cram --disable-digest --enable-scram --disable-otp --disable-gssapi --with-dblib=none --with-pic && \
    make && \
    make install

RUN cd $(mktemp -d) && \
    wget -nv -O v1.3.0.tar.gz https://github.com/edenhill/librdkafka/archive/v1.3.0.tar.gz && \
    tar -xz --strip-components=1 -f v1.3.0.tar.gz && \
    rm -f v1.3.0.tar.gz && \
    ./configure --prefix=/usr --enable-sasl && \
    make -j && \
    make install

ENV GO111MODULE=on \
    GOOS=linux \
    GOARCH=amd64

RUN go get -d -v

RUN go build -tags musl -o app

EXPOSE 8092
## Our start command which kicks off
## our newly created binary executable
CMD ["./app"]