package agentic

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("agentic", "SWAGGER_OVERRIDE_/api/v2/agentic")
	agenticCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("agentic"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(agenticCmd)
}

func Cmdagentic() *cobra.Command {
	return agenticCmd
}
