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
server:
        Port: "portnumber"
web3:
        url: "web3gatewayurl"
contracts:
	token: "deployedAddr"
```

Run Relay
```
./relay start
```
