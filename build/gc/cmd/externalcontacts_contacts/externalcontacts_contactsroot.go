package externalcontacts_contacts

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/externalcontacts_contacts_notes"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/externalcontacts_contacts_schemas"
)

func init() {
	externalcontacts_contactsCmd.AddCommand(externalcontacts_contacts_notes.Cmdexternalcontacts_contacts_notes())
	externalcontacts_contactsCmd.AddCommand(externalcontacts_contacts_schemas.Cmdexternalcontacts_contacts_schemas())
	externalcontacts_contactsCmd.Short = utils.GenerateCustomDescription(externalcontacts_contactsCmd.Short, externalcontacts_contacts_notes.Description, externalcontacts_contacts_schemas.Description, )
	externalcontacts_contactsCmd.Long = externalcontacts_contactsCmd.Short
}
