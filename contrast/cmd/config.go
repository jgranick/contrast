package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

/**

TODO: Allow for a simple command to print or write a token that connects to the online dashboard

For example:

> contrast config
$CONTRAST_TOKEN=dG9rZW4dG9rZW4

> contrast config dG9rZW4dG9rZW4
Writes to .contrast (creates if not present) and sets the given token value

Another idea would be to provide a wizard to generate a new token locally
(prompt for app ID, etc)

**/

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure remote security dashboard features",
	Long:  `Configure remote security dashboard features`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("TODO: Print .contrast config with no arguments, or write config with arguments")
	},
}

func init() {
	rootCmd.AddCommand(configCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
