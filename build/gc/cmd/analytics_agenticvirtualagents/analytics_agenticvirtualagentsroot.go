package analytics_agenticvirtualagents

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_agenticvirtualagents_aggregates"
)

func init() {
	analytics_agenticvirtualagentsCmd.AddCommand(analytics_agenticvirtualagents_aggregates.Cmdanalytics_agenticvirtualagents_aggregates())
	analytics_agenticvirtualagentsCmd.Short = utils.GenerateCustomDescription(analytics_agenticvirtualagentsCmd.Short, analytics_agenticvirtualagents_aggregates.Description, )
	analytics_agenticvirtualagentsCmd.Long = analytics_agenticvirtualagentsCmd.Short
}
