package billing_contracts

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/billing_contracts_billingperiods"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/billing_contracts_invoices"
)

func init() {
	billing_contractsCmd.AddCommand(billing_contracts_billingperiods.Cmdbilling_contracts_billingperiods())
	billing_contractsCmd.AddCommand(billing_contracts_invoices.Cmdbilling_contracts_invoices())
	billing_contractsCmd.Short = utils.GenerateCustomDescription(billing_contractsCmd.Short, billing_contracts_billingperiods.Description, billing_contracts_invoices.Description, )
	billing_contractsCmd.Long = billing_contractsCmd.Short
}
