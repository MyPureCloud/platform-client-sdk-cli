package workforcemanagement_workplanbids

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_workplanbids_workplans"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_workplanbids_preferences"
)

func init() {
	workforcemanagement_workplanbidsCmd.AddCommand(workforcemanagement_workplanbids_workplans.Cmdworkforcemanagement_workplanbids_workplans())
	workforcemanagement_workplanbidsCmd.AddCommand(workforcemanagement_workplanbids_preferences.Cmdworkforcemanagement_workplanbids_preferences())
	workforcemanagement_workplanbidsCmd.Short = utils.GenerateCustomDescription(workforcemanagement_workplanbidsCmd.Short, workforcemanagement_workplanbids_workplans.Description, workforcemanagement_workplanbids_preferences.Description, )
	workforcemanagement_workplanbidsCmd.Long = workforcemanagement_workplanbidsCmd.Short
}
