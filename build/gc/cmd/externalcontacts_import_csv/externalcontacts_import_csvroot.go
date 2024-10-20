package externalcontacts_import_csv

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/externalcontacts_import_csv_jobs"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/externalcontacts_import_csv_settings"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/externalcontacts_import_csv_uploads"
)

func init() {
	externalcontacts_import_csvCmd.AddCommand(externalcontacts_import_csv_jobs.Cmdexternalcontacts_import_csv_jobs())
	externalcontacts_import_csvCmd.AddCommand(externalcontacts_import_csv_settings.Cmdexternalcontacts_import_csv_settings())
	externalcontacts_import_csvCmd.AddCommand(externalcontacts_import_csv_uploads.Cmdexternalcontacts_import_csv_uploads())
	externalcontacts_import_csvCmd.Short = utils.GenerateCustomDescription(externalcontacts_import_csvCmd.Short, externalcontacts_import_csv_jobs.Description, externalcontacts_import_csv_settings.Description, externalcontacts_import_csv_uploads.Description, )
	externalcontacts_import_csvCmd.Long = externalcontacts_import_csvCmd.Short
}
