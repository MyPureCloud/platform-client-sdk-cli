package usage_simplesearch

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/usage_simplesearch_results"
)

func init() {
	usage_simplesearchCmd.AddCommand(usage_simplesearch_results.Cmdusage_simplesearch_results())
	usage_simplesearchCmd.Short = utils.GenerateCustomDescription(usage_simplesearchCmd.Short, usage_simplesearch_results.Description, )
	usage_simplesearchCmd.Long = usage_simplesearchCmd.Short
}
