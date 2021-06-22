package billing

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/billing_trusteebillingoverview"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/billing_reports"
)

func init() {
	billingCmd.AddCommand(billing_trusteebillingoverview.Cmdbilling_trusteebillingoverview())
	billingCmd.AddCommand(billing_reports.Cmdbilling_reports())
	billingCmd.Short = utils.GenerateCustomDescription(billingCmd.Short, billing_trusteebillingoverview.Description, billing_reports.Description, )
	billingCmd.Long = billingCmd.Short
}
