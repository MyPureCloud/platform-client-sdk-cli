package alerting

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/alerting_alerts"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/alerting_interactionstats"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/alerting_rules"
)

func init() {
	alertingCmd.AddCommand(alerting_alerts.Cmdalerting_alerts())
	alertingCmd.AddCommand(alerting_interactionstats.Cmdalerting_interactionstats())
	alertingCmd.AddCommand(alerting_rules.Cmdalerting_rules())
	alertingCmd.Short = utils.GenerateCustomDescription(alertingCmd.Short, alerting_alerts.Description, alerting_interactionstats.Description, alerting_rules.Description, )
	alertingCmd.Long = alertingCmd.Short
}
