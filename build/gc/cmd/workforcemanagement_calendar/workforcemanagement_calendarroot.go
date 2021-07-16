package workforcemanagement_calendar

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_calendar_data"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_calendar_url"
)

func init() {
	workforcemanagement_calendarCmd.AddCommand(workforcemanagement_calendar_data.Cmdworkforcemanagement_calendar_data())
	workforcemanagement_calendarCmd.AddCommand(workforcemanagement_calendar_url.Cmdworkforcemanagement_calendar_url())
	workforcemanagement_calendarCmd.Short = utils.GenerateCustomDescription(workforcemanagement_calendarCmd.Short, workforcemanagement_calendar_data.Description, workforcemanagement_calendar_url.Description, )
	workforcemanagement_calendarCmd.Long = workforcemanagement_calendarCmd.Short
}
