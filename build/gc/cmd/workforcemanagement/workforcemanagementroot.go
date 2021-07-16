package workforcemanagement

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_calendar"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_managementunits"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_businessunits"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_agents"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_adherence"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_timeoffrequests"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_agentschedules"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_schedules"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_shifttrades"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_notifications"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_historicaldata"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_adhocmodelingjobs"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_schedulingjobs"
)

func init() {
	workforcemanagementCmd.AddCommand(workforcemanagement_calendar.Cmdworkforcemanagement_calendar())
	workforcemanagementCmd.AddCommand(workforcemanagement_managementunits.Cmdworkforcemanagement_managementunits())
	workforcemanagementCmd.AddCommand(workforcemanagement_businessunits.Cmdworkforcemanagement_businessunits())
	workforcemanagementCmd.AddCommand(workforcemanagement_agents.Cmdworkforcemanagement_agents())
	workforcemanagementCmd.AddCommand(workforcemanagement_adherence.Cmdworkforcemanagement_adherence())
	workforcemanagementCmd.AddCommand(workforcemanagement_timeoffrequests.Cmdworkforcemanagement_timeoffrequests())
	workforcemanagementCmd.AddCommand(workforcemanagement_agentschedules.Cmdworkforcemanagement_agentschedules())
	workforcemanagementCmd.AddCommand(workforcemanagement_schedules.Cmdworkforcemanagement_schedules())
	workforcemanagementCmd.AddCommand(workforcemanagement_shifttrades.Cmdworkforcemanagement_shifttrades())
	workforcemanagementCmd.AddCommand(workforcemanagement_notifications.Cmdworkforcemanagement_notifications())
	workforcemanagementCmd.AddCommand(workforcemanagement_historicaldata.Cmdworkforcemanagement_historicaldata())
	workforcemanagementCmd.AddCommand(workforcemanagement_adhocmodelingjobs.Cmdworkforcemanagement_adhocmodelingjobs())
	workforcemanagementCmd.AddCommand(workforcemanagement_schedulingjobs.Cmdworkforcemanagement_schedulingjobs())
	workforcemanagementCmd.Short = utils.GenerateCustomDescription(workforcemanagementCmd.Short, workforcemanagement_calendar.Description, workforcemanagement_managementunits.Description, workforcemanagement_businessunits.Description, workforcemanagement_agents.Description, workforcemanagement_adherence.Description, workforcemanagement_timeoffrequests.Description, workforcemanagement_agentschedules.Description, workforcemanagement_schedules.Description, workforcemanagement_shifttrades.Description, workforcemanagement_notifications.Description, workforcemanagement_historicaldata.Description, workforcemanagement_adhocmodelingjobs.Description, workforcemanagement_schedulingjobs.Description, )
	workforcemanagementCmd.Long = workforcemanagementCmd.Short
}
