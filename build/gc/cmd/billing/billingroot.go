package billing

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/billing_reports"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/billing_trusteebillingoverview"
)

func init() {
	billingCmd.AddCommand(billing_reports.Cmdbilling_reports())
	billingCmd.AddCommand(billing_trusteebillingoverview.Cmdbilling_trusteebillingoverview())
	billingCmd.Short = utils.GenerateCustomDescription(billingCmd.Short, billing_reports.Description, billing_trusteebillingoverview.Description, )
	billingCmd.Long = billingCmd.Short
}
