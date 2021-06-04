package speechandtextanalytics_programs_general

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("speechandtextanalytics_programs_general", "SWAGGER_OVERRIDE_/api/v2/speechandtextanalytics/programs/general")
	speechandtextanalytics_programs_generalCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("speechandtextanalytics_programs_general"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(speechandtextanalytics_programs_generalCmd)
}

func Cmdspeechandtextanalytics_programs_general() *cobra.Command {
	return speechandtextanalytics_programs_generalCmd
}
