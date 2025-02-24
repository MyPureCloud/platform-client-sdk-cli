package analytics_conversations_details_jobs

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_conversations_details_jobs_availability"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_conversations_details_jobs_results"
)

func init() {
	analytics_conversations_details_jobsCmd.AddCommand(analytics_conversations_details_jobs_availability.Cmdanalytics_conversations_details_jobs_availability())
	analytics_conversations_details_jobsCmd.AddCommand(analytics_conversations_details_jobs_results.Cmdanalytics_conversations_details_jobs_results())
	analytics_conversations_details_jobsCmd.Short = utils.GenerateCustomDescription(analytics_conversations_details_jobsCmd.Short, analytics_conversations_details_jobs_availability.Description, analytics_conversations_details_jobs_results.Description, )
	analytics_conversations_details_jobsCmd.Long = analytics_conversations_details_jobsCmd.Short
}
