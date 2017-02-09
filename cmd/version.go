package cmd

import (
	"fmt"

	"github.com/leonjza/go-observe/utils"
	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show the version of go-observe",
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Printf("%s\n", utils.Version)
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
