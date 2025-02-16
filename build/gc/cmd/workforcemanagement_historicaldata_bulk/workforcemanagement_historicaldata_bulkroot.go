package workforcemanagement_historicaldata_bulk

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_historicaldata_bulk_remove"
)

func init() {
	workforcemanagement_historicaldata_bulkCmd.AddCommand(workforcemanagement_historicaldata_bulk_remove.Cmdworkforcemanagement_historicaldata_bulk_remove())
	workforcemanagement_historicaldata_bulkCmd.Short = utils.GenerateCustomDescription(workforcemanagement_historicaldata_bulkCmd.Short, workforcemanagement_historicaldata_bulk_remove.Description, )
	workforcemanagement_historicaldata_bulkCmd.Long = workforcemanagement_historicaldata_bulkCmd.Short
}
