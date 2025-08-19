package analytics_users_details

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_users_details_query"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_users_details_jobs"
)

func init() {
	analytics_users_detailsCmd.AddCommand(analytics_users_details_query.Cmdanalytics_users_details_query())
	analytics_users_detailsCmd.AddCommand(analytics_users_details_jobs.Cmdanalytics_users_details_jobs())
	analytics_users_detailsCmd.Short = utils.GenerateCustomDescription(analytics_users_detailsCmd.Short, analytics_users_details_query.Description, analytics_users_details_jobs.Description, )
	analytics_users_detailsCmd.Long = analytics_users_detailsCmd.Short
}
