package notifications

import (
	"github.com/spf13/cobra"
)

var notificationsCmd = &cobra.Command{
	Use:   "notifications",
	Short: "Manages Genesys Cloud notifications",
	Long:  `Manages Genesys Cloud notifications`,
}

func Cmdnotifications() *cobra.Command {
	return notificationsCmd
}
