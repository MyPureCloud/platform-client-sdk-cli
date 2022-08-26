package knowledge_knowledgebases_export

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("knowledge_knowledgebases_export", "SWAGGER_OVERRIDE_/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/export")
	knowledge_knowledgebases_exportCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("knowledge_knowledgebases_export"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(knowledge_knowledgebases_exportCmd)
}

func Cmdknowledge_knowledgebases_export() *cobra.Command {
	return knowledge_knowledgebases_exportCmd
}
