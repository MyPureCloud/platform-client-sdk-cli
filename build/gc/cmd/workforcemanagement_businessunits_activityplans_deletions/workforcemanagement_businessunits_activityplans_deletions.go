package workforcemanagement_businessunits_activityplans_deletions

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("workforcemanagement_businessunits_activityplans_deletions", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/businessunits/{businessUnitId}/activityplans/{activityPlanId}/deletions")
	workforcemanagement_businessunits_activityplans_deletionsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement_businessunits_activityplans_deletions"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagement_businessunits_activityplans_deletionsCmd)
}

func Cmdworkforcemanagement_businessunits_activityplans_deletions() *cobra.Command {
	return workforcemanagement_businessunits_activityplans_deletionsCmd
}
