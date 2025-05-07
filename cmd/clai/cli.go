package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/v4nderstruck/clai/cmd/clai/cmdler"
)

const (
	appVersion = "0.0.1"
	appRepo    = "github.com/v4nderstruck/clai"
)

func printVersion() {
	fmt.Printf("clai version %s, %s\n", appVersion, appRepo)
}

var rootCmd = &cobra.Command{
	Use:   "clai",
	Short: "clai is a cli tool to prompt AI models",
  Args: cobra.MinimumNArgs(1),
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		getVersionFlag, err := cmd.Flags().GetBool("version")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to access version flag: %v", err)
			os.Exit(1)
		}
		if getVersionFlag {
			printVersion()
			os.Exit(0)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {},
}

var rootCmdVersionFlag bool

func main() {
	cmdler.Init()
	rootCmd.PersistentFlags().BoolVarP(&rootCmdVersionFlag, "version", "v", false, "Prints version")
	rootCmd.AddCommand(cmdler.Cmdler)
	rootCmd.Execute()
}
