package workforcemanagement_managementunits_users

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_managementunits_users_timeoffrequests"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_managementunits_users_timeoffbalance"
)

func init() {
	workforcemanagement_managementunits_usersCmd.AddCommand(workforcemanagement_managementunits_users_timeoffrequests.Cmdworkforcemanagement_managementunits_users_timeoffrequests())
	workforcemanagement_managementunits_usersCmd.AddCommand(workforcemanagement_managementunits_users_timeoffbalance.Cmdworkforcemanagement_managementunits_users_timeoffbalance())
	workforcemanagement_managementunits_usersCmd.Short = utils.GenerateCustomDescription(workforcemanagement_managementunits_usersCmd.Short, workforcemanagement_managementunits_users_timeoffrequests.Description, workforcemanagement_managementunits_users_timeoffbalance.Description, )
	workforcemanagement_managementunits_usersCmd.Long = workforcemanagement_managementunits_usersCmd.Short
}