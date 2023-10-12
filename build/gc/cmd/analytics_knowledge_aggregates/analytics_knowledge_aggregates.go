package analytics_knowledge_aggregates

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("analytics_knowledge_aggregates", "SWAGGER_OVERRIDE_/api/v2/analytics/knowledge/aggregates")
	analytics_knowledge_aggregatesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("analytics_knowledge_aggregates"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(analytics_knowledge_aggregatesCmd)
}

func Cmdanalytics_knowledge_aggregates() *cobra.Command {
	return analytics_knowledge_aggregatesCmd
}
