package speechandtextanalytics_reprocessing

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("speechandtextanalytics_reprocessing", "SWAGGER_OVERRIDE_/api/v2/speechandtextanalytics/reprocessing")
	speechandtextanalytics_reprocessingCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("speechandtextanalytics_reprocessing"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(speechandtextanalytics_reprocessingCmd)
}

func Cmdspeechandtextanalytics_reprocessing() *cobra.Command {
	return speechandtextanalytics_reprocessingCmd
}
