package workforcemanagement_historicaldata

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_historicaldata_deletejob"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_historicaldata_importstatus"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_historicaldata_validate"
)

func init() {
	workforcemanagement_historicaldataCmd.AddCommand(workforcemanagement_historicaldata_deletejob.Cmdworkforcemanagement_historicaldata_deletejob())
	workforcemanagement_historicaldataCmd.AddCommand(workforcemanagement_historicaldata_importstatus.Cmdworkforcemanagement_historicaldata_importstatus())
	workforcemanagement_historicaldataCmd.AddCommand(workforcemanagement_historicaldata_validate.Cmdworkforcemanagement_historicaldata_validate())
	workforcemanagement_historicaldataCmd.Short = utils.GenerateCustomDescription(workforcemanagement_historicaldataCmd.Short, workforcemanagement_historicaldata_deletejob.Description, workforcemanagement_historicaldata_importstatus.Description, workforcemanagement_historicaldata_validate.Description, )
	workforcemanagement_historicaldataCmd.Long = workforcemanagement_historicaldataCmd.Short
}