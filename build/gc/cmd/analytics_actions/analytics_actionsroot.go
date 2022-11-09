package analytics_actions

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_actions_aggregates"
)

func init() {
	analytics_actionsCmd.AddCommand(analytics_actions_aggregates.Cmdanalytics_actions_aggregates())
	analytics_actionsCmd.Short = utils.GenerateCustomDescription(analytics_actionsCmd.Short, analytics_actions_aggregates.Description, )
	analytics_actionsCmd.Long = analytics_actionsCmd.Short
}
