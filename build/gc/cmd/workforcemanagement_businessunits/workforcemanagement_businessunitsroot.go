package workforcemanagement_businessunits

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_businessunits_activitycodes"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_businessunits_intraday"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_businessunits_weeks"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_businessunits_managementunits"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_businessunits_divisionviews"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_businessunits_planninggroups"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_businessunits_servicegoaltemplates"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_businessunits_scheduling"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_businessunits_agentschedules"
)

func init() {
	workforcemanagement_businessunitsCmd.AddCommand(workforcemanagement_businessunits_activitycodes.Cmdworkforcemanagement_businessunits_activitycodes())
	workforcemanagement_businessunitsCmd.AddCommand(workforcemanagement_businessunits_intraday.Cmdworkforcemanagement_businessunits_intraday())
	workforcemanagement_businessunitsCmd.AddCommand(workforcemanagement_businessunits_weeks.Cmdworkforcemanagement_businessunits_weeks())
	workforcemanagement_businessunitsCmd.AddCommand(workforcemanagement_businessunits_managementunits.Cmdworkforcemanagement_businessunits_managementunits())
	workforcemanagement_businessunitsCmd.AddCommand(workforcemanagement_businessunits_divisionviews.Cmdworkforcemanagement_businessunits_divisionviews())
	workforcemanagement_businessunitsCmd.AddCommand(workforcemanagement_businessunits_planninggroups.Cmdworkforcemanagement_businessunits_planninggroups())
	workforcemanagement_businessunitsCmd.AddCommand(workforcemanagement_businessunits_servicegoaltemplates.Cmdworkforcemanagement_businessunits_servicegoaltemplates())
	workforcemanagement_businessunitsCmd.AddCommand(workforcemanagement_businessunits_scheduling.Cmdworkforcemanagement_businessunits_scheduling())
	workforcemanagement_businessunitsCmd.AddCommand(workforcemanagement_businessunits_agentschedules.Cmdworkforcemanagement_businessunits_agentschedules())
	workforcemanagement_businessunitsCmd.Short = utils.GenerateCustomDescription(workforcemanagement_businessunitsCmd.Short, workforcemanagement_businessunits_activitycodes.Description, workforcemanagement_businessunits_intraday.Description, workforcemanagement_businessunits_weeks.Description, workforcemanagement_businessunits_managementunits.Description, workforcemanagement_businessunits_divisionviews.Description, workforcemanagement_businessunits_planninggroups.Description, workforcemanagement_businessunits_servicegoaltemplates.Description, workforcemanagement_businessunits_scheduling.Description, workforcemanagement_businessunits_agentschedules.Description, )
	workforcemanagement_businessunitsCmd.Long = workforcemanagement_businessunitsCmd.Short
}