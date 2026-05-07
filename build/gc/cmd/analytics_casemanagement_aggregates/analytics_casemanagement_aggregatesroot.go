package analytics_casemanagement_aggregates

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_casemanagement_aggregates_query"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_casemanagement_aggregates_jobs"
)

func init() {
	analytics_casemanagement_aggregatesCmd.AddCommand(analytics_casemanagement_aggregates_query.Cmdanalytics_casemanagement_aggregates_query())
	analytics_casemanagement_aggregatesCmd.AddCommand(analytics_casemanagement_aggregates_jobs.Cmdanalytics_casemanagement_aggregates_jobs())
	analytics_casemanagement_aggregatesCmd.Short = utils.GenerateCustomDescription(analytics_casemanagement_aggregatesCmd.Short, analytics_casemanagement_aggregates_query.Description, analytics_casemanagement_aggregates_jobs.Description, )
	analytics_casemanagement_aggregatesCmd.Long = analytics_casemanagement_aggregatesCmd.Short
}
