package workforcemanagement_managementunits

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_managementunits_activitycodes"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_managementunits_move"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_managementunits_divisionviews"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_managementunits_agents"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_managementunits_users"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_managementunits_shifttrades"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_managementunits_weeks"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_managementunits_workplans"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_managementunits_workplanrotations"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_managementunits_historicaladherencequery"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_managementunits_adherence"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_managementunits_agentschedules"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_managementunits_schedules"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_managementunits_shrinkage"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_managementunits_timeoffrequests"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_managementunits_timeofflimits"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_managementunits_timeoffplans"
)

func init() {
	workforcemanagement_managementunitsCmd.AddCommand(workforcemanagement_managementunits_activitycodes.Cmdworkforcemanagement_managementunits_activitycodes())
	workforcemanagement_managementunitsCmd.AddCommand(workforcemanagement_managementunits_move.Cmdworkforcemanagement_managementunits_move())
	workforcemanagement_managementunitsCmd.AddCommand(workforcemanagement_managementunits_divisionviews.Cmdworkforcemanagement_managementunits_divisionviews())
	workforcemanagement_managementunitsCmd.AddCommand(workforcemanagement_managementunits_agents.Cmdworkforcemanagement_managementunits_agents())
	workforcemanagement_managementunitsCmd.AddCommand(workforcemanagement_managementunits_users.Cmdworkforcemanagement_managementunits_users())
	workforcemanagement_managementunitsCmd.AddCommand(workforcemanagement_managementunits_shifttrades.Cmdworkforcemanagement_managementunits_shifttrades())
	workforcemanagement_managementunitsCmd.AddCommand(workforcemanagement_managementunits_weeks.Cmdworkforcemanagement_managementunits_weeks())
	workforcemanagement_managementunitsCmd.AddCommand(workforcemanagement_managementunits_workplans.Cmdworkforcemanagement_managementunits_workplans())
	workforcemanagement_managementunitsCmd.AddCommand(workforcemanagement_managementunits_workplanrotations.Cmdworkforcemanagement_managementunits_workplanrotations())
	workforcemanagement_managementunitsCmd.AddCommand(workforcemanagement_managementunits_historicaladherencequery.Cmdworkforcemanagement_managementunits_historicaladherencequery())
	workforcemanagement_managementunitsCmd.AddCommand(workforcemanagement_managementunits_adherence.Cmdworkforcemanagement_managementunits_adherence())
	workforcemanagement_managementunitsCmd.AddCommand(workforcemanagement_managementunits_agentschedules.Cmdworkforcemanagement_managementunits_agentschedules())
	workforcemanagement_managementunitsCmd.AddCommand(workforcemanagement_managementunits_schedules.Cmdworkforcemanagement_managementunits_schedules())
	workforcemanagement_managementunitsCmd.AddCommand(workforcemanagement_managementunits_shrinkage.Cmdworkforcemanagement_managementunits_shrinkage())
	workforcemanagement_managementunitsCmd.AddCommand(workforcemanagement_managementunits_timeoffrequests.Cmdworkforcemanagement_managementunits_timeoffrequests())
	workforcemanagement_managementunitsCmd.AddCommand(workforcemanagement_managementunits_timeofflimits.Cmdworkforcemanagement_managementunits_timeofflimits())
	workforcemanagement_managementunitsCmd.AddCommand(workforcemanagement_managementunits_timeoffplans.Cmdworkforcemanagement_managementunits_timeoffplans())
	workforcemanagement_managementunitsCmd.Short = utils.GenerateCustomDescription(workforcemanagement_managementunitsCmd.Short, workforcemanagement_managementunits_activitycodes.Description, workforcemanagement_managementunits_move.Description, workforcemanagement_managementunits_divisionviews.Description, workforcemanagement_managementunits_agents.Description, workforcemanagement_managementunits_users.Description, workforcemanagement_managementunits_shifttrades.Description, workforcemanagement_managementunits_weeks.Description, workforcemanagement_managementunits_workplans.Description, workforcemanagement_managementunits_workplanrotations.Description, workforcemanagement_managementunits_historicaladherencequery.Description, workforcemanagement_managementunits_adherence.Description, workforcemanagement_managementunits_agentschedules.Description, workforcemanagement_managementunits_schedules.Description, workforcemanagement_managementunits_shrinkage.Description, workforcemanagement_managementunits_timeoffrequests.Description, workforcemanagement_managementunits_timeofflimits.Description, workforcemanagement_managementunits_timeoffplans.Description, )
	workforcemanagement_managementunitsCmd.Long = workforcemanagement_managementunitsCmd.Short
}
