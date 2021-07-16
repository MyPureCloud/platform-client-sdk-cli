package alerting

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/alerting_alerts"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/alerting_interactionstats"
)

func init() {
	alertingCmd.AddCommand(alerting_alerts.Cmdalerting_alerts())
	alertingCmd.AddCommand(alerting_interactionstats.Cmdalerting_interactionstats())
	alertingCmd.Short = utils.GenerateCustomDescription(alertingCmd.Short, alerting_alerts.Description, alerting_interactionstats.Description, )
	alertingCmd.Long = alertingCmd.Short
}
