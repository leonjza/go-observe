package cmd

import (
	"fmt"

	"github.com/leonjza/go-observe/observatory"
	"github.com/leonjza/go-observe/utils"
	"github.com/spf13/cobra"
)

// fileresultCmd represents the fileresult command
var fileresultCmd = &cobra.Command{
	Use:   "fileresult",
	Short: "Get scan results for URLs / Hostnames in a file",
	Long: `Gets scan results from the Observatory for URLs / Hostnames in
a file. A file should have newline seperated entries to be read in.

Examples:
	go-observe fileresult hostnames`,

	Run: func(cmd *cobra.Command, args []string) {

		if len(args) <= 0 {
			fmt.Println("Must provide a filename for results")
			return
		}

		hosts, err := utils.ParseHostsFile(args[0])
		if err != nil {
			fmt.Printf("Failed to parse hosts from file with error: %s\n", err)
			return
		}

		for _, host := range hosts {

			result, err := observatory.GetObservatoryResults(host)
			if err != nil {
				fmt.Printf("Failed to get results for %s with error: %s\n", host, err)
				continue
			}

			fmt.Printf("[%s] Scan State: 		%s\n", host, result.State)
			fmt.Printf("[%s] Scan Start Time: 	%s\n", host, result.StartTime)

			if result.State == "FINISHED" {
				fmt.Printf("[%s] Grade:			%s\n", host, result.Grade)
				fmt.Printf("[%s] Score:			%d\n", host, result.Score)
				fmt.Printf("[%s] Failed/Passed/Total: 	%d/%d/%d\n", host,
					result.TestsFailed, result.TestsPassed, result.TestsQuantity)
				fmt.Printf("[%s] Scan ID:		%d\n", host, result.ScanID)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(fileresultCmd)
}
