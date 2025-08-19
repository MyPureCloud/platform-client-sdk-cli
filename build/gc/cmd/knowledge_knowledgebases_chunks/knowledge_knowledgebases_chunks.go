package knowledge_knowledgebases_chunks

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("knowledge_knowledgebases_chunks", "SWAGGER_OVERRIDE_/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/chunks")
	knowledge_knowledgebases_chunksCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("knowledge_knowledgebases_chunks"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(knowledge_knowledgebases_chunksCmd)
}

func Cmdknowledge_knowledgebases_chunks() *cobra.Command {
	return knowledge_knowledgebases_chunksCmd
}
