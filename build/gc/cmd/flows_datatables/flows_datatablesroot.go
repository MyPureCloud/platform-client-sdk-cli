package flows_datatables

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/flows_datatables_export"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/flows_datatables_import"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/flows_datatables_rows"
)

func init() {
	flows_datatablesCmd.AddCommand(flows_datatables_export.Cmdflows_datatables_export())
	flows_datatablesCmd.AddCommand(flows_datatables_import.Cmdflows_datatables_import())
	flows_datatablesCmd.AddCommand(flows_datatables_rows.Cmdflows_datatables_rows())
	flows_datatablesCmd.Short = utils.GenerateCustomDescription(flows_datatablesCmd.Short, flows_datatables_export.Description, flows_datatables_import.Description, flows_datatables_rows.Description, )
	flows_datatablesCmd.Long = flows_datatablesCmd.Short
}
