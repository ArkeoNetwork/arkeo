#!/bin/bash

set -o pipefail
set -ex

CHAIN_ID="arkeo-testnet-v2"
STAKE="50000000000000000uarkeo"
TOKEN="uarkeo"
USER="ark"

add_module() {
	jq --arg ADDRESS "$1" --arg ASSET "$2" --arg AMOUNT "$3" --arg NAME "$4" '.app_state.auth.accounts += [{
        "@type": "/cosmos.auth.v1beta1.ModuleAccount",
        "base_account": {
          "address": $ADDRESS,
          "pub_key": null,
          "sequence": "0"
        },
        "name": $NAME,
        "permissions": []
  }]' <~/.arkeo/config/genesis.json >/tmp/genesis.json
	mv /tmp/genesis.json ~/.arkeo/config/genesis.json

	jq --arg ADDRESS "$1" --arg ASSET "$2" --arg AMOUNT "$3" '.app_state.bank.balances += [{
        "address": $ADDRESS,
        "coins": [ { "denom": $ASSET, "amount": $AMOUNT } ],
    }]' <~/.arkeo/config/genesis.json >/tmp/genesis.json
	mv /tmp/genesis.json ~/.arkeo/config/genesis.json
}

add_account() {
	jq --arg ADDRESS "$1" --arg ASSET "$2" --arg AMOUNT "$3" '.app_state.auth.accounts += [{
        "@type": "/cosmos.auth.v1beta1.BaseAccount",
        "address": $ADDRESS,
        "pub_key": null,
        "account_number": "0",
        "sequence": "0"
    }]' <~/.arkeo/config/genesis.json >/tmp/genesis.json
	mv /tmp/genesis.json ~/.arkeo/config/genesis.json

	jq --arg ADDRESS "$1" --arg ASSET "$2" --arg AMOUNT "$3" '.app_state.bank.balances += [{
        "address": $ADDRESS,
        "coins": [ { "denom": $ASSET, "amount": $AMOUNT } ],
    }]' <~/.arkeo/config/genesis.json >/tmp/genesis.json
	mv /tmp/genesis.json ~/.arkeo/config/genesis.json
}

add_claim_records() {
	jq --arg CHAIN "$1" --arg ADDRESS "$2" --arg AMOUNTCLAIM "$3" --arg AMOUNTVOTE "$4" --arg AMOUNTDELEGATE "$5" --arg ISTRANSFERABLE "$6" '.app_state.claimarkeo.claim_records += [{
        "chain": $CHAIN,
		"address": $ADDRESS,
        "amount_claim": { "denom": "uarkeo", "amount": $AMOUNTCLAIM },
        "amount_vote": { "denom": "uarkeo", "amount": $AMOUNTVOTE },
        "amount_delegate": { "denom": "uarkeo", "amount": $AMOUNTDELEGATE },
        "is_transferable": true,
    }]' <~/.arkeo/config/genesis.json >/tmp/genesis.json
	mv /tmp/genesis.json ~/.arkeo/config/genesis.json
}

if [ ! -f ~/.arkeo/config/priv_validator_key.json ]; then
	# remove the original generate genesis file, as below will init chain again
	rm -rf ~/.arkeo/config/genesis.json
fi

if [ ! -f ~/.arkeo/config/genesis.json ]; then
	arkeod init local --staking-bond-denom $TOKEN --chain-id "$CHAIN_ID"
	arkeod keys add $USER --keyring-backend test
	arkeod add-genesis-account $USER $STAKE --keyring-backend test
	arkeod keys list --keyring-backend test
	arkeod gentx $USER $STAKE --chain-id $CHAIN_ID --keyring-backend test
	arkeod collect-gentxs

	arkeod keys add faucet --keyring-backend test
	FAUCET=$(arkeod keys show faucet -a --keyring-backend test)
	add_account "$FAUCET" $TOKEN 10000000000000000 # faucet, 100m

	if [ "$NET" = "mocknet" ] || [ "$NET" = "testnet" ]; then
		add_module tarkeo1d0m97ywk2y4vq58ud6q5e0r3q9khj9e3unfe4t $TOKEN 10000000000000000 'arkeo-reserve' # reserve, 100m
		add_module tarkeo14tmx70mvve3u7hfmd45vle49kvylk6s2wllxny $TOKEN 10000000000000000 'claimarkeo'

		echo "shoulder heavy loyal save patient deposit crew bag pull club escape eyebrow hip verify border into wire start pact faint fame festival solve shop" | arkeod keys add alice --keyring-backend test --recover
		ALICE=$(arkeod keys show alice -a --keyring-backend test)
		add_account "$ALICE" $TOKEN 1100000000000000 # alice, 11m

		echo "clog swear steak glide artwork glory solution short company borrow aerobic idle corn climb believe wink forum destroy miracle oak cover solid valve make" | arkeod keys add bob --keyring-backend test --recover
		BOB=$(arkeod keys show bob -a --keyring-backend test)
		add_account "$BOB" $TOKEN 1000000000000000 # bob, 10m
		add_claim_records "ARKEO" "$BOB" 1000 1000 1000 true

		# Thorchain derived test addresses
		add_account "tarkeo1dllfyp57l4xj5umqfcqy6c2l3xfk0qk6zpc3t7" $TOKEN 1000000000000000 # bob, 10m
		add_claim_records "ARKEO" "tarkeo1dllfyp57l4xj5umqfcqy6c2l3xfk0qk6zpc3t7" 1000 1000 1000 true
		add_account "tarkeo1xrz7z3zwtpc45xm72tpnevuf3wn53re8q4u4nr" $TOKEN 1000000000000000
		add_claim_records "ARKEO" "tarkeo1xrz7z3zwtpc45xm72tpnevuf3wn53re8q4u4nr" 1000 1000 1000 true

		# add_claim_records "ARKEO" "{YOUR ARKEO ADDRESS}" 500000 500000 500000 true
		# add_account "{YOUR ARKEO ADDRESS}" $TOKEN 1000000000000000

		# add_claim_records "ETHEREUM" "{YOUR ETH ADDRESS}" 500000 600000 700000 true
		add_claim_records "ETHEREUM" "0x92E14917A0508Eb56C90C90619f5F9Adbf49f47d" 500000 600000 700000 true

		# enable CORs on testnet/localnet
		sed -i 's/enabled-unsafe-cors = false/enabled-unsafe-cors = true/g' ~/.arkeo/config/app.toml
		sed -i 's/cors_allowed_origins = \[\]/cors_allowed_origins = \["*"\]/g' ~/.arkeo/config/config.toml
	fi

	sed -i 's/"stake"/"uarkeo"/g' ~/.arkeo/config/genesis.json
	sed -i 's/enable = false/enable = true/g' ~/.arkeo/config/app.toml
	sed -i 's/127.0.0.1:26657/0.0.0.0:26657/g' ~/.arkeo/config/config.toml

	set -e
	arkeod validate-genesis --trace
fi

arkeod start
