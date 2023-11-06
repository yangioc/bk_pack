package util

import (
	"github.com/bwmarrin/snowflake"
)

var _instants_snowflake *snowflake.Node

func GenStrUUID(serviceNodeNum int64) string {
	if _instants_snowflake == nil {
		var err error
		_instants_snowflake, err = snowflake.NewNode(serviceNodeNum)
		if err != nil {
			panic(err)
		}
	}
	return _instants_snowflake.Generate().String()
}
