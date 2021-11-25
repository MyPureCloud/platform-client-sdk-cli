package analytics_conversations_details

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_conversations_details_query"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_conversations_details_properties"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_conversations_details_jobs"
)

func init() {
	analytics_conversations_detailsCmd.AddCommand(analytics_conversations_details_query.Cmdanalytics_conversations_details_query())
	analytics_conversations_detailsCmd.AddCommand(analytics_conversations_details_properties.Cmdanalytics_conversations_details_properties())
	analytics_conversations_detailsCmd.AddCommand(analytics_conversations_details_jobs.Cmdanalytics_conversations_details_jobs())
	analytics_conversations_detailsCmd.Short = utils.GenerateCustomDescription(analytics_conversations_detailsCmd.Short, analytics_conversations_details_query.Description, analytics_conversations_details_properties.Description, analytics_conversations_details_jobs.Description, )
	analytics_conversations_detailsCmd.Long = analytics_conversations_detailsCmd.Short
}
