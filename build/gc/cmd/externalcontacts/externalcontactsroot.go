
package externalcontacts

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/contacts_externalcontacts"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/organizations"
)

func init() {
	externalcontactsCmd.AddCommand(contacts_externalcontacts.Cmdcontacts_externalcontacts())
	externalcontactsCmd.AddCommand(organizations.Cmdorganizations())
}
