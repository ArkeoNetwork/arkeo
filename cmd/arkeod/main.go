//go:build !regtest
// +build !regtest

package main

import (
	"os"

	"github.com/arkeonetwork/arkeo/app"
	"github.com/arkeonetwork/arkeo/arkeocli"

	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
)

func main() {

	rootCmd, _ := NewRootCmd()
	// rootCmd, _ := NewRootCmd(
	// 	app.Name,
	// 	app.AccountAddressPrefix,
	// 	app.DefaultNodeHome,
	// 	xstrings.NoDash(app.Name),
	// 	app.ModuleBasics,
	// 	app.New,
	// 	// this line is used by starport scaffolding # root/arguments
	// )
	// add in arkeo specific utilities
	rootCmd.AddCommand(arkeocli.GetArkeoCmd())
	if err := svrcmd.Execute(rootCmd, "", app.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}
