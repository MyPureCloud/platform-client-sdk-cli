package analytics_agentcopilots

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_agentcopilots_aggregates"
)

func init() {
	analytics_agentcopilotsCmd.AddCommand(analytics_agentcopilots_aggregates.Cmdanalytics_agentcopilots_aggregates())
	analytics_agentcopilotsCmd.Short = utils.GenerateCustomDescription(analytics_agentcopilotsCmd.Short, analytics_agentcopilots_aggregates.Description, )
	analytics_agentcopilotsCmd.Long = analytics_agentcopilotsCmd.Short
}
