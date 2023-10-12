package analytics_knowledge

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("analytics_knowledge", "SWAGGER_OVERRIDE_/api/v2/analytics/knowledge")
	analytics_knowledgeCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("analytics_knowledge"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(analytics_knowledgeCmd)
}

func Cmdanalytics_knowledge() *cobra.Command {
	return analytics_knowledgeCmd
}
