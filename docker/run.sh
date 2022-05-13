#!/bin/sh
#set -e

peers=""
#hosts=(${NODE_HOSTS//,/ })
#for host in "${hosts[@]}"
#do
  host=$NODE_HOSTS
  gnchaind status --node tcp://$host:26657 2>&1 > /dev/null
  if [ $? -eq 0 ]; then
    pid=`gnchaind status --node tcp://$host:26657 | jq .NodeInfo.id | sed 's/\"//g'`
    if [ -z $GNCHAIND_P2P_PERSISTENT_PEERS ]; then
      export GNCHAIND_P2P_PERSISTENT_PEERS=$pid@$host:26656
    else
      export GNCHAIND_P2P_PERSISTENT_PEERS=$GNCHAIND_P2P_PERSISTENT_PEERS,$pid@$host
    fi
  fi
#done

echo "peers" $GNCHAIND_P2P_PERSISTENT_PEERS

gnchaind start --rpc.laddr tcp://0.0.0.0:26657
