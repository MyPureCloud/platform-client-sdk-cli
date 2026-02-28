package knowledge_search

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("knowledge_search", "SWAGGER_OVERRIDE_/api/v2/knowledge/search")
	knowledge_searchCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("knowledge_search"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(knowledge_searchCmd)
}

func Cmdknowledge_search() *cobra.Command {
	return knowledge_searchCmd
}
