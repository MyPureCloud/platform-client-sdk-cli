package knowledge_knowledgebases_uploads

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("knowledge_knowledgebases_uploads", "SWAGGER_OVERRIDE_/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/uploads")
	knowledge_knowledgebases_uploadsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("knowledge_knowledgebases_uploads"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(knowledge_knowledgebases_uploadsCmd)
}

func Cmdknowledge_knowledgebases_uploads() *cobra.Command {
	return knowledge_knowledgebases_uploadsCmd
}
