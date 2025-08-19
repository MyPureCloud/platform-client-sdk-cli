package notifications

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/notifications_availabletopics"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/notifications_channels"
)

func init() {
	notificationsCmd.AddCommand(notifications_availabletopics.Cmdnotifications_availabletopics())
	notificationsCmd.AddCommand(notifications_channels.Cmdnotifications_channels())
	notificationsCmd.Short = utils.GenerateCustomDescription(notificationsCmd.Short, notifications_availabletopics.Description, notifications_channels.Description, )
	notificationsCmd.Long = notificationsCmd.Short
}
