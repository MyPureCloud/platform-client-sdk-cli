package knowledge_knowledgebases_operations_users

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("knowledge_knowledgebases_operations_users", "SWAGGER_OVERRIDE_/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/operations/users")
	knowledge_knowledgebases_operations_usersCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("knowledge_knowledgebases_operations_users"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(knowledge_knowledgebases_operations_usersCmd)
}

func Cmdknowledge_knowledgebases_operations_users() *cobra.Command {
	return knowledge_knowledgebases_operations_usersCmd
}
