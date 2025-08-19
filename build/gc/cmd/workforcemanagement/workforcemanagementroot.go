package workforcemanagement

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_calendar"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_managementunits"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_businessunits"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_agents"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_adherence"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_teams"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_alternativeshifts"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_timeofflimits"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_timeoffrequests"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_workplanbids"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_agentschedules"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_schedules"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_shifttrades"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_notifications"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_historicaldata"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_integrations"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_schedulingjobs"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_shrinkage"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_timeoffbalance"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_users"
)

func init() {
	workforcemanagementCmd.AddCommand(workforcemanagement_calendar.Cmdworkforcemanagement_calendar())
	workforcemanagementCmd.AddCommand(workforcemanagement_managementunits.Cmdworkforcemanagement_managementunits())
	workforcemanagementCmd.AddCommand(workforcemanagement_businessunits.Cmdworkforcemanagement_businessunits())
	workforcemanagementCmd.AddCommand(workforcemanagement_agents.Cmdworkforcemanagement_agents())
	workforcemanagementCmd.AddCommand(workforcemanagement_adherence.Cmdworkforcemanagement_adherence())
	workforcemanagementCmd.AddCommand(workforcemanagement_teams.Cmdworkforcemanagement_teams())
	workforcemanagementCmd.AddCommand(workforcemanagement_alternativeshifts.Cmdworkforcemanagement_alternativeshifts())
	workforcemanagementCmd.AddCommand(workforcemanagement_timeofflimits.Cmdworkforcemanagement_timeofflimits())
	workforcemanagementCmd.AddCommand(workforcemanagement_timeoffrequests.Cmdworkforcemanagement_timeoffrequests())
	workforcemanagementCmd.AddCommand(workforcemanagement_workplanbids.Cmdworkforcemanagement_workplanbids())
	workforcemanagementCmd.AddCommand(workforcemanagement_agentschedules.Cmdworkforcemanagement_agentschedules())
	workforcemanagementCmd.AddCommand(workforcemanagement_schedules.Cmdworkforcemanagement_schedules())
	workforcemanagementCmd.AddCommand(workforcemanagement_shifttrades.Cmdworkforcemanagement_shifttrades())
	workforcemanagementCmd.AddCommand(workforcemanagement_notifications.Cmdworkforcemanagement_notifications())
	workforcemanagementCmd.AddCommand(workforcemanagement_historicaldata.Cmdworkforcemanagement_historicaldata())
	workforcemanagementCmd.AddCommand(workforcemanagement_integrations.Cmdworkforcemanagement_integrations())
	workforcemanagementCmd.AddCommand(workforcemanagement_schedulingjobs.Cmdworkforcemanagement_schedulingjobs())
	workforcemanagementCmd.AddCommand(workforcemanagement_shrinkage.Cmdworkforcemanagement_shrinkage())
	workforcemanagementCmd.AddCommand(workforcemanagement_timeoffbalance.Cmdworkforcemanagement_timeoffbalance())
	workforcemanagementCmd.AddCommand(workforcemanagement_users.Cmdworkforcemanagement_users())
	workforcemanagementCmd.Short = utils.GenerateCustomDescription(workforcemanagementCmd.Short, workforcemanagement_calendar.Description, workforcemanagement_managementunits.Description, workforcemanagement_businessunits.Description, workforcemanagement_agents.Description, workforcemanagement_adherence.Description, workforcemanagement_teams.Description, workforcemanagement_alternativeshifts.Description, workforcemanagement_timeofflimits.Description, workforcemanagement_timeoffrequests.Description, workforcemanagement_workplanbids.Description, workforcemanagement_agentschedules.Description, workforcemanagement_schedules.Description, workforcemanagement_shifttrades.Description, workforcemanagement_notifications.Description, workforcemanagement_historicaldata.Description, workforcemanagement_integrations.Description, workforcemanagement_schedulingjobs.Description, workforcemanagement_shrinkage.Description, workforcemanagement_timeoffbalance.Description, workforcemanagement_users.Description, )
	workforcemanagementCmd.Long = workforcemanagementCmd.Short
}
