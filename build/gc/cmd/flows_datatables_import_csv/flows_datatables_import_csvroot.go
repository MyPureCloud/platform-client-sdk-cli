package flows_datatables_import_csv

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/flows_datatables_import_csv_jobs"
)

func init() {
	flows_datatables_import_csvCmd.AddCommand(flows_datatables_import_csv_jobs.Cmdflows_datatables_import_csv_jobs())
	flows_datatables_import_csvCmd.Short = utils.GenerateCustomDescription(flows_datatables_import_csvCmd.Short, flows_datatables_import_csv_jobs.Description, )
	flows_datatables_import_csvCmd.Long = flows_datatables_import_csvCmd.Short
}
