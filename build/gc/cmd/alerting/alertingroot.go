package alerting

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/alerting_interactionstats"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/alerting_alerts"
)

func init() {
	alertingCmd.AddCommand(alerting_interactionstats.Cmdalerting_interactionstats())
	alertingCmd.AddCommand(alerting_alerts.Cmdalerting_alerts())
	alertingCmd.Short = utils.GenerateCustomDescription(alertingCmd.Short, alerting_interactionstats.Description, alerting_alerts.Description, )
	alertingCmd.Long = alertingCmd.Short
}
