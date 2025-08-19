package knowledge_knowledgebases_documents_versions_bulk

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("knowledge_knowledgebases_documents_versions_bulk", "SWAGGER_OVERRIDE_/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/versions/bulk")
	knowledge_knowledgebases_documents_versions_bulkCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("knowledge_knowledgebases_documents_versions_bulk"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(knowledge_knowledgebases_documents_versions_bulkCmd)
}

func Cmdknowledge_knowledgebases_documents_versions_bulk() *cobra.Command {
	return knowledge_knowledgebases_documents_versions_bulkCmd
}
