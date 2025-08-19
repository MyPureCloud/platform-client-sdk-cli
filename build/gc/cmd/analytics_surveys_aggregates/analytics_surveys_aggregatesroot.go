package analytics_surveys_aggregates

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_surveys_aggregates_query"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_surveys_aggregates_jobs"
)

func init() {
	analytics_surveys_aggregatesCmd.AddCommand(analytics_surveys_aggregates_query.Cmdanalytics_surveys_aggregates_query())
	analytics_surveys_aggregatesCmd.AddCommand(analytics_surveys_aggregates_jobs.Cmdanalytics_surveys_aggregates_jobs())
	analytics_surveys_aggregatesCmd.Short = utils.GenerateCustomDescription(analytics_surveys_aggregatesCmd.Short, analytics_surveys_aggregates_query.Description, analytics_surveys_aggregates_jobs.Description, )
	analytics_surveys_aggregatesCmd.Long = analytics_surveys_aggregatesCmd.Short
}
