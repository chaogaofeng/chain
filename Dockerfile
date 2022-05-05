FROM golang:1.17-alpine3.15 AS go-builder

# this comes from standard alpine nightly file
#  https://github.com/rust-lang/docker-rust-nightly/blob/master/alpine3.12/Dockerfile
# with some changes to support our toolchain, etc
RUN set -eux; apk add --no-cache ca-certificates build-base;

# Set up dependencies
ENV PACKAGES make gcc git libc-dev bash openssl
# NOTE: add these to run with LEDGER_ENABLED=true
# RUN apk add libusb-dev linux-headers

# Install minimum necessary dependencies
RUN apk add $PACKAGES

WORKDIR /code
COPY . /code/

RUN sed -i '/\.\./d' go.mod
RUN sed -i 's/\/\///g' go.mod

# See https://github.com/CosmWasm/wasmvm/releases
ADD https://github.com/CosmWasm/wasmvm/releases/download/v1.0.0-beta6/libwasmvm_muslc.a /lib/libwasmvm_muslc.a
RUN sha256sum /lib/libwasmvm_muslc.a | grep e9cb9517585ce3477905e2d4e37553d85f6eac29bdc3b9c25c37c8f5e554045c

# force it to use static lib (from above) not standard libgo_cosmwasm.so file
RUN LEDGER_ENABLED=false BUILD_TAGS=muslc make build

# --------------------------------------------------------
FROM alpine:3.15

## Set up dependencies
#ENV PACKAGES make gcc perl wget
#
## Install openssl 3.0.0
#RUN apk add $PACKAGES \
#    && wget https://github.com/openssl/openssl/archive/openssl-3.0.2.tar.gz \
#    && tar -xzvf openssl-3.0.2.tar.gz \
#    && cd openssl-openssl-3.0.2 && ./config \
#    && make install \
#    && cd ../ && rm -fr openssl-openssl-3.0.2

COPY --from=go-builder /code/build/gnchaind /usr/bin/gnchaind

WORKDIR /root

COPY docker/* /root

# rest server
EXPOSE 1317
# tendermint p2p
EXPOSE 26656
# tendermint rpc
EXPOSE 26657

CMD ["/root/run.sh"]