package speechandtextanalytics_topics_generatedphrases

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("speechandtextanalytics_topics_generatedphrases", "SWAGGER_OVERRIDE_/api/v2/speechandtextanalytics/topics/generatedphrases")
	speechandtextanalytics_topics_generatedphrasesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("speechandtextanalytics_topics_generatedphrases"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(speechandtextanalytics_topics_generatedphrasesCmd)
}

func Cmdspeechandtextanalytics_topics_generatedphrases() *cobra.Command {
	return speechandtextanalytics_topics_generatedphrasesCmd
}
