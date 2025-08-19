package alerting_rules

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/alerting_rules_bulk"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/alerting_rules_query"
)

func init() {
	alerting_rulesCmd.AddCommand(alerting_rules_bulk.Cmdalerting_rules_bulk())
	alerting_rulesCmd.AddCommand(alerting_rules_query.Cmdalerting_rules_query())
	alerting_rulesCmd.Short = utils.GenerateCustomDescription(alerting_rulesCmd.Short, alerting_rules_bulk.Description, alerting_rules_query.Description, )
	alerting_rulesCmd.Long = alerting_rulesCmd.Short
}
