package knowledge_knowledgebases_languages

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("knowledge_knowledgebases_languages", "SWAGGER_OVERRIDE_/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages")
	knowledge_knowledgebases_languagesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("knowledge_knowledgebases_languages"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(knowledge_knowledgebases_languagesCmd)
}

func Cmdknowledge_knowledgebases_languages() *cobra.Command {
	return knowledge_knowledgebases_languagesCmd
}
