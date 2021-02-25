
package externalcontacts

import (
	"github.com/spf13/cobra"
)

var externalcontactsCmd = &cobra.Command{
	Use:   "externalcontacts",
	Short: "Manages Genesys Cloud externalcontacts",
	Long:  `Manages Genesys Cloud externalcontacts`,
}

func Cmdexternalcontacts() *cobra.Command {
	return externalcontactsCmd
}
