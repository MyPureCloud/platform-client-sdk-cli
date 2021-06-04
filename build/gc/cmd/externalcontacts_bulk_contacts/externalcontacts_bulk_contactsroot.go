package externalcontacts_bulk_contacts

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/externalcontacts_bulk_contacts_update"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/externalcontacts_bulk_contacts_add"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/externalcontacts_bulk_contacts_remove"
)

func init() {
	externalcontacts_bulk_contactsCmd.AddCommand(externalcontacts_bulk_contacts_update.Cmdexternalcontacts_bulk_contacts_update())
	externalcontacts_bulk_contactsCmd.AddCommand(externalcontacts_bulk_contacts_add.Cmdexternalcontacts_bulk_contacts_add())
	externalcontacts_bulk_contactsCmd.AddCommand(externalcontacts_bulk_contacts_remove.Cmdexternalcontacts_bulk_contacts_remove())
	externalcontacts_bulk_contactsCmd.Short = utils.GenerateCustomDescription(externalcontacts_bulk_contactsCmd.Short, externalcontacts_bulk_contacts_update.Description, externalcontacts_bulk_contacts_add.Description, externalcontacts_bulk_contacts_remove.Description, )
	externalcontacts_bulk_contactsCmd.Long = externalcontacts_bulk_contactsCmd.Short
}
