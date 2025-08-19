package analytics_bots

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_bots_aggregates"
)

func init() {
	analytics_botsCmd.AddCommand(analytics_bots_aggregates.Cmdanalytics_bots_aggregates())
	analytics_botsCmd.Short = utils.GenerateCustomDescription(analytics_botsCmd.Short, analytics_bots_aggregates.Description, )
	analytics_botsCmd.Long = analytics_botsCmd.Short
}
