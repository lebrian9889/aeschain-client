# Aeternalism Go client

Aeternalism Go client will listen to LockNFT event from Locker's smart contracts from aeschain-ethereum repo then call NFT module in Aeschain to create NFT in the cosmos

## Install and deploy contracts

1. Prerequisites:
- Golang version >= 1.16
- Create .env file in the root folder then add
  - ETHEREUM_NETWORK_WS: websocket endpoint for ethereum ropsten network (create free account at infura.io)
  - LOCKER_ADDRESS: locker's smart contracts address that deployed to ropsten network from aeschain-ethereum repo
  - COSMOS_PREFIX: cosmos address prefix (aes for the demo)
  - COSMOS_ACCOUNT: cosmos account that used for interacting with Aeschain (brian for the demo)

2. Install dependencies:
Run command in the project root to install dependencies
```
go mod tidy
```

3. Run the application
- Run the command to start listening event
```
go run main.go
```
