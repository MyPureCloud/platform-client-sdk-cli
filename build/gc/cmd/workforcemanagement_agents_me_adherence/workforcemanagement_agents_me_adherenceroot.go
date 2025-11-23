package workforcemanagement_agents_me_adherence

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_agents_me_adherence_historical"
)

func init() {
	workforcemanagement_agents_me_adherenceCmd.AddCommand(workforcemanagement_agents_me_adherence_historical.Cmdworkforcemanagement_agents_me_adherence_historical())
	workforcemanagement_agents_me_adherenceCmd.Short = utils.GenerateCustomDescription(workforcemanagement_agents_me_adherenceCmd.Short, workforcemanagement_agents_me_adherence_historical.Description, )
	workforcemanagement_agents_me_adherenceCmd.Long = workforcemanagement_agents_me_adherenceCmd.Short
}
