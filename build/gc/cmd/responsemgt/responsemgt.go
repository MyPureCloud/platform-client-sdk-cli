package responsemgt

import (
	"github.com/spf13/cobra"
)

var responsemgtCmd = &cobra.Command{
	Use:   "responsemgt",
	Short: "Manages Genesys Cloud responsemgt",
	Long:  `Manages Genesys Cloud responsemgt`,
}

func Cmdresponsemgt() *cobra.Command {
	return responsemgtCmd
}