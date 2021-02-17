
package externalcontacts

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/contacts"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/organizations"
)

func init() {
	externalcontactsCmd.AddCommand(contacts.Cmdcontacts())
	externalcontactsCmd.AddCommand(organizations.Cmdorganizations())
}
