package main

import (
	"fmt"
	"github.com/b2gdevs/ge/lib"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{Use: "ge"}
	subCmd := ge.Uninstall()
	subCmd.Flags().BoolP("parent", "p", false, "remove parent direcorties of package")
	rootCmd.AddCommand(subCmd)

	if len(os.Args) == 1 {
		fmt.Print("Error: No arguments supplied")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "uninstall":
		rootCmd.Execute()
	case "help":
		ge.Help()
	default:
		ge.ExecuteGoCmd()
	}

}
