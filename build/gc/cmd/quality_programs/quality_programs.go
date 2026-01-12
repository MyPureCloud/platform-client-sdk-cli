package quality_programs

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("quality_programs", "SWAGGER_OVERRIDE_/api/v2/quality/programs")
	quality_programsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("quality_programs"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(quality_programsCmd)
}

func Cmdquality_programs() *cobra.Command {
	return quality_programsCmd
}
