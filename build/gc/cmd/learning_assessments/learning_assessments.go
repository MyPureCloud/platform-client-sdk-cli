package learning_assessments

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("learning_assessments", "SWAGGER_OVERRIDE_/api/v2/learning/assessments")
	learning_assessmentsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("learning_assessments"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(learning_assessmentsCmd)
}

func Cmdlearning_assessments() *cobra.Command {
	return learning_assessmentsCmd
}
