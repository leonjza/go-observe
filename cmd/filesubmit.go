package cmd

import (
	"fmt"

	"github.com/leonjza/go-observe/observatory"
	"github.com/leonjza/go-observe/utils"
	"github.com/spf13/cobra"
)

// flag variables are taken from `submit.go` ;)

// filesubmitCmd represents the filesubmit command
var filesubmitCmd = &cobra.Command{
	Use:   "filesubmit",
	Short: "Submit URLs / Hostnames from a file to the Observatory for analysis",
	Long: `Submits URLs / Hostnames from a file to the Observatory for analysis. A
file should have newline seperated entries to be read in.

Examples:
	go-observe filesubmit hostnames
	go-observe filesubmit hostnames --rescan
	go-observe filesubmit hostnames --rescan --no-hide`,

	Run: func(cmd *cobra.Command, args []string) {

		if len(args) <= 0 {
			fmt.Println("Must provide a filename to submit")
			return
		}

		// the slice of hosts we are going to submit
		hosts, err := utils.ParseHostsFile(args[0])
		if err != nil {
			fmt.Printf("Failed to parse the file %s with error: %s\n", args[0], err)
			return
		}

		fmt.Printf("Submitting %d hosts parsed from file to the Observatory\n", len(hosts))

		for _, host := range hosts {

			result, err := observatory.SubmitObservatoryAnalysis(host, noHidden, forceRescan)
			if err != nil {
				fmt.Printf("Failed to submit host %s with error: %s\n", host, err)
				continue
			}

			fmt.Printf("Submitted %s. Observatory status: %s\n", host, result.State)
		}
	},
}

func init() {
	RootCmd.AddCommand(filesubmitCmd)

	filesubmitCmd.Flags().BoolVarP(&forceRescan, "rescan", "r", false, "Force a rescan from the Observatory")
	filesubmitCmd.Flags().BoolVarP(&noHidden, "no-hide", "n", false, "Remove the flag that hides scan results from the Observatory")
}
