package flows_datatables_import

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/flows_datatables_import_jobs"
)

func init() {
	flows_datatables_importCmd.AddCommand(flows_datatables_import_jobs.Cmdflows_datatables_import_jobs())
	flows_datatables_importCmd.Short = utils.GenerateCustomDescription(flows_datatables_importCmd.Short, flows_datatables_import_jobs.Description, )
	flows_datatables_importCmd.Long = flows_datatables_importCmd.Short
}
