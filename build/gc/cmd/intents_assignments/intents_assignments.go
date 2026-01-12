package intents_assignments

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("intents_assignments", "SWAGGER_OVERRIDE_/api/v2/intents/assignments")
	intents_assignmentsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("intents_assignments"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(intents_assignmentsCmd)
}

func Cmdintents_assignments() *cobra.Command {
	return intents_assignmentsCmd
}
