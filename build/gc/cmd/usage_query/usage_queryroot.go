package usage_query

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/usage_query_results"
)

func init() {
	usage_queryCmd.AddCommand(usage_query_results.Cmdusage_query_results())
	usage_queryCmd.Short = utils.GenerateCustomDescription(usage_queryCmd.Short, usage_query_results.Description, )
	usage_queryCmd.Long = usage_queryCmd.Short
}
