package utils

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
)

var TxWallet *hdwallet.Wallet

func UpdateTxWallet(mnemonic string) error {
	var err error
	// 生成 txwallet 失败的原因，添加助记词时需要检查这个助记词是否存在问题
	TxWallet, err = hdwallet.NewFromMnemonic(mnemonic)
	return err
}

func DerivationPath(userUID string) accounts.DerivationPath {
	return hdwallet.MustParseDerivationPath(fmt.Sprintf("%s/%s", Config.Mnemonic.Path, userUID))
}

func ValidMnemonic(mnemonic string) bool {
	_, err := hdwallet.NewFromMnemonic(mnemonic)
	return err == nil
}

func GenerateAddress(mn string, userUID string) (string, string, error) {
	account, err := TxWallet.Derive(DerivationPath(userUID), false)
	if err != nil {
		return "", "", err
	}
	privateKey, err := TxWallet.PrivateKeyHex(account)
	if err != nil {
		return "", "", err
	}
	return account.Address.Hex(), privateKey, nil
}
