package workforcemanagement_businessunits_intraday

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_businessunits_intraday_planninggroups"
)

func init() {
	workforcemanagement_businessunits_intradayCmd.AddCommand(workforcemanagement_businessunits_intraday_planninggroups.Cmdworkforcemanagement_businessunits_intraday_planninggroups())
	workforcemanagement_businessunits_intradayCmd.Short = utils.GenerateCustomDescription(workforcemanagement_businessunits_intradayCmd.Short, workforcemanagement_businessunits_intraday_planninggroups.Description, )
	workforcemanagement_businessunits_intradayCmd.Long = workforcemanagement_businessunits_intradayCmd.Short
}
