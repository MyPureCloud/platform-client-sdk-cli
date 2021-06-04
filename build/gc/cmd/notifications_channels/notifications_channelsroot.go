package notifications_channels

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/notifications_channels_subscriptions"
)

func init() {
	notifications_channelsCmd.AddCommand(notifications_channels_subscriptions.Cmdnotifications_channels_subscriptions())
	notifications_channelsCmd.Short = utils.GenerateCustomDescription(notifications_channelsCmd.Short, notifications_channels_subscriptions.Description, )
	notifications_channelsCmd.Long = notifications_channelsCmd.Short
}
