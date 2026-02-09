package speechandtextanalytics_programs_topiclinks

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("speechandtextanalytics_programs_topiclinks", "SWAGGER_OVERRIDE_/api/v2/speechandtextanalytics/programs/topiclinks")
	speechandtextanalytics_programs_topiclinksCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("speechandtextanalytics_programs_topiclinks"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(speechandtextanalytics_programs_topiclinksCmd)
}

func Cmdspeechandtextanalytics_programs_topiclinks() *cobra.Command {
	return speechandtextanalytics_programs_topiclinksCmd
}
