package knowledge

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("knowledge", "SWAGGER_OVERRIDE_/api/v2/knowledge")
	knowledgeCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("knowledge"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(knowledgeCmd)
}

func Cmdknowledge() *cobra.Command {
	return knowledgeCmd
}
