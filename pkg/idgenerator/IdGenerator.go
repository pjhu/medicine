package idgenerator

import (
	"sync"

	"github.com/bwmarrin/snowflake"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

var (
	once   sync.Once
	node *snowflake.Node
)

func init()  {
	once.Do(func() {
		var err error
		node, err = snowflake.NewNode(1)
		if err != nil {
			logrus.Error(errors.Wrap(err, "fail to id generator"))
			panic("fail to id generator")
		}
	})
}

// NewId for create id
func NewId() int64 {
	return node.Generate().Int64()
}