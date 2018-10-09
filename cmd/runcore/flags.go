package main

import (
	"srcd/cmd/utils"

	"gopkg.in/urfave/cli.v1"
)

var (
	// flags that configure the node
	nodeFlags = []cli.Flag{
		utils.IdentityFlag,
		utils.DataDirFlag,
		configFileFlag,
	}

	rpcFlags = []cli.Flag{}
)