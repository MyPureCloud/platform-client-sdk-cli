package analytics_conversations_activity

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_conversations_activity_query"
)

func init() {
	analytics_conversations_activityCmd.AddCommand(analytics_conversations_activity_query.Cmdanalytics_conversations_activity_query())
	analytics_conversations_activityCmd.Short = utils.GenerateCustomDescription(analytics_conversations_activityCmd.Short, analytics_conversations_activity_query.Description, )
	analytics_conversations_activityCmd.Long = analytics_conversations_activityCmd.Short
}
