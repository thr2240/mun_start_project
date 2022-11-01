#!/bin/bash

#Build Flag
BINARY=mund
NETWORK=devnet
FUNCTION=$1
CATEGORY=$2
PARAM_1=$3
PARAM_2=$4
PARAM_3=$5

ADDR_LUCA="mun1unr06d9xl8c6f32m7qmqx60cwhp6dm2zuaqfzy"

case $NETWORK in
  devnet)
    NODE="http://localhost:26657"
    DENOM=utmun
    CHAIN_ID=mun
    WALLET="--from luca"
    ADDR_ADMIN=$ADDR_LUCA
    ;;
#   testnet)
#     NODE="https://rpc.juno.giansalex.dev:443"
#     DENOM=ujunox
#     CHAIN_ID=uni-3
#     LP_TOKEN_CODE_ID=123
#     WALLET="--from finalmarble"
#     ADDR_ADMIN=$ADDR_MARBLE
#     TOKEN_MARBLE="juno15s50e6k9s8mac9cmrg2uq85cgw7fxxfh24xhr0chems2rjxsfjjs8kmuje"
#     ;;
#   mainnet)
#     NODE="https://rpc-juno.itastakers.com:443"
#     DENOM=ujuno
#     CHAIN_ID=juno-1
#     LP_TOKEN_CODE_ID=1
#     WALLET="--from finalmarble"
#     ADDR_ADMIN=$ADDR_MARBLE
#     TOKEN_MARBLE="juno1g2g7ucurum66d42g8k5twk34yegdq8c82858gz0tq2fc75zy7khssgnhjl"
#     TOKEN_BLOCK="juno1y9rf7ql6ffwkv02hsgd4yruz23pn4w97p75e2slsnkm0mnamhzysvqnxaq"
#     ;;
esac

NODECHAIN=" --node $NODE --chain-id $CHAIN_ID"
TXFLAG=" $NODECHAIN --gas-prices 0.0$DENOM --gas auto --gas-adjustment 1.3"

###################################################################################################
###################################################################################################
###################################################################################################
###################################################################################################
#Environment Functions

Airdrop() {
    echo "================================================="
    echo "Claim Airdrop"
    UPLOADTX=$($BINARY tx claim claim 1 i $WALLET $TXFLAG --output json -y | jq -r '.txhash')
    
}
#################################################################################
PrintWalletBalance() {
    echo "native balance"
    echo "========================================="
    $BINARY query bank balances $ADDR_ADMIN $NODECHAIN
    echo "========================================="
}

#################################### End of Function ###################################################
if [[ $FUNCTION == "" ]]; then
    printf "y\npassword\n" | Airdrop
    sleep 3
    PrintWalletBalance
else
    $FUNCTION $CATEGORY
fi
