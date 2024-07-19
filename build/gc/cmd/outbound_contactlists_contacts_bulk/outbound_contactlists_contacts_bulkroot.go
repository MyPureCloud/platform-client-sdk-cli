package outbound_contactlists_contacts_bulk

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/outbound_contactlists_contacts_bulk_remove"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/outbound_contactlists_contacts_bulk_update"
)

func init() {
	outbound_contactlists_contacts_bulkCmd.AddCommand(outbound_contactlists_contacts_bulk_remove.Cmdoutbound_contactlists_contacts_bulk_remove())
	outbound_contactlists_contacts_bulkCmd.AddCommand(outbound_contactlists_contacts_bulk_update.Cmdoutbound_contactlists_contacts_bulk_update())
	outbound_contactlists_contacts_bulkCmd.Short = utils.GenerateCustomDescription(outbound_contactlists_contacts_bulkCmd.Short, outbound_contactlists_contacts_bulk_remove.Description, outbound_contactlists_contacts_bulk_update.Description, )
	outbound_contactlists_contacts_bulkCmd.Long = outbound_contactlists_contacts_bulkCmd.Short
}
