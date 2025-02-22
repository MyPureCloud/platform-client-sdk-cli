package workforcemanagement_businessunits_mainforecast_continuousforecast

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_businessunits_mainforecast_continuousforecast_session"
)

func init() {
	workforcemanagement_businessunits_mainforecast_continuousforecastCmd.AddCommand(workforcemanagement_businessunits_mainforecast_continuousforecast_session.Cmdworkforcemanagement_businessunits_mainforecast_continuousforecast_session())
	workforcemanagement_businessunits_mainforecast_continuousforecastCmd.Short = utils.GenerateCustomDescription(workforcemanagement_businessunits_mainforecast_continuousforecastCmd.Short, workforcemanagement_businessunits_mainforecast_continuousforecast_session.Description, )
	workforcemanagement_businessunits_mainforecast_continuousforecastCmd.Long = workforcemanagement_businessunits_mainforecast_continuousforecastCmd.Short
}
