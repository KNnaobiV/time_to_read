package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var Version = "development"

var rootCmd = &cobra.Command{
	Use: "dunn",
	Short: "Dunn is a file read time checker",
	Long: `A read time checker for text files built with Go.
			Complete documentation is available at 
			https://github.com/knnaobiv/time_to_read`,
	Example: "dunn check read",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

var versionCmd = &cobra.Command{
	Use: "version",
	Short: "Prints the version number of Dunn to the CLI",
	Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Dunn CLI version: %s\n", Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

func Execute() {
	// cobra.CheckErr(rootCmd.Execute())
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}