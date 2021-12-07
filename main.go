package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/spf13/viper"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	aesnft "github.com/lebrian9889/aeschain-client/ethereum/contracts/aesnft"
	locker "github.com/lebrian9889/aeschain-client/ethereum/contracts/locker"

	aeternalismtypes "github.com/lebrian9889/aeternalism/x/nft/types"
	"github.com/tendermint/starport/starport/pkg/cosmosclient"
)

func getEnv(key string) string {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	val, ok := viper.Get(key).(string)
	if !ok {
		log.Fatal("Invalid type")
	}

	return val
}

func main() {
	ethNetwork := getEnv("ETHEREUM_NETWORK_WS")
	lockerAddr := getEnv("LOCKER_ADDRESS")

	client, err := ethclient.Dial(ethNetwork)
	if err != nil {
		log.Fatal(err)
	}

	chainId, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected chain", chainId)
	fmt.Println("Listening events...")

	// load smart contract
	lockerAddress := common.HexToAddress(lockerAddr)
	lockerContract, err := locker.NewLocker(lockerAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	_ = lockerContract

	query := ethereum.FilterQuery{
		Addresses: []common.Address{lockerAddress},
	}

	logs := make(chan types.Log)

	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	constractAbi, err := abi.JSON(strings.NewReader(string(locker.LockerABI)))
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			event := struct {
				NftAddress common.Address
				From       common.Address
				TokenId    *big.Int
				Nounce     *big.Int
				Data       string
			}{}
			err := constractAbi.UnpackIntoInterface(&event, "LockNFT", vLog.Data)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("NFT Address: ", event.NftAddress)
			fmt.Println("From: ", event.From)
			fmt.Println("TokenID: ", event.TokenId)
			fmt.Println("Nounce: ", event.Nounce)
			fmt.Println("Arbitrage data: ", event.Data)

			fmt.Println("Calling minting module")

			// load nft contract
			aesnftContract, err := aesnft.NewAesnft(event.NftAddress, client)
			if err != nil {
				log.Fatal(err)
			}

			mediaUrl, err := aesnftContract.TokenURI(&bind.CallOpts{}, event.TokenId)
			if err != nil {
				log.Fatal(err)
			}

			createErr := createNft(mediaUrl, event.Data, event.TokenId.String(), event.From.String(), "ETHEREUM", event.NftAddress.String(), event.Nounce.String())
			if createErr != nil {
				log.Fatal(createErr)
			}
		}
	}

}

func createNft(mediaUrl string, owner string, orgId string, orgOwner string, orgChain string, orgAddress string, orgNounce string) error {
	addressPrefix := getEnv("COSMOS_PREFIX")
	caller := getEnv("COSMOS_ACCOUNT")

	aes, err := cosmosclient.New(context.Background(), cosmosclient.WithAddressPrefix(addressPrefix))
	if err != nil {
		return err
	}

	// get info from caller
	account, err := aes.Account(caller)
	if err != nil {
		return err
	}
	address := account.Address(addressPrefix)

	// create-nft msg
	msg := &aeternalismtypes.MsgCreateNft{
		Creator:    address,
		MediaUrl:   mediaUrl,
		Owner:      owner,
		OrgId:      orgId,
		OrgOwner:   orgOwner,
		OrgChain:   orgChain,
		OrgAddress: orgAddress,
		OrgNounce:  orgNounce,
	}

	txRes, err := aes.BroadcastTx(caller, msg)
	if err != nil {
		return err
	}

	fmt.Println("Create NFT response")
	fmt.Println(txRes)

	// query new NFT
	queryClient := aeternalismtypes.NewQueryClient(aes.Context)

	nfts, err := queryClient.NftAll(context.Background(), &aeternalismtypes.QueryAllNftRequest{})
	if err != nil {
		return err
	}

	fmt.Println("All nfts")
	fmt.Println(nfts)

	return nil
}
