FROM golang:1.17-alpine AS go-builder
ARG ALGO=ed25519

# Install minimum necessary dependencies
RUN apk update && apk add --no-cache build-base make git libc-dev openssl
# NOTE: add these to run with LEDGER_ENABLED=true
# RUN apk add libusb-dev linux-headers

# See https://github.com/CosmWasm/wasmvm/releases
ADD https://github.com/CosmWasm/wasmvm/releases/download/v1.0.0/libwasmvm_muslc.x86_64.a /lib/libwasmvm_muslc.a
RUN sha256sum /lib/libwasmvm_muslc.a | grep f6282df732a13dec836cda1f399dd874b1e3163504dbd9607c6af915b2740479

WORKDIR /code
COPY . /code/
# force it to use static lib (from above) not standard libgo_cosmwasm.so file
RUN ALGO=$ALGO LEDGER_ENABLED=false BUILD_TAGS=muslc LINK_STATICALLY=true make build

# --------------------------------------------------------
FROM alpine
RUN apk add --no-cache jq
ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.9.0/wait /wait
RUN chmod +x /wait

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
COPY docker/* /

WORKDIR /root

# rest server
EXPOSE 1317
# tendermint p2p
EXPOSE 26656
# tendermint rpc
EXPOSE 26657

CMD ["sh", "-c", "/wait && /run.sh"]
