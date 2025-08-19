package alerting_rules_bulk

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/alerting_rules_bulk_remove"
)

func init() {
	alerting_rules_bulkCmd.AddCommand(alerting_rules_bulk_remove.Cmdalerting_rules_bulk_remove())
	alerting_rules_bulkCmd.Short = utils.GenerateCustomDescription(alerting_rules_bulkCmd.Short, alerting_rules_bulk_remove.Description, )
	alerting_rules_bulkCmd.Long = alerting_rules_bulkCmd.Short
}
