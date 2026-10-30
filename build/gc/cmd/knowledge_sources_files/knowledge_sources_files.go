package knowledge_sources_files

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("knowledge_sources_files", "SWAGGER_OVERRIDE_/api/v2/knowledge/sources/files")
	knowledge_sources_filesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("knowledge_sources_files"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(knowledge_sources_filesCmd)
}

func Cmdknowledge_sources_files() *cobra.Command {
	return knowledge_sources_filesCmd
}
