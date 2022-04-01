package uploads_workforcemanagement_historicaldata

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/uploads_workforcemanagement_historicaldata_csv"
)

func init() {
	uploads_workforcemanagement_historicaldataCmd.AddCommand(uploads_workforcemanagement_historicaldata_csv.Cmduploads_workforcemanagement_historicaldata_csv())
	uploads_workforcemanagement_historicaldataCmd.Short = utils.GenerateCustomDescription(uploads_workforcemanagement_historicaldataCmd.Short, uploads_workforcemanagement_historicaldata_csv.Description, )
	uploads_workforcemanagement_historicaldataCmd.Long = uploads_workforcemanagement_historicaldataCmd.Short
}
