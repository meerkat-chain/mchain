package root

import (
	"fmt"
	"os"

	"github.com/meerkat-chain/mchain/command/backup"
	"github.com/meerkat-chain/mchain/command/genesis"
	"github.com/meerkat-chain/mchain/command/helper"
	"github.com/meerkat-chain/mchain/command/ibft"
	"github.com/meerkat-chain/mchain/command/license"
	"github.com/meerkat-chain/mchain/command/loadbot"
	"github.com/meerkat-chain/mchain/command/monitor"
	"github.com/meerkat-chain/mchain/command/peers"
	"github.com/meerkat-chain/mchain/command/secrets"
	"github.com/meerkat-chain/mchain/command/server"
	"github.com/meerkat-chain/mchain/command/status"
	"github.com/meerkat-chain/mchain/command/txpool"
	"github.com/meerkat-chain/mchain/command/version"
	"github.com/meerkat-chain/mchain/command/whitelist"
	"github.com/spf13/cobra"
)

type RootCommand struct {
	baseCmd *cobra.Command
}

func NewRootCommand() *RootCommand {
	rootCommand := &RootCommand{
		baseCmd: &cobra.Command{
			Short: "Meerkat Chain is a framework for building Ethereum-compatible Blockchain networks",
		},
	}

	helper.RegisterJSONOutputFlag(rootCommand.baseCmd)

	rootCommand.registerSubCommands()

	return rootCommand
}

func (rc *RootCommand) registerSubCommands() {
	rc.baseCmd.AddCommand(
		version.GetCommand(),
		txpool.GetCommand(),
		status.GetCommand(),
		secrets.GetCommand(),
		peers.GetCommand(),
		monitor.GetCommand(),
		loadbot.GetCommand(),
		ibft.GetCommand(),
		backup.GetCommand(),
		genesis.GetCommand(),
		server.GetCommand(),
		whitelist.GetCommand(),
		license.GetCommand(),
	)
}

func (rc *RootCommand) Execute() {
	if err := rc.baseCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)

		os.Exit(1)
	}
}
