package externalcontacts_scan_contacts

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/externalcontacts_scan_contacts_divisionviews"
)

func init() {
	externalcontacts_scan_contactsCmd.AddCommand(externalcontacts_scan_contacts_divisionviews.Cmdexternalcontacts_scan_contacts_divisionviews())
	externalcontacts_scan_contactsCmd.Short = utils.GenerateCustomDescription(externalcontacts_scan_contactsCmd.Short, externalcontacts_scan_contacts_divisionviews.Description, )
	externalcontacts_scan_contactsCmd.Long = externalcontacts_scan_contactsCmd.Short
}
