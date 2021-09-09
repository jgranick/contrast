package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

/**

TODO: Run an OSS scan

For example:

> contrast audit
(Detects project type(s) automatically and prints vulnerabilities)

> contrast audit ../custom/path
(Same as above, just with a specified path)

The results should be local and require no tokens. Adding a token can allow
the results to upload back to the security dashboard if desired. The token
should be read from ENV or from the .contrast config file

RE: Existing contrast-cli behavior

Cataloging should be implied by running this command
Language should be detected automatically, and the command run (and aggregating) all detected languages
Backend identifiers should be implied by the (optional!) token
Reporting should be clear and concise, --verbose for wordier response or --silent for no standard output
I think all levels of vulnerability should be reported
Rather than breaking on a threshold, I think there should be a way to ignore a vulnerability (contrast ignore command?)

Possible flags:

-v, --verbose
-p, --proxy
-s, --silent


>

--yaml_path string Used only if you want to run the command with a yaml
--api_key string (required): An agent API key as provided by Contrast UI
--authorization string (required): An agent Authorization credentials as provided by Contrast UI
--organization_id string (required): The ID of your organization in Contrast UI
--application_id string (required): The ID of the application cataloged by Contrast UI
--host string (required): Provide the name of the host and optionally the port expressed as <host>:<port>.
--application_name string (optional): The name of the application cataloged by Contrast UI
--catalogue_application (required for catalogue): Provide this if you want to catalogue an application
--language string (required for catalogue): Valid values are JAVA, DOTNET, NODE, PYTHON and RUBY. If there are multiple project configuration files in the project path, language is also required.
--project_path string (optional): The directory root of a project/application that you would like analyzed. Defaults to current directory.
--app_groups string (optional for catalogue): Assign your application to one or more pre-existing groups when using the catalogue command. Group lists should be comma separated.
--maven_settings_path string (optional): Allows you to specify an alternative location for your maven settings.xml file.
--proxy string (optional): Allows for connection via a proxy server. If authentication is required please provide the username and password with the protocol, host and port. For instance: 'http://username:password@:'.
--silent (optional): Silences JSON output.
-v, --version Displays CLI Version you are currently on.
--sub_project string (optional): Specify the sub project within your gradle application.
-h, --help Display usage guide.
-r, --report Display vulnerability information for this application.
-f, --fail Set the process to fail if this option is set in combination with the --report and --cveseverity.
-s, --cve_severity _type Combined with the --report command, allows the user to report libraries with vulnerabilities above a chosen severity level. For example, cveseverity=medium only reports libraries with vulnerabilities at medium or higher severity. Values for level are high, medium or low.
--cve_threshold _number The number of CVE's that must be exceeded to fail a build

**/

// auditCmd represents the audit command
var auditCmd = &cobra.Command{
	Use:   "audit",
	Short: "Audit project dependencies for known vulnerabilities",
	Long:  `Run a security audit for known vulnerabilities`,
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: How to document optional `path` argument?
		if len(args) > 0 {
			path = args[0]
			fmt.Println("TODO: Run OSS audit in relative path=\"" + path + "\"")
		} else {
			fmt.Println("TODO: Run OSS audit in current directory")
		}
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
