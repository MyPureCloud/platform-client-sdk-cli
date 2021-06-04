package alerting_interactionstats_alerts

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/alerting_interactionstats_alerts_unread"
)

func init() {
	alerting_interactionstats_alertsCmd.AddCommand(alerting_interactionstats_alerts_unread.Cmdalerting_interactionstats_alerts_unread())
	alerting_interactionstats_alertsCmd.Short = utils.GenerateCustomDescription(alerting_interactionstats_alertsCmd.Short, alerting_interactionstats_alerts_unread.Description, )
	alerting_interactionstats_alertsCmd.Long = alerting_interactionstats_alertsCmd.Short
}
