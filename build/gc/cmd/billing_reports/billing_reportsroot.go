package billing_reports

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/billing_reports_billableusage"
)

func init() {
	billing_reportsCmd.AddCommand(billing_reports_billableusage.Cmdbilling_reports_billableusage())
	billing_reportsCmd.Short = utils.GenerateCustomDescription(billing_reportsCmd.Short, billing_reports_billableusage.Description, )
	billing_reportsCmd.Long = billing_reportsCmd.Short
}
