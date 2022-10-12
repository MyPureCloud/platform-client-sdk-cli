package externalcontacts_merge

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/externalcontacts_merge_contacts"
)

func init() {
	externalcontacts_mergeCmd.AddCommand(externalcontacts_merge_contacts.Cmdexternalcontacts_merge_contacts())
	externalcontacts_mergeCmd.Short = utils.GenerateCustomDescription(externalcontacts_mergeCmd.Short, externalcontacts_merge_contacts.Description, )
	externalcontacts_mergeCmd.Long = externalcontacts_mergeCmd.Short
}
