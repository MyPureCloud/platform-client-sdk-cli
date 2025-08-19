package speechandtextanalytics_topics_testphrase

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("speechandtextanalytics_topics_testphrase", "SWAGGER_OVERRIDE_/api/v2/speechandtextanalytics/topics/testphrase")
	speechandtextanalytics_topics_testphraseCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("speechandtextanalytics_topics_testphrase"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(speechandtextanalytics_topics_testphraseCmd)
}

func Cmdspeechandtextanalytics_topics_testphrase() *cobra.Command {
	return speechandtextanalytics_topics_testphraseCmd
}
