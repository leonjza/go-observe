package cmd

import (
	"fmt"

	"github.com/leonjza/go-observe/observatory"
	"github.com/leonjza/go-observe/utils"
	"github.com/spf13/cobra"
)

// variables for flags
var details bool

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
			return
		}

		result, err := observatory.GetObservatoryResults(target)
		if err != nil {
			fmt.Printf("Failed to get results with error: %s\n", err)
			return
		}

		// echo the results
		fmt.Println("")
		fmt.Printf("Scan State: 		%s\n", result.State)
		fmt.Printf("Scan Start Time: 	%s\n", result.StartTime)

		if result.State == "FINISHED" {
			fmt.Printf("Grade:			%s\n", result.Grade)
			fmt.Printf("Score:			%d\n", result.Score)
			fmt.Printf("Failed/Passed/Total: 	%d/%d/%d\n", result.TestsFailed, result.TestsPassed, result.TestsQuantity)
			fmt.Printf("Scan ID:		%d\n", result.ScanID)
		}

		// if we need the details, get that too!
		if details && result.State == "FINISHED" {

			details, err := observatory.GetObservatoryDetails(result.ScanID)
			if err != nil {
				fmt.Printf("An error occured when trying to get the details: %s\n", err)
				return
			}

			fmt.Println("")
			for _, v := range details {

				fmt.Printf("Pass: %t	Score: %d	Description: %s\n", v.Pass, v.ScoreModifier, v.ScoreDescription)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(resultCmd)

	resultCmd.Flags().BoolVarP(&details, "detail", "d", false, "Get issue details too")
}
