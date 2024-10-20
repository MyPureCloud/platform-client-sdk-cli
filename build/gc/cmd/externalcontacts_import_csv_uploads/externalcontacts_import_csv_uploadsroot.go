package externalcontacts_import_csv_uploads

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/externalcontacts_import_csv_uploads_details"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/externalcontacts_import_csv_uploads_preview"
)

func init() {
	externalcontacts_import_csv_uploadsCmd.AddCommand(externalcontacts_import_csv_uploads_details.Cmdexternalcontacts_import_csv_uploads_details())
	externalcontacts_import_csv_uploadsCmd.AddCommand(externalcontacts_import_csv_uploads_preview.Cmdexternalcontacts_import_csv_uploads_preview())
	externalcontacts_import_csv_uploadsCmd.Short = utils.GenerateCustomDescription(externalcontacts_import_csv_uploadsCmd.Short, externalcontacts_import_csv_uploads_details.Description, externalcontacts_import_csv_uploads_preview.Description, )
	externalcontacts_import_csv_uploadsCmd.Long = externalcontacts_import_csv_uploadsCmd.Short
}
