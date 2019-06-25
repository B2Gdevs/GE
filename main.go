package main

import (
	"fmt"
	"github.com/b2gdevs/ge/lib"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = cobra.Command{
		Use:   "uninstall",
		Short: "uninstall a go package",
		Long:  `Uninstall a go package from bin src, and pkg folders`,
		Run: func(cmd *cobra.Command, args []string) {
			ge.RemoveFiles(args)
		},
	}

	if len(os.Args) == 1 {
		fmt.Print("Error: No arguments supplied")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "uninstall":
		ge.Uninstall(rootCmd)
	case "help":
		ge.Help()
	default:
		ge.ExecuteGoCmd()
	}

}
