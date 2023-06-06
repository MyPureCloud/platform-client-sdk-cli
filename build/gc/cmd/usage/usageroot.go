package usage

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/usage_query"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/usage_simplesearch"
)

func init() {
	usageCmd.AddCommand(usage_query.Cmdusage_query())
	usageCmd.AddCommand(usage_simplesearch.Cmdusage_simplesearch())
	usageCmd.Short = utils.GenerateCustomDescription(usageCmd.Short, usage_query.Description, usage_simplesearch.Description, )
	usageCmd.Long = usageCmd.Short
}
