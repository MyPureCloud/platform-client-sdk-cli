package notifications

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/channels"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/topics"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/subscriptions"

	"github.com/spf13/cobra"
)

var notificationsCmd = &cobra.Command{
	Use:   "notifications",
	Short: "Manages Genesys Cloud notifications",
	Long:  `Manages Genesys Cloud notifications`,
}

func Cmdnotifications() *cobra.Command {
	notificationsCmd.AddCommand(topics.Cmdtopics())
	notificationsCmd.AddCommand(channels.Cmdchannels())
	notificationsCmd.AddCommand(subscriptions.Cmdsubscriptions())

	return notificationsCmd
}
