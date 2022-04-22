#!/bin/sh
set -o errexit -o nounset -o pipefail

PASSWORD=${PASSWORD:-1234567890} # 钱包密码
STAKE=${STAKE_TOKEN:-stake} # 权益代币
FEE=${FEE_TOKEN:-ugnc} # 手续费代币
CHAIN_ID=${CHAIN_ID:-testing} # 链ID
MONIKER=${MONIKER:-node} # 节点名称

gnchaind init --chain-id "$CHAIN_ID" "$MONIKER"
# staking/governance token is hardcoded in config, change this
sed -i "s/\"stake\"/\"$STAKE\"/" "$HOME"/.gnchain/config/genesis.json
# this is essential for sub-1s block times (or header times go crazy)
sed -i 's/"time_iota_ms": "1000"/"time_iota_ms": "10"/' "$HOME"/.gnchain/config/genesis.json

sed -i "s/\"os\"/\"test\"/" "$HOME"/.gnchain/config/client.toml

if ! gnchaind keys show validator; then
  (echo "$PASSWORD"; echo "$PASSWORD") | gnchaind keys add validator
fi
# hardcode the validator account for this instance
echo "$PASSWORD" | gnchaind add-genesis-account validator "10000000000$STAKE,10000000000$FEE"

# (optionally) add a few more genesis accounts
for addr in "$@"; do
  echo $addr
  gnchaind add-genesis-account "$addr" "10000000000$FEE"
done

# submit a genesis validator tx
(echo "$PASSWORD"; echo "$PASSWORD"; echo "$PASSWORD") | gnchaind gentx validator "1000000$STAKE" --chain-id="$CHAIN_ID" 
gnchaind collect-gentxs
