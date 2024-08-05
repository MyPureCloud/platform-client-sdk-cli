package workforcemanagement_users

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_users_workplanbidranks"
)

func init() {
	workforcemanagement_usersCmd.AddCommand(workforcemanagement_users_workplanbidranks.Cmdworkforcemanagement_users_workplanbidranks())
	workforcemanagement_usersCmd.Short = utils.GenerateCustomDescription(workforcemanagement_usersCmd.Short, workforcemanagement_users_workplanbidranks.Description, )
	workforcemanagement_usersCmd.Long = workforcemanagement_usersCmd.Short
}
