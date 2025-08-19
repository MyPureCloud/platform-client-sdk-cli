package analytics_users_activity

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_users_activity_query"
)

func init() {
	analytics_users_activityCmd.AddCommand(analytics_users_activity_query.Cmdanalytics_users_activity_query())
	analytics_users_activityCmd.Short = utils.GenerateCustomDescription(analytics_users_activityCmd.Short, analytics_users_activity_query.Description, )
	analytics_users_activityCmd.Long = analytics_users_activityCmd.Short
}
