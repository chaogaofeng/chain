#!/bin/sh
set -e

if ! [ -f /root/.gnchain/config/genesis.json ]; then
  if [ -n "$RECOVER" ]; then
    echo "$RECOVER" | gnchaind init moniker --recover > /dev/null
  else
    gnchaind init moniker > /dev/null
  fi

  gnchaind chain --home chain
  cp chain/config/genesis.json /root/.gnchain/config/genesis.json
  rm -rf chain
fi