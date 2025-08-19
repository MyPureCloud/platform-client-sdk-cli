package workforcemanagement_notifications

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_notifications_update"
)

func init() {
	workforcemanagement_notificationsCmd.AddCommand(workforcemanagement_notifications_update.Cmdworkforcemanagement_notifications_update())
	workforcemanagement_notificationsCmd.Short = utils.GenerateCustomDescription(workforcemanagement_notificationsCmd.Short, workforcemanagement_notifications_update.Description, )
	workforcemanagement_notificationsCmd.Long = workforcemanagement_notificationsCmd.Short
}
