package knowledge_knowledgebases_parse

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("knowledge_knowledgebases_parse", "SWAGGER_OVERRIDE_/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/parse")
	knowledge_knowledgebases_parseCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("knowledge_knowledgebases_parse"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(knowledge_knowledgebases_parseCmd)
}

func Cmdknowledge_knowledgebases_parse() *cobra.Command {
	return knowledge_knowledgebases_parseCmd
}
