package flows_versions_intents

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("flows_versions_intents", "SWAGGER_OVERRIDE_/api/v2/flows/{flowId}/versions/{versionId}/intents")
	flows_versions_intentsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("flows_versions_intents"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(flows_versions_intentsCmd)
}

func Cmdflows_versions_intents() *cobra.Command {
	return flows_versions_intentsCmd
}
