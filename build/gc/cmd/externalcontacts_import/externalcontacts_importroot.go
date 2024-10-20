package externalcontacts_import

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/externalcontacts_import_csv"
)

func init() {
	externalcontacts_importCmd.AddCommand(externalcontacts_import_csv.Cmdexternalcontacts_import_csv())
	externalcontacts_importCmd.Short = utils.GenerateCustomDescription(externalcontacts_importCmd.Short, externalcontacts_import_csv.Description, )
	externalcontacts_importCmd.Long = externalcontacts_importCmd.Short
}
