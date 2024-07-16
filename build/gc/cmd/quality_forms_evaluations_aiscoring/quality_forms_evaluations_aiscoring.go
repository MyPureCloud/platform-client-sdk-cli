package quality_forms_evaluations_aiscoring

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("quality_forms_evaluations_aiscoring", "SWAGGER_OVERRIDE_/api/v2/quality/forms/evaluations/{formId}/aiscoring")
	quality_forms_evaluations_aiscoringCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("quality_forms_evaluations_aiscoring"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(quality_forms_evaluations_aiscoringCmd)
}

func Cmdquality_forms_evaluations_aiscoring() *cobra.Command {
	return quality_forms_evaluations_aiscoringCmd
}
