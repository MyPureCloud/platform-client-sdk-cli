package workforcemanagement_managementunits_users_timeoffbalance

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("workforcemanagement_managementunits_users_timeoffbalance", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/managementunits/{managementUnitId}/users/{userId}/timeoffbalance")
	workforcemanagement_managementunits_users_timeoffbalanceCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement_managementunits_users_timeoffbalance"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagement_managementunits_users_timeoffbalanceCmd)
}

func Cmdworkforcemanagement_managementunits_users_timeoffbalance() *cobra.Command {
	return workforcemanagement_managementunits_users_timeoffbalanceCmd
}
