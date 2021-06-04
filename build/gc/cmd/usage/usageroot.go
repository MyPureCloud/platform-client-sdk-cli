package usage

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/usage_query"
)

func init() {
	usageCmd.AddCommand(usage_query.Cmdusage_query())
	usageCmd.Short = utils.GenerateCustomDescription(usageCmd.Short, usage_query.Description, )
	usageCmd.Long = usageCmd.Short
}
