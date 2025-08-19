package billing

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/billing_reports"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/billing_trusteebillingoverview"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/billing_contracts"
)

func init() {
	billingCmd.AddCommand(billing_reports.Cmdbilling_reports())
	billingCmd.AddCommand(billing_trusteebillingoverview.Cmdbilling_trusteebillingoverview())
	billingCmd.AddCommand(billing_contracts.Cmdbilling_contracts())
	billingCmd.Short = utils.GenerateCustomDescription(billingCmd.Short, billing_reports.Description, billing_trusteebillingoverview.Description, billing_contracts.Description, )
	billingCmd.Long = billingCmd.Short
}
