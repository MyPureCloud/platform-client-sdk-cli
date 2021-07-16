package workforcemanagement_calendar_data

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_calendar_data_ics"
)

func init() {
	workforcemanagement_calendar_dataCmd.AddCommand(workforcemanagement_calendar_data_ics.Cmdworkforcemanagement_calendar_data_ics())
	workforcemanagement_calendar_dataCmd.Short = utils.GenerateCustomDescription(workforcemanagement_calendar_dataCmd.Short, workforcemanagement_calendar_data_ics.Description, )
	workforcemanagement_calendar_dataCmd.Long = workforcemanagement_calendar_dataCmd.Short
}
