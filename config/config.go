package config

import (
	"github.com/Colin1989/goactor/constant"
)

type AppConfig struct {
	NodeId   string
	NodeType string
}

func (a AppConfig) String() string {
	return a.NodeType + "_" + a.NodeId
}

func NewDefaultAppConfig() AppConfig {
	return AppConfig{
		NodeId:   constant.DefaultNodeId,
		NodeType: constant.DefaultNodeType,
	}
}
