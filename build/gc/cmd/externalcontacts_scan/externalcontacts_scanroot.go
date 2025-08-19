package externalcontacts_scan

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/externalcontacts_scan_contacts"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/externalcontacts_scan_organizations"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/externalcontacts_scan_notes"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/externalcontacts_scan_relationships"
)

func init() {
	externalcontacts_scanCmd.AddCommand(externalcontacts_scan_contacts.Cmdexternalcontacts_scan_contacts())
	externalcontacts_scanCmd.AddCommand(externalcontacts_scan_organizations.Cmdexternalcontacts_scan_organizations())
	externalcontacts_scanCmd.AddCommand(externalcontacts_scan_notes.Cmdexternalcontacts_scan_notes())
	externalcontacts_scanCmd.AddCommand(externalcontacts_scan_relationships.Cmdexternalcontacts_scan_relationships())
	externalcontacts_scanCmd.Short = utils.GenerateCustomDescription(externalcontacts_scanCmd.Short, externalcontacts_scan_contacts.Description, externalcontacts_scan_organizations.Description, externalcontacts_scan_notes.Description, externalcontacts_scan_relationships.Description, )
	externalcontacts_scanCmd.Long = externalcontacts_scanCmd.Short
}
