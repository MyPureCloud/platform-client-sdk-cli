package workforcemanagement_unavailabletimes_validation

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("workforcemanagement_unavailabletimes_validation", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/unavailabletimes/validation")
	workforcemanagement_unavailabletimes_validationCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement_unavailabletimes_validation"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagement_unavailabletimes_validationCmd)
}

func Cmdworkforcemanagement_unavailabletimes_validation() *cobra.Command {
	return workforcemanagement_unavailabletimes_validationCmd
}
