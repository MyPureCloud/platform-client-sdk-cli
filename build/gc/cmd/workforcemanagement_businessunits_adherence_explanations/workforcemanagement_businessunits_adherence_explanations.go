package workforcemanagement_businessunits_adherence_explanations

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("workforcemanagement_businessunits_adherence_explanations", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/businessunits/{businessUnitId}/adherence/explanations")
	workforcemanagement_businessunits_adherence_explanationsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement_businessunits_adherence_explanations"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagement_businessunits_adherence_explanationsCmd)
}

func Cmdworkforcemanagement_businessunits_adherence_explanations() *cobra.Command {
	return workforcemanagement_businessunits_adherence_explanationsCmd
}
