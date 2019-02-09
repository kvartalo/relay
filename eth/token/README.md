# Kvartalo token
Token code: https://github.com/kvartalo/token

## Token go build
```
solc --abi --bin contracts/Token.sol -o build
abigen --bin=./build/Token.bin --abi=./build/Token.abi --pkg=token --out=token.go
```
