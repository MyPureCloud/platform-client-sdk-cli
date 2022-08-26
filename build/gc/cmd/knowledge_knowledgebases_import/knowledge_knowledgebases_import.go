package knowledge_knowledgebases_import

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("knowledge_knowledgebases_import", "SWAGGER_OVERRIDE_/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/import")
	knowledge_knowledgebases_importCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("knowledge_knowledgebases_import"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(knowledge_knowledgebases_importCmd)
}

func Cmdknowledge_knowledgebases_import() *cobra.Command {
	return knowledge_knowledgebases_importCmd
}
