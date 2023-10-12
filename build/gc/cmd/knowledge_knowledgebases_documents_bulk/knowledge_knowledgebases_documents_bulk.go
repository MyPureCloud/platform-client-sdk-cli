package knowledge_knowledgebases_documents_bulk

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("knowledge_knowledgebases_documents_bulk", "SWAGGER_OVERRIDE_/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/bulk")
	knowledge_knowledgebases_documents_bulkCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("knowledge_knowledgebases_documents_bulk"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(knowledge_knowledgebases_documents_bulkCmd)
}

func Cmdknowledge_knowledgebases_documents_bulk() *cobra.Command {
	return knowledge_knowledgebases_documents_bulkCmd
}
