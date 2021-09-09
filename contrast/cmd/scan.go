package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

/**

TODO: Run a SAST scan locally

For example:

> contrast scan
(Detects project source(s) automatically and runs a scan)

> contrast scan ../custom/path
(Same as above, just with a specified path)

Ideally this is a local scan with no tokens. However, an error message could
print that says that a token is required. I'm inclined to think that there should
be an -u/--upload or -r/--remote flag that is required in order to upload code to the backend. This
could allow us to run local scans and/or remote scans if customers want to use our services to burden
the local machine less? This is certainly simpler if it is a local-only command

Adding a token can allow the results to upload back to the security dashboard if desired. The token
should be read from ENV or from the .contrast config file

Possible flags:

-v, --verbose
-p, --proxy
-s, --silent
-r, --remote (to perform a remote scan rather than local)

**/

// scanCmd represents the scan command
var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Scan custom code to detect vulnerabilities",
	Long:  `Run a static application (SAST) scan and report detected vulnerabilities`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: How to document optional `path` argument?
		if len(args) > 0 {
			path = args[0]
			fmt.Println("TODO: Run scan in relative path \"" + path + "\"")
		} else {
			fmt.Println("TODO: Run scan in current directory")
		}
	},
}

func init() {
	rootCmd.AddCommand(scanCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// scanCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// scanCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
