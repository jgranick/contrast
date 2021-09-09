package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// auditCmd represents the audit command
var auditCmd = &cobra.Command{
	Use:   "audit",
	Short: "Audit project dependencies for known vulnerabilities",
	Long:  `Run a security audit for known vulnerabilities`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("audit called")
	},
}

func init() {
	rootCmd.AddCommand(auditCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// auditCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// auditCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
