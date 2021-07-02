package notifications

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/notifications_channels"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/notifications_availabletopics"
)

func init() {
	notificationsCmd.AddCommand(notifications_channels.Cmdnotifications_channels())
	notificationsCmd.AddCommand(notifications_availabletopics.Cmdnotifications_availabletopics())
	notificationsCmd.Short = utils.GenerateCustomDescription(notificationsCmd.Short, notifications_channels.Description, notifications_availabletopics.Description, )
	notificationsCmd.Long = notificationsCmd.Short
}
