// Copyright (c) 2018 IoTeX
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/iotexproject/iotex-core/crypto"
	"github.com/iotexproject/iotex-core/logger"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate [# number]",
	Short: "Generates n number of iotex address key pairs.",
	Long:  `Generates n number of iotex address key pairs.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(generate(args))
	},
}

var _addrNum int

func generate(args []string) string {
	items := make([]string, _addrNum)
	for i := 0; i < _addrNum; i++ {
		public, private, err := crypto.EC283.NewKeyPair()
		if err != nil {
			logger.Fatal().Err(err).Msg("failed to create key pair")
		}
		items[i] = fmt.Sprintf(
			"{\"PublicKey\": \"%x\", \"PrivateKey\": \"%x\"}",
			private,
			public,
		)
	}
	return "[" + strings.Join(items, ",") + "]"
}

func init() {
	generateCmd.Flags().IntVarP(&_addrNum, "number", "n", 10, "number of addresses to be generated")
	rootCmd.AddCommand(generateCmd)
}
