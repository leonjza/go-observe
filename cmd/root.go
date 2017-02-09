package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var cfgFile string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "go-observe",
	Short: "A command line Mozilla Observatory client",
	Long: `
                      __
  ___ ____  _______  / /  ___ ___ _____  _____
 / _  / _ \/___/ _ \/ _ \(_-</ -_) __/ |/ / -_)
 \_, /\___/    \___/_.__/___/\__/_/  |___/\__/
/___/ @leonjza`,
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
