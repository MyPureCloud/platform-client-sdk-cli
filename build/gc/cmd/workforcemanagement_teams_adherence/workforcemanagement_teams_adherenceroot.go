package workforcemanagement_teams_adherence

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_teams_adherence_historical"
)

func init() {
	workforcemanagement_teams_adherenceCmd.AddCommand(workforcemanagement_teams_adherence_historical.Cmdworkforcemanagement_teams_adherence_historical())
	workforcemanagement_teams_adherenceCmd.Short = utils.GenerateCustomDescription(workforcemanagement_teams_adherenceCmd.Short, workforcemanagement_teams_adherence_historical.Description, )
	workforcemanagement_teams_adherenceCmd.Long = workforcemanagement_teams_adherenceCmd.Short
}
