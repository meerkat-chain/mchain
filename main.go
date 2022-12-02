package main

import (
	_ "embed"

	"github.com/meerkat-chain/mchain/command/root"
	"github.com/meerkat-chain/mchain/licenses"
)

var (
	//go:embed LICENSE
	license string
)

func main() {
	licenses.SetLicense(license)

	root.NewRootCommand().Execute()
}
