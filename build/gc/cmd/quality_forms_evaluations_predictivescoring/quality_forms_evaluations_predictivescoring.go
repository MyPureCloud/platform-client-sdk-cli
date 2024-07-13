package quality_forms_evaluations_predictivescoring

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("quality_forms_evaluations_predictivescoring", "SWAGGER_OVERRIDE_/api/v2/quality/forms/evaluations/{formId}/predictivescoring")
	quality_forms_evaluations_predictivescoringCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("quality_forms_evaluations_predictivescoring"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(quality_forms_evaluations_predictivescoringCmd)
}

func Cmdquality_forms_evaluations_predictivescoring() *cobra.Command {
	return quality_forms_evaluations_predictivescoringCmd
}
