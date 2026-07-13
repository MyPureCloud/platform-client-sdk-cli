package externalcontacts_relationships

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/externalcontacts_relationships_exports"
)

func init() {
	externalcontacts_relationshipsCmd.AddCommand(externalcontacts_relationships_exports.Cmdexternalcontacts_relationships_exports())
	externalcontacts_relationshipsCmd.Short = utils.GenerateCustomDescription(externalcontacts_relationshipsCmd.Short, externalcontacts_relationships_exports.Description, )
	externalcontacts_relationshipsCmd.Long = externalcontacts_relationshipsCmd.Short
}
