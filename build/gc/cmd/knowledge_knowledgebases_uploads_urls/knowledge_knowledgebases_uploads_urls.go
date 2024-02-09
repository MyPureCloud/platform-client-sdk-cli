package knowledge_knowledgebases_uploads_urls

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("knowledge_knowledgebases_uploads_urls", "SWAGGER_OVERRIDE_/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/uploads/urls")
	knowledge_knowledgebases_uploads_urlsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("knowledge_knowledgebases_uploads_urls"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(knowledge_knowledgebases_uploads_urlsCmd)
}

func Cmdknowledge_knowledgebases_uploads_urls() *cobra.Command {
	return knowledge_knowledgebases_uploads_urlsCmd
}
