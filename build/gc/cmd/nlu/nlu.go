
package nlu

import (
	"github.com/spf13/cobra"
)

var nluCmd = &cobra.Command{
	Use:   "nlu",
	Short: "Manages Genesys Cloud nlu",
	Long:  `Manages Genesys Cloud nlu`,
}

func Cmdnlu() *cobra.Command {
	return nluCmd
}
