package speechandtextanalytics

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("speechandtextanalytics", "SWAGGER_OVERRIDE_/api/v2/speechandtextanalytics")
	speechandtextanalyticsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("speechandtextanalytics"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(speechandtextanalyticsCmd)
}

func Cmdspeechandtextanalytics() *cobra.Command {
	return speechandtextanalyticsCmd
}
