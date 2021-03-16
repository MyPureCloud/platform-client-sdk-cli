
package contactlists

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/contacts_contactlists"
)

func init() {
	contactlistsCmd.AddCommand(contacts_contactlists.Cmdcontacts_contactlists())
}
