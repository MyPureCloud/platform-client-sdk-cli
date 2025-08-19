package flows_versions_intents_utterances

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("flows_versions_intents_utterances", "SWAGGER_OVERRIDE_/api/v2/flows/{flowId}/versions/{versionId}/intents/{intentId}/utterances")
	flows_versions_intents_utterancesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("flows_versions_intents_utterances"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(flows_versions_intents_utterancesCmd)
}

func Cmdflows_versions_intents_utterances() *cobra.Command {
	return flows_versions_intents_utterancesCmd
}
