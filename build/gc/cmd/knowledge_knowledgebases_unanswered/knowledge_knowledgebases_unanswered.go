package knowledge_knowledgebases_unanswered

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("knowledge_knowledgebases_unanswered", "SWAGGER_OVERRIDE_/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/unanswered")
	knowledge_knowledgebases_unansweredCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("knowledge_knowledgebases_unanswered"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(knowledge_knowledgebases_unansweredCmd)
}

func Cmdknowledge_knowledgebases_unanswered() *cobra.Command {
	return knowledge_knowledgebases_unansweredCmd
}
