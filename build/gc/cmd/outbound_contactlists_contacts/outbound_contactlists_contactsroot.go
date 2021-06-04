package outbound_contactlists_contacts

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/outbound_contactlists_contacts_bulk"
)

func init() {
	outbound_contactlists_contactsCmd.AddCommand(outbound_contactlists_contacts_bulk.Cmdoutbound_contactlists_contacts_bulk())
	outbound_contactlists_contactsCmd.Short = utils.GenerateCustomDescription(outbound_contactlists_contactsCmd.Short, outbound_contactlists_contacts_bulk.Description, )
	outbound_contactlists_contactsCmd.Long = outbound_contactlists_contactsCmd.Short
}
