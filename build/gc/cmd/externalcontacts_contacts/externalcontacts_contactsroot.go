package externalcontacts_contacts

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/externalcontacts_contacts_notes"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/externalcontacts_contacts_promotion"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/externalcontacts_contacts_identifiers"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/externalcontacts_contacts_schemas"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/externalcontacts_contacts_journey"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/externalcontacts_contacts_unresolved"
)

func init() {
	externalcontacts_contactsCmd.AddCommand(externalcontacts_contacts_notes.Cmdexternalcontacts_contacts_notes())
	externalcontacts_contactsCmd.AddCommand(externalcontacts_contacts_promotion.Cmdexternalcontacts_contacts_promotion())
	externalcontacts_contactsCmd.AddCommand(externalcontacts_contacts_identifiers.Cmdexternalcontacts_contacts_identifiers())
	externalcontacts_contactsCmd.AddCommand(externalcontacts_contacts_schemas.Cmdexternalcontacts_contacts_schemas())
	externalcontacts_contactsCmd.AddCommand(externalcontacts_contacts_journey.Cmdexternalcontacts_contacts_journey())
	externalcontacts_contactsCmd.AddCommand(externalcontacts_contacts_unresolved.Cmdexternalcontacts_contacts_unresolved())
	externalcontacts_contactsCmd.Short = utils.GenerateCustomDescription(externalcontacts_contactsCmd.Short, externalcontacts_contacts_notes.Description, externalcontacts_contacts_promotion.Description, externalcontacts_contacts_identifiers.Description, externalcontacts_contacts_schemas.Description, externalcontacts_contacts_journey.Description, externalcontacts_contacts_unresolved.Description, )
	externalcontacts_contactsCmd.Long = externalcontacts_contactsCmd.Short
}
