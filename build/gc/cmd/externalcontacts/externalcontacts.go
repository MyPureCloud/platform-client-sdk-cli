package externalcontacts

import (
	"github.com/spf13/cobra"
)

var externalcontactsCmd = &cobra.Command{
	Use:   "externalcontacts",
	Short: "Manages Genesys Cloud external contacts",
	Long:  `Manages Genesys Cloud external contacts`,
}

func Cmdexternalcontacts() *cobra.Command {
	return externalcontactsCmd
}