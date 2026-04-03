package externalcontacts_identifierlookup

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/externalcontacts_identifierlookup_organizations"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/externalcontacts_identifierlookup_contacts"
)

func init() {
	externalcontacts_identifierlookupCmd.AddCommand(externalcontacts_identifierlookup_organizations.Cmdexternalcontacts_identifierlookup_organizations())
	externalcontacts_identifierlookupCmd.AddCommand(externalcontacts_identifierlookup_contacts.Cmdexternalcontacts_identifierlookup_contacts())
	externalcontacts_identifierlookupCmd.Short = utils.GenerateCustomDescription(externalcontacts_identifierlookupCmd.Short, externalcontacts_identifierlookup_organizations.Description, externalcontacts_identifierlookup_contacts.Description, )
	externalcontacts_identifierlookupCmd.Long = externalcontacts_identifierlookupCmd.Short
}
