# Relay
Kvartalo relay


## Usage
Create new wallet keystore
```
./relay wallet new
```

Set the config file `./config.yaml`
```yaml
keystorage:
        address: "0xaddress"
        password: "secretpassword"
server:
        Port: "portnumber"
web3:
        url: "web3gatewayurl"
```

Run Relay
```
./relay start
```
