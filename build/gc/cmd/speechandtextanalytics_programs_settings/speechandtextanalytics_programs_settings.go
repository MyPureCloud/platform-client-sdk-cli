package speechandtextanalytics_programs_settings

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("speechandtextanalytics_programs_settings", "SWAGGER_OVERRIDE_/api/v2/speechandtextanalytics/programs/{programId}/settings")
	speechandtextanalytics_programs_settingsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("speechandtextanalytics_programs_settings"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(speechandtextanalytics_programs_settingsCmd)
}

func Cmdspeechandtextanalytics_programs_settings() *cobra.Command {
	return speechandtextanalytics_programs_settingsCmd
}
