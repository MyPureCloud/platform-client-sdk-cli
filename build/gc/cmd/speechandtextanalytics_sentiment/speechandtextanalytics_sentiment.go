package speechandtextanalytics_sentiment

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("speechandtextanalytics_sentiment", "SWAGGER_OVERRIDE_/api/v2/speechandtextanalytics/sentiment")
	speechandtextanalytics_sentimentCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("speechandtextanalytics_sentiment"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(speechandtextanalytics_sentimentCmd)
}

func Cmdspeechandtextanalytics_sentiment() *cobra.Command {
	return speechandtextanalytics_sentimentCmd
}
