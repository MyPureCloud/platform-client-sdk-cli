package alerting_alerts

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/alerting_alerts_active"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/alerting_alerts_bulk"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/alerting_alerts_query"
)

func init() {
	alerting_alertsCmd.AddCommand(alerting_alerts_active.Cmdalerting_alerts_active())
	alerting_alertsCmd.AddCommand(alerting_alerts_bulk.Cmdalerting_alerts_bulk())
	alerting_alertsCmd.AddCommand(alerting_alerts_query.Cmdalerting_alerts_query())
	alerting_alertsCmd.Short = utils.GenerateCustomDescription(alerting_alertsCmd.Short, alerting_alerts_active.Description, alerting_alerts_bulk.Description, alerting_alerts_query.Description, )
	alerting_alertsCmd.Long = alerting_alertsCmd.Short
}
