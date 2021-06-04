package alerting_alerts

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/alerting_alerts_active"
)

func init() {
	alerting_alertsCmd.AddCommand(alerting_alerts_active.Cmdalerting_alerts_active())
	alerting_alertsCmd.Short = utils.GenerateCustomDescription(alerting_alertsCmd.Short, alerting_alerts_active.Description, )
	alerting_alertsCmd.Long = alerting_alertsCmd.Short
}
