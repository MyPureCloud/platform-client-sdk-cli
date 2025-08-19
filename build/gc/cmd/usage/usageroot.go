package usage

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/usage_events"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/usage_query"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/usage_simplesearch"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/usage_client"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/usage_aggregates"
)

func init() {
	usageCmd.AddCommand(usage_events.Cmdusage_events())
	usageCmd.AddCommand(usage_query.Cmdusage_query())
	usageCmd.AddCommand(usage_simplesearch.Cmdusage_simplesearch())
	usageCmd.AddCommand(usage_client.Cmdusage_client())
	usageCmd.AddCommand(usage_aggregates.Cmdusage_aggregates())
	usageCmd.Short = utils.GenerateCustomDescription(usageCmd.Short, usage_events.Description, usage_query.Description, usage_simplesearch.Description, usage_client.Description, usage_aggregates.Description, )
	usageCmd.Long = usageCmd.Short
}
