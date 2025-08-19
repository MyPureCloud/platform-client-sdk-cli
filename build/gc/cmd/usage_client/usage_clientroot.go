package usage_client

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/usage_client_aggregates"
)

func init() {
	usage_clientCmd.AddCommand(usage_client_aggregates.Cmdusage_client_aggregates())
	usage_clientCmd.Short = utils.GenerateCustomDescription(usage_clientCmd.Short, usage_client_aggregates.Description, )
	usage_clientCmd.Long = usage_clientCmd.Short
}
