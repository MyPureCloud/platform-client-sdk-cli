package analytics_users_details_jobs

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_users_details_jobs_results"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_users_details_jobs_availability"
)

func init() {
	analytics_users_details_jobsCmd.AddCommand(analytics_users_details_jobs_results.Cmdanalytics_users_details_jobs_results())
	analytics_users_details_jobsCmd.AddCommand(analytics_users_details_jobs_availability.Cmdanalytics_users_details_jobs_availability())
	analytics_users_details_jobsCmd.Short = utils.GenerateCustomDescription(analytics_users_details_jobsCmd.Short, analytics_users_details_jobs_results.Description, analytics_users_details_jobs_availability.Description, )
	analytics_users_details_jobsCmd.Long = analytics_users_details_jobsCmd.Short
}
