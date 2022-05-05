FROM golang:1.17-alpine AS go-builder

# Install minimum necessary dependencies
RUN apk update && apk add --no-cache build-base make git libc-dev openssl
# NOTE: add these to run with LEDGER_ENABLED=true
# RUN apk add libusb-dev linux-headers

# See https://github.com/CosmWasm/wasmvm/releases
ADD https://github.com/CosmWasm/wasmvm/releases/download/v1.0.0-beta10/libwasmvm_muslc.x86_64.a /lib/libwasmvm_muslc.a
RUN sha256sum /lib/libwasmvm_muslc.a | grep 2f44efa9c6c1cda138bd1f46d8d53c5ebfe1f4a53cf3457b01db86472c4917ac

WORKDIR /code
COPY . /code/
# force it to use static lib (from above) not standard libgo_cosmwasm.so file
RUN LEDGER_ENABLED=false BUILD_TAGS=muslc LINK_STATICALLY=true make build

# --------------------------------------------------------
FROM alpine

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