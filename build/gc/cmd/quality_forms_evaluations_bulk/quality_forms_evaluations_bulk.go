package quality_forms_evaluations_bulk

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("quality_forms_evaluations_bulk", "SWAGGER_OVERRIDE_/api/v2/quality/forms/evaluations/bulk")
	quality_forms_evaluations_bulkCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("quality_forms_evaluations_bulk"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(quality_forms_evaluations_bulkCmd)
}

func Cmdquality_forms_evaluations_bulk() *cobra.Command {
	return quality_forms_evaluations_bulkCmd
}
