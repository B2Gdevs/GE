package ge

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"

	"github.com/spf13/cobra"
)

// RemoveFiles removes all files found with the package that was installed
// with go get [pkgname] or ge get [pkgname].  Those files are found in the
// bin, src, and pkg folders. The removal from the pkg folder only works on
// windows.
func RemoveFiles(cmd *cobra.Command, args []string) {
	flag, _ := cmd.LocalFlags().GetBool("parent")
	dir := args[0]
	if flag {
		dir = path.Dir(args[0])
	}
	srcPath := path.Join("src", dir)
	// Windows only
	pkgPath := path.Join("pkg", "windows_amd64", path.Dir(dir))

	// go clean needs to have the src folder existing to remove the bin
	// in the bin folder.
	goCmd := exec.Command("go", "clean", "-i", dir)

	if err := goCmd.Run(); err != nil {
		log.Fatal(err)
	}

	if _, err := os.Stat(srcPath); err == nil {

		cmd := exec.Command("rm", "-rf", srcPath)
		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println(srcPath)

	if _, err := os.Stat(pkgPath); err == nil {
		cmd := exec.Command("rm", "-rf", pkgPath)
		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}
	}
}

// ExecuteGoCmd executes any go commands that would normally be given to
// the go cli tool.
func ExecuteGoCmd() {
	out, err := exec.Command("go", os.Args[1:]...).Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(string(out))
}

// Uninstall executes the root command which
// executes the RemoveFiles function.
func Uninstall() *cobra.Command {
	return &cobra.Command{
		Use:   "uninstall",
		Short: "uninstall a go package",
		Long:  `Uninstall a go package from bin src, and pkg folders`,
		Run: func(cmd *cobra.Command, args []string) {
			RemoveFiles(cmd, args)
		},
	}
}

// Help prints out a message that explains what Go Extension does and
// the help message that the Go cli tools prints out.
func Help() {

	str := "\nGE, Go Extension, is a wrapper for the Go cli tool.\n\n" +
		"The only command this wrapper has is uninstall.\n\n" +

		"The uninstall command will remove the binary files, src files,\n" +
		"and the pkg files for the pkg that was installed with go get \n" +
		"or ge get commands.\n\n" +

		"Usage: ge uninstall [package name]\n\n" +

		"The rest that follows is the output from the Go cli.\n\n" +
		"Usage: ge [cmd] [args]\n\n"

	out, err := exec.Command("go", "help").Output()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf(str)
	fmt.Println(string(out))
}