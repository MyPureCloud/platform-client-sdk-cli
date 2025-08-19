package workforcemanagement_teams

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_teams_adherence"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_teams_shrinkage"
)

func init() {
	workforcemanagement_teamsCmd.AddCommand(workforcemanagement_teams_adherence.Cmdworkforcemanagement_teams_adherence())
	workforcemanagement_teamsCmd.AddCommand(workforcemanagement_teams_shrinkage.Cmdworkforcemanagement_teams_shrinkage())
	workforcemanagement_teamsCmd.Short = utils.GenerateCustomDescription(workforcemanagement_teamsCmd.Short, workforcemanagement_teams_adherence.Description, workforcemanagement_teams_shrinkage.Description, )
	workforcemanagement_teamsCmd.Long = workforcemanagement_teamsCmd.Short
}
