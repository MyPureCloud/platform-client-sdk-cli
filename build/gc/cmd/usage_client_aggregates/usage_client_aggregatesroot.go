package usage_client_aggregates

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/usage_client_aggregates_query"
)

func init() {
	usage_client_aggregatesCmd.AddCommand(usage_client_aggregates_query.Cmdusage_client_aggregates_query())
	usage_client_aggregatesCmd.Short = utils.GenerateCustomDescription(usage_client_aggregatesCmd.Short, usage_client_aggregates_query.Description, )
	usage_client_aggregatesCmd.Long = usage_client_aggregatesCmd.Short
}
