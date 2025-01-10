#
go install github.com/ethereum/go-ethereum/cmd/abigen@latest
~/go/bin/abigen --abi=./usdt_token_abi.json --pkg=token --out=token.go
