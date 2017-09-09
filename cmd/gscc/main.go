// Copyright 2017 The go-socicoin Authors
// This file is part of go-socicoin.
//
//
package main

import (
	"fmt"
	"os"
	//"flag"

	"github.com/urfave/cli"
)

var (
	//mode *string = flag.String("mode", "apend", "open mode")
	//hex  *bool   = flag.Bool("hex", false, "hex mode")
	datadir_flag = cli.StringFlag{
		Name:  "datadir",
		Usage: "data directory",
		Value: "./",
	}
	testnet_flag = cli.BoolFlag{
		Name:  "testnet",
		Usage: "testnet for private network",
	}

	app = cli.NewApp()
)

func init() {

	app.Action = gscc
	app.HideVersion = true // we have a command to print the version
	app.Copyright = "Copyright 2017-2020 The go-socicoin Authors"

	app.Commands = []cli.Command{
		// See chaincmd.go:
		initCommand,
	}

	app.Flags = []cli.Flag{
		datadir_flag,
		testnet_flag,
	}
}

func main() {
	fmt.Println("go-socicoin starting...")
	fmt.Println(os.Args)

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func gscc(ctx *cli.Context) error {
	fmt.Println("gscc()")
	return nil
}
