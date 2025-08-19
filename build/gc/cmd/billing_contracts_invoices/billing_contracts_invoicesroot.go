package billing_contracts_invoices

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/billing_contracts_invoices_document"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/billing_contracts_invoices_lines"
)

func init() {
	billing_contracts_invoicesCmd.AddCommand(billing_contracts_invoices_document.Cmdbilling_contracts_invoices_document())
	billing_contracts_invoicesCmd.AddCommand(billing_contracts_invoices_lines.Cmdbilling_contracts_invoices_lines())
	billing_contracts_invoicesCmd.Short = utils.GenerateCustomDescription(billing_contracts_invoicesCmd.Short, billing_contracts_invoices_document.Description, billing_contracts_invoices_lines.Description, )
	billing_contracts_invoicesCmd.Long = billing_contracts_invoicesCmd.Short
}
