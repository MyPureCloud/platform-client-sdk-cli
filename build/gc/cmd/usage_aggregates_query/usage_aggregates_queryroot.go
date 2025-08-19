package usage_aggregates_query

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/usage_aggregates_query_jobs"
)

func init() {
	usage_aggregates_queryCmd.AddCommand(usage_aggregates_query_jobs.Cmdusage_aggregates_query_jobs())
	usage_aggregates_queryCmd.Short = utils.GenerateCustomDescription(usage_aggregates_queryCmd.Short, usage_aggregates_query_jobs.Description, )
	usage_aggregates_queryCmd.Long = usage_aggregates_queryCmd.Short
}
