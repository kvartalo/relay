# Relay [![Go Report Card](https://goreportcard.com/badge/github.com/kvartalo/relay)](https://goreportcard.com/report/github.com/kvartalo/relay) [![GoDoc](https://godoc.org/github.com/kvartalo/relay?status.svg)](https://godoc.org/github.com/kvartalo/relay)
Kvartalo relay


## Usage
Create new wallet keystore:
```
./relay wallet new
```

Deploy the [token](https://github.com/kvartalo/token) smart contract:
```
./relay contracts token deploy
```

(For docs about how to build the Go code for `token` contract, see https://github.com/kvartalo/relay/tree/master/eth/token/README.md )


Set the config file `./config.yaml`
```yaml
keystorage:
        address: "0xaddress"
        password: "secretpassword"
        keyjsonpath: "./keystorage/keystoragepath"
server:
        Port: "portnumber"
web3:
        url: "web3gateway-url"
        startscanblock: fist-block-to-scan-transfer-txs

contracts:
        token: "0xdeployedAddr"
storage:
        path: ./databasepath

```

Run Relay
```
./relay start
```

### Commands
```sh
# initializes the database
./relay init

# info about the relay address, eth and tokens
./relay info

# create new keystore
./relay wallet new

# info of the current keystores
./relay wallet info

# deploy the token smart contract
./relay contracts token deploy

# mint tokens that will go to the relay address
./relay contracts token mint [amount]

# transfer tokens to a specified address
./relay contracts token transfer [address] [amount]

# start the relay server
./relay start
```


### API Endpoints

- **GET /balance/:addr**

Returns:
```json
{
  "addr": "0x104d6ea578124aa891163a0da51315cccc823939",
  "balance": "23300"
}
```

- **GET /history/:addr**

Returns:
```json
{
  "addr": "0x104d6ea578124aa891163a0da51315cccc823939",
  "count": 2,
  "transfers": [
    {
      "Timestamp": 1556194387,
      "From": "0x104d6ea578124aa891163a0da51315cccc823939",
      "To": "0x653ad143cd78e47db6846d1bd3cc0f69a6fb9225",
      "Value": 3000
    },
    {
      "Timestamp": 1556013327,
      "From": "0x104d6ea578124aa891163a0da51315cccc823939",
      "To": "0x3a05826f99a4ae1460e7c9b39d87dba74b09c3ad",
      "Value": 3000
    }
  ]
}
```

- **GET /tx/nonce/:addr**

Returns:
```json
{
  "addr": "0x104d6ea578124aa891163a0da51315cccc823939",
  "nonce": 11
}
```

- **POST /tx**

Input:
```json
{
  "from": "0x104d6ea578124aa891163a0da51315cccc823939",
  "to": "0x653ad143cd78e47db6846d1bd3cc0f69a6fb9225",
  "value": 3000,
  "r": "sig.r parameter in hex",
  "s": "sig.s parameter in hex",
  "v": 0
}
```

Returns:
```json
{
  "ethTx": "0xtxhash"
}
```
