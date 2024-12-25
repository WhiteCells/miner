package utils

import "github.com/bwmarrin/snowflake"

func GenUID() (string, error) {
	node, err := snowflake.NewNode(1) // 1 表示分片标识
	if err != nil {
		return "", err
	}
	id := node.Generate()
	return id.String(), nil
}
