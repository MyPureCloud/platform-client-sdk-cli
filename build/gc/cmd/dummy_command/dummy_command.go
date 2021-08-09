package dummy_command

import (
	"fmt"
	"github.com/spf13/cobra"
)

var dummyCmd = &cobra.Command{
	Use:   "dummy_command",
	Short: "Dummy command that does nothing",
	Long:  `Dummy command that does nothing`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Running dummy command...")
	},
}

func Cmddummy_command() *cobra.Command {
	return dummyCmd
}
