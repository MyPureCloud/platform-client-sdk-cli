package usage_client_aggregates_query

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/usage_client_aggregates_query_jobs"
)

func init() {
	usage_client_aggregates_queryCmd.AddCommand(usage_client_aggregates_query_jobs.Cmdusage_client_aggregates_query_jobs())
	usage_client_aggregates_queryCmd.Short = utils.GenerateCustomDescription(usage_client_aggregates_queryCmd.Short, usage_client_aggregates_query_jobs.Description, )
	usage_client_aggregates_queryCmd.Long = usage_client_aggregates_queryCmd.Short
}
