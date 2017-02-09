package cmd

import (
	"fmt"

	"github.com/leonjza/go-observe/observatory"
	"github.com/leonjza/go-observe/utils"
	"github.com/spf13/cobra"
)

// variables for flags
var (
	details bool
)

// resultCmd represents the result command
var resultCmd = &cobra.Command{
	Use:   "result [url / hostname]",
	Short: "Retrieve results for a URL / Hostname from the Observatory",
	Long: `Retreives scan results for a URL from the Observatory.

When the --detail flag is provided, a list of the findings details will also
be retreived.`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) <= 0 {
			fmt.Println("Must provide a URL / Hostname to submit")
			return
		}

		// validate the url/host
		target, err := utils.ValidateAndGetURLHost(args[0])
		if err != nil {
			fmt.Printf("Failed to parse url/host with error: %s\n", err)
		}

		response, err := observatory.GetObservatoryResults(target)
		if err != nil {
			fmt.Printf("Failed to get results with error: %s\n", err)
			return
		}
	},
}

func init() {
	RootCmd.AddCommand(resultCmd)

	submitCmd.Flags().BoolVarP(&details, "detail", "d", false, "Get issue details too")
}
