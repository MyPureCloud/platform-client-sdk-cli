package uploads_workforcemanagement

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/uploads_workforcemanagement_historicaldata"
)

func init() {
	uploads_workforcemanagementCmd.AddCommand(uploads_workforcemanagement_historicaldata.Cmduploads_workforcemanagement_historicaldata())
	uploads_workforcemanagementCmd.Short = utils.GenerateCustomDescription(uploads_workforcemanagementCmd.Short, uploads_workforcemanagement_historicaldata.Description, )
	uploads_workforcemanagementCmd.Long = uploads_workforcemanagementCmd.Short
}
