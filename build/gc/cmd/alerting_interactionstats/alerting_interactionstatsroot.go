package alerting_interactionstats

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/alerting_interactionstats_rules"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/alerting_interactionstats_alerts"
)

func init() {
	alerting_interactionstatsCmd.AddCommand(alerting_interactionstats_rules.Cmdalerting_interactionstats_rules())
	alerting_interactionstatsCmd.AddCommand(alerting_interactionstats_alerts.Cmdalerting_interactionstats_alerts())
	alerting_interactionstatsCmd.Short = utils.GenerateCustomDescription(alerting_interactionstatsCmd.Short, alerting_interactionstats_rules.Description, alerting_interactionstats_alerts.Description, )
	alerting_interactionstatsCmd.Long = alerting_interactionstatsCmd.Short
}
