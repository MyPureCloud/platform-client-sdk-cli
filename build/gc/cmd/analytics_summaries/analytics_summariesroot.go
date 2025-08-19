package analytics_summaries

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_summaries_aggregates"
)

func init() {
	analytics_summariesCmd.AddCommand(analytics_summaries_aggregates.Cmdanalytics_summaries_aggregates())
	analytics_summariesCmd.Short = utils.GenerateCustomDescription(analytics_summariesCmd.Short, analytics_summaries_aggregates.Description, )
	analytics_summariesCmd.Long = analytics_summariesCmd.Short
}
