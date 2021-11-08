package idgenerator

import (
	log "github.com/sirupsen/logrus"
	"github.com/bwmarrin/snowflake"
	"github.com/pkg/errors"
)

var node *snowflake.Node

func init()  {

	var err error
	node, err = snowflake.NewNode(1)
	if err != nil {
		log.Error(errors.Wrap(err, "fail to id generator"))
		panic("fail to id generator")
	}
}

// NewId for create id
func NewId() (int64) {
	return node.Generate().Int64()
}