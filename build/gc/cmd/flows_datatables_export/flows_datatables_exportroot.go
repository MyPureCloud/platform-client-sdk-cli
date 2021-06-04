package flows_datatables_export

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/flows_datatables_export_jobs"
)

func init() {
	flows_datatables_exportCmd.AddCommand(flows_datatables_export_jobs.Cmdflows_datatables_export_jobs())
	flows_datatables_exportCmd.Short = utils.GenerateCustomDescription(flows_datatables_exportCmd.Short, flows_datatables_export_jobs.Description, )
	flows_datatables_exportCmd.Long = flows_datatables_exportCmd.Short
}
