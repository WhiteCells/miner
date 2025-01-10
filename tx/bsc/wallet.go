package bsc

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"miner/tx/bsc/usdt"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
)

type Wallet struct{}

func NewWallet() *Wallet {
	return &Wallet{}
}

// 创建钱包
func (w *Wallet) GenerateWallet() (privateKeyHex, address string) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Error("failed to generate wallet", err.Error())
		return "", ""
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)
	privateKeyHex = hex.EncodeToString(privateKeyBytes)

	publicKey := privateKey.PublicKey
	address = crypto.PubkeyToAddress(publicKey).Hex()

	return privateKeyHex, address
}

// 获取余额
func (w *Wallet) GetBalance(client *ethclient.Client, tokenAddress, userAddress string) (*big.Int, error) {
	tokenContract, err := usdt.NewToken(common.HexToAddress(tokenAddress), client)
	if err != nil {
		fmt.Println("token.NewToken", err.Error())
		return nil, err
	}

	balance, err := tokenContract.BalanceOf(&bind.CallOpts{}, common.HexToAddress(userAddress))
	if err != nil {
		fmt.Println("tokenContract.BalanceOf", err.Error())
		return nil, err
	}

	return balance, nil
}

// 发送代币
func (w *Wallet) SendTokenTransaction(client *ethclient.Client, privateKeyHex, toAddress, tokenAddress string, amount *big.Int) (string, error) {
	// 将私钥从十六进制转换为 ECDSA 格式
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		fmt.Println("crypto.HexToECDSA")
		return "", err
	}

	// 获取发送方地址
	publicKey := privateKey.PublicKey
	fromAddress := crypto.PubkeyToAddress(publicKey)

	// 获取链 ID（例如 BSC 主网链 ID 为 56）
	chainID := big.NewInt(56) // 如果是 BSC 测试网，使用 97

	// 创建交易签名器（用链 ID 确保签名的唯一性）
	txOpts, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		fmt.Println("bind.NewKeyedTransactorWithChainID", err.Error())
		return "", err
	}

	// 获取 nonce
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		fmt.Println("client.PendingNonceAt", err.Error())
		return "", err
	}
	txOpts.Nonce = big.NewInt(int64(nonce))

	// 设置交易选项
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		fmt.Println("client.SuggestGasPrice", err.Error())
		return "", err
	}
	txOpts.Value = big.NewInt(0)    // 发送 BSC-20 代币时无需附加 BNB
	txOpts.GasPrice = gasPrice      // 推荐的 Gas 价格
	txOpts.GasLimit = uint64(60000) // 根据代币合约需求设置 Gas 限制

	// 加载代币合约
	tokenContract, err := usdt.NewToken(common.HexToAddress(tokenAddress), client)
	if err != nil {
		fmt.Println("token.NewToken", err.Error())
		return "", err
	}

	balance, err := tokenContract.BalanceOf(&bind.CallOpts{}, common.HexToAddress(fromAddress.Hex()))
	if err != nil {
		fmt.Printf("Failed to check balance: %v\n", err.Error())
		return "", err
	}

	fmt.Println(balance)

	// 调用代币合约的 Transfer 方法
	tx, err := tokenContract.Transfer(txOpts, common.HexToAddress(toAddress), amount)
	if err != nil {
		fmt.Println("tokenContract.Transfer", err.Error())
		return "", err
	}

	fmt.Printf("Transaction sent: %s", tx.Hash().Hex())
	return tx.Hash().Hex(), nil
}

// 验证交易哈希是否有效
func (w *Wallet) validTxHash(txHash string) bool {
	return strings.HasPrefix(txHash, "0x") && len(txHash) == 66
}

// 查询交易结果
func (w *Wallet) TransactionReceipt(client *ethclient.Client, txHash string) (bool, error) {
	// 检查 txHash 是否为有效的哈希字符串
	if !w.validTxHash(txHash) {
		return false, errors.New("invalid transaction hash")
	}

	// 转换为以太坊的交易哈希类型
	hash := common.HexToHash(txHash)

	// 查询交易收据
	receipt, err := client.TransactionReceipt(context.Background(), hash)
	if err != nil {
		fmt.Println("Error retrieving transaction receipt:", err.Error())
		return false, err
	}

	// 检查交易状态
	if receipt.Status == 1 {
		fmt.Println("Transaction successful")
		return true, nil
	} else {
		fmt.Println("Transaction failed")
		return false, nil
	}
}
