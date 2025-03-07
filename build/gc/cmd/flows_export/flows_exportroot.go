package flows_export

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/flows_export_jobs"
)

func init() {
	flows_exportCmd.AddCommand(flows_export_jobs.Cmdflows_export_jobs())
	flows_exportCmd.Short = utils.GenerateCustomDescription(flows_exportCmd.Short, flows_export_jobs.Description, )
	flows_exportCmd.Long = flows_exportCmd.Short
}
