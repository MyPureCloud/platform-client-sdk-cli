package outbound_campaigns_performance

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/outbound_campaigns_performance_query"
)

func init() {
	outbound_campaigns_performanceCmd.AddCommand(outbound_campaigns_performance_query.Cmdoutbound_campaigns_performance_query())
	outbound_campaigns_performanceCmd.Short = utils.GenerateCustomDescription(outbound_campaigns_performanceCmd.Short, outbound_campaigns_performance_query.Description, )
	outbound_campaigns_performanceCmd.Long = outbound_campaigns_performanceCmd.Short
}
