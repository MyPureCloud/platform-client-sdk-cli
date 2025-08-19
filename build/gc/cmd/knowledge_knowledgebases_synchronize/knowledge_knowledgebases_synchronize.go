package knowledge_knowledgebases_synchronize

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("knowledge_knowledgebases_synchronize", "SWAGGER_OVERRIDE_/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/synchronize")
	knowledge_knowledgebases_synchronizeCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("knowledge_knowledgebases_synchronize"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(knowledge_knowledgebases_synchronizeCmd)
}

func Cmdknowledge_knowledgebases_synchronize() *cobra.Command {
	return knowledge_knowledgebases_synchronizeCmd
}
