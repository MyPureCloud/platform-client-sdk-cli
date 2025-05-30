package externalcontacts

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/externalcontacts_bulk"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/externalcontacts_contacts"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/externalcontacts_merge"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/externalcontacts_scan"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/externalcontacts_conversations"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/externalcontacts_organizations"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/externalcontacts_externalsources"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/externalcontacts_identifierlookup"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/externalcontacts_relationships"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/externalcontacts_reversewhitepageslookup"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/externalcontacts_import"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/externalcontacts_settings"
)

func init() {
	externalcontactsCmd.AddCommand(externalcontacts_bulk.Cmdexternalcontacts_bulk())
	externalcontactsCmd.AddCommand(externalcontacts_contacts.Cmdexternalcontacts_contacts())
	externalcontactsCmd.AddCommand(externalcontacts_merge.Cmdexternalcontacts_merge())
	externalcontactsCmd.AddCommand(externalcontacts_scan.Cmdexternalcontacts_scan())
	externalcontactsCmd.AddCommand(externalcontacts_conversations.Cmdexternalcontacts_conversations())
	externalcontactsCmd.AddCommand(externalcontacts_organizations.Cmdexternalcontacts_organizations())
	externalcontactsCmd.AddCommand(externalcontacts_externalsources.Cmdexternalcontacts_externalsources())
	externalcontactsCmd.AddCommand(externalcontacts_identifierlookup.Cmdexternalcontacts_identifierlookup())
	externalcontactsCmd.AddCommand(externalcontacts_relationships.Cmdexternalcontacts_relationships())
	externalcontactsCmd.AddCommand(externalcontacts_reversewhitepageslookup.Cmdexternalcontacts_reversewhitepageslookup())
	externalcontactsCmd.AddCommand(externalcontacts_import.Cmdexternalcontacts_import())
	externalcontactsCmd.AddCommand(externalcontacts_settings.Cmdexternalcontacts_settings())
	externalcontactsCmd.Short = utils.GenerateCustomDescription(externalcontactsCmd.Short, externalcontacts_bulk.Description, externalcontacts_contacts.Description, externalcontacts_merge.Description, externalcontacts_scan.Description, externalcontacts_conversations.Description, externalcontacts_organizations.Description, externalcontacts_externalsources.Description, externalcontacts_identifierlookup.Description, externalcontacts_relationships.Description, externalcontacts_reversewhitepageslookup.Description, externalcontacts_import.Description, externalcontacts_settings.Description, )
	externalcontactsCmd.Long = externalcontactsCmd.Short
}
