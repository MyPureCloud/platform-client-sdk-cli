package workforcemanagement_agents_adherence

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_agents_adherence_explanations"
)

func init() {
	workforcemanagement_agents_adherenceCmd.AddCommand(workforcemanagement_agents_adherence_explanations.Cmdworkforcemanagement_agents_adherence_explanations())
	workforcemanagement_agents_adherenceCmd.Short = utils.GenerateCustomDescription(workforcemanagement_agents_adherenceCmd.Short, workforcemanagement_agents_adherence_explanations.Description, )
	workforcemanagement_agents_adherenceCmd.Long = workforcemanagement_agents_adherenceCmd.Short
}
