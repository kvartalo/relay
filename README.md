# Relay
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
contracts:
        token: "0xdeployedAddr"
```

Run Relay
```
./relay start
```

### Commands
```sh
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

