package bsc

import (
	"fmt"

	"github.com/ethereum/go-ethereum/ethclient"
)

var Client *ethclient.Client

func InitClient() {
	var err error
	Client, err = ethclient.Dial("https://bsc.blockrazor.xyz")
	if err != nil {
		fmt.Printf("failed to connect to the bsc node: %v", err)
		return
	}
}
