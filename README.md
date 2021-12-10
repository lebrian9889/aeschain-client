# Aeternalism Go client
Aeternalism Go client will listen to the LockNFT event from Locker's smart contracts from aeschain-ethereum repo and then call the NFT module in Aeschain to create NFT in the cosmos.

## Install and deploy contracts.

1. Prerequisite:
- Golang version >= 1.16
- Create .env file in the root folder, then add
  - ETHEREUM_NETWORK_WS: websocket endpoint for the ethereum ropsten network (create a free account at infura.io)
  - LOCKER_ADDRESS: locker's smart contract address that is deployed to the ropsten network from the aeschain-ethereum repo
  - COSMOS_PREFIX: cosmos address prefix (aes for the demo)
  - COSMOS_ACCOUNT: Cosmos account that is used for interacting with Aeschain (brian for the demo)

2. Install dependencies:

Run the command in the project root to install dependencies.
```
go mod tidy
```
3. Run the application
- Run the command to start listening to events.
```
go run main.go
```
