package workforcemanagement_users_workplanbidranks

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_users_workplanbidranks_bulk"
)

func init() {
	workforcemanagement_users_workplanbidranksCmd.AddCommand(workforcemanagement_users_workplanbidranks_bulk.Cmdworkforcemanagement_users_workplanbidranks_bulk())
	workforcemanagement_users_workplanbidranksCmd.Short = utils.GenerateCustomDescription(workforcemanagement_users_workplanbidranksCmd.Short, workforcemanagement_users_workplanbidranks_bulk.Description, )
	workforcemanagement_users_workplanbidranksCmd.Long = workforcemanagement_users_workplanbidranksCmd.Short
}
