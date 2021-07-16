package workforcemanagement_calendar_url

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_calendar_url_ics"
)

func init() {
	workforcemanagement_calendar_urlCmd.AddCommand(workforcemanagement_calendar_url_ics.Cmdworkforcemanagement_calendar_url_ics())
	workforcemanagement_calendar_urlCmd.Short = utils.GenerateCustomDescription(workforcemanagement_calendar_urlCmd.Short, workforcemanagement_calendar_url_ics.Description, )
	workforcemanagement_calendar_urlCmd.Long = workforcemanagement_calendar_urlCmd.Short
}
