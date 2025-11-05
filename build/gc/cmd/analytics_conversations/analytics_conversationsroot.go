package analytics_conversations

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_conversations_aggregates"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_conversations_activity"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_conversations_details"
)

func init() {
	analytics_conversationsCmd.AddCommand(analytics_conversations_aggregates.Cmdanalytics_conversations_aggregates())
	analytics_conversationsCmd.AddCommand(analytics_conversations_activity.Cmdanalytics_conversations_activity())
	analytics_conversationsCmd.AddCommand(analytics_conversations_details.Cmdanalytics_conversations_details())
	analytics_conversationsCmd.Short = utils.GenerateCustomDescription(analytics_conversationsCmd.Short, analytics_conversations_aggregates.Description, analytics_conversations_activity.Description, analytics_conversations_details.Description, )
	analytics_conversationsCmd.Long = analytics_conversationsCmd.Short
}
