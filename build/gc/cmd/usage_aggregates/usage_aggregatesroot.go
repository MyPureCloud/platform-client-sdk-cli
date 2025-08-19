package usage_aggregates

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/usage_aggregates_query"
)

func init() {
	usage_aggregatesCmd.AddCommand(usage_aggregates_query.Cmdusage_aggregates_query())
	usage_aggregatesCmd.Short = utils.GenerateCustomDescription(usage_aggregatesCmd.Short, usage_aggregates_query.Description, )
	usage_aggregatesCmd.Long = usage_aggregatesCmd.Short
}
