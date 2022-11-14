package workforcemanagement_businessunits_adherence

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("workforcemanagement_businessunits_adherence", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/businessunits/{businessUnitId}/adherence")
	workforcemanagement_businessunits_adherenceCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement_businessunits_adherence"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagement_businessunits_adherenceCmd)
}

func Cmdworkforcemanagement_businessunits_adherence() *cobra.Command {
	return workforcemanagement_businessunits_adherenceCmd
}
