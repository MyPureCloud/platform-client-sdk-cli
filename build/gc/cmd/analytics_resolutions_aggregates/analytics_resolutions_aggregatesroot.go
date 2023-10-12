package analytics_resolutions_aggregates

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_resolutions_aggregates_jobs"
)

func init() {
	analytics_resolutions_aggregatesCmd.AddCommand(analytics_resolutions_aggregates_jobs.Cmdanalytics_resolutions_aggregates_jobs())
	analytics_resolutions_aggregatesCmd.Short = utils.GenerateCustomDescription(analytics_resolutions_aggregatesCmd.Short, analytics_resolutions_aggregates_jobs.Description, )
	analytics_resolutions_aggregatesCmd.Long = analytics_resolutions_aggregatesCmd.Short
}
