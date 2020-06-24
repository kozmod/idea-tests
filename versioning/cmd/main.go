package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/kozmod/idea-tests/versioning/version"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "versioning application",
	Short: "test of use versioning + cobra",
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "version of app (+ golang version and build time)",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(fmt.Sprintf("Go Version:         %s", runtime.Version()))
		fmt.Println(fmt.Sprintf("Build Version:      %s", version.Version))
		fmt.Println(fmt.Sprintf("Build Version Time: %s", version.VersionTime))
	},
}

func main() {
	rootCmd.AddCommand(versionCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
