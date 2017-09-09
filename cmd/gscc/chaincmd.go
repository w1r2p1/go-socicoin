// Copyright 2017 The go-socicoin Authors
// This file is part of go-socicoin.
//
//
package main

import (
	"fmt"

	"github.com/urfave/cli"
)

var (
	initCommand = cli.Command{
		Action:      initGenesis,
		Name:        "init",
		Usage:       "initialize a new genesis block",
		ArgsUsage:   "<genesisPath>",
		Category:    "BLOCKCHAIN COMMANDS",
		Description: `The init command initializes a new genesis block.`,
	}
)

func initGenesis(ctx *cli.Context) error {
	fmt.Println("initGenesis()")
	genesisPath := ctx.Args().First()
	fmt.Println("genesis file: ", genesisPath)
	fmt.Println("datadir :", ctx.GlobalString(datadir_flag.Name))

	fmt.Println(ctx.Args())
	return nil
}
