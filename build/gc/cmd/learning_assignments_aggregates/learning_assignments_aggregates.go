package learning_assignments_aggregates

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("learning_assignments_aggregates", "SWAGGER_OVERRIDE_/api/v2/learning/assignments/aggregates")
	learning_assignments_aggregatesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("learning_assignments_aggregates"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(learning_assignments_aggregatesCmd)
}

func Cmdlearning_assignments_aggregates() *cobra.Command {
	return learning_assignments_aggregatesCmd
}
