package utils

import "github.com/bwmarrin/snowflake"

func GenerateId() string{
	node, err := snowflake.NewNode(1)
	if err != nil {
		panic(err)
	}
	// Generate a snowflake ID.
	id := node.Generate()
	return id.String()
}
