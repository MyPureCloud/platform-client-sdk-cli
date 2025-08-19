package speechandtextanalytics_translations

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("speechandtextanalytics_translations", "SWAGGER_OVERRIDE_/api/v2/speechandtextanalytics/translations")
	speechandtextanalytics_translationsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("speechandtextanalytics_translations"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(speechandtextanalytics_translationsCmd)
}

func Cmdspeechandtextanalytics_translations() *cobra.Command {
	return speechandtextanalytics_translationsCmd
}
