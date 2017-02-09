package cmd

import (
	"fmt"

	"github.com/leonjza/go-observe/observatory"
	"github.com/leonjza/go-observe/utils"
	"github.com/spf13/cobra"
)

// variables for flags
var (
	noHidden    bool
	forceRescan bool
)

// submitCmd represents the submit command
var submitCmd = &cobra.Command{
	Use:   "submit [url / hostname to submit]",
	Short: "Submit a URL / Hostname to the Observatory for analysis",
	Long: `Submits a URL or a Hostname to the Observatory for analysis.

If a URL / Hostname has previously been submitted to the Observatory,
then the API will respond with a "FINISHED". A rescan can be forced with
the --rescan flag.

Examples:
	go-observe submit https://duckduckgo.com
	go-observe submit duckduckgo.com --rescan
	go-observe submit duckduckgo.com --rescan -n`,

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

		// Submit the analysis
		response, err := observatory.SubmitObservatoryAnalysis(target, noHidden, forceRescan)
		if err != nil {
			fmt.Printf("Failed to send host for analysis for error: %s\n", err)
			return
		}

		// check the response
		if response.State == "FINISHED" {
			fmt.Printf("Scan has already been finished on %s.\n", response.EndTime)
			fmt.Printf("Use `go-observe results %s` to get scan results.\n", target)
		} else {
			fmt.Printf("Scan has been submitted and now has status: %s\n", response.State)
		}
	},
}

func init() {
	RootCmd.AddCommand(submitCmd)

	submitCmd.Flags().BoolVarP(&forceRescan, "rescan", "r", false, "Force a rescan from the Observatory")
	submitCmd.Flags().BoolVarP(&noHidden, "no-hide", "n", false, "Remove the flag that hides scan results from the Observatory")
}
