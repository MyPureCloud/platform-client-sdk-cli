package analytics_copilots

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_copilots_aggregates"
)

func init() {
	analytics_copilotsCmd.AddCommand(analytics_copilots_aggregates.Cmdanalytics_copilots_aggregates())
	analytics_copilotsCmd.Short = utils.GenerateCustomDescription(analytics_copilotsCmd.Short, analytics_copilots_aggregates.Description, )
	analytics_copilotsCmd.Long = analytics_copilotsCmd.Short
}
