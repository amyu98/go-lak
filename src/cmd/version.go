package cmd

import (
	"fmt"
	"github.com/amyu98/go-lak/src/lakversion"
	"github.com/spf13/cobra"
)

// versionCmd represents the lakversion command
var versionCmd = &cobra.Command{
	Use:   "lakversion",
	Short: "Print the lakversion number of generated code example",
	Long:  `All software has versions. This is generated code example`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Build Date:", lakversion.BuildDate)
		fmt.Println("Git Commit:", lakversion.GitCommit)
		fmt.Println("Version:", lakversion.Version)
		fmt.Println("Go Version:", lakversion.GoVersion)
		fmt.Println("OS / Arch:", lakversion.OsArch)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
