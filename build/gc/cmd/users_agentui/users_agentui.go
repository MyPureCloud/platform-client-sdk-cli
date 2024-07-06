package users_agentui

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("users_agentui", "SWAGGER_OVERRIDE_/api/v2/users/agentui")
	users_agentuiCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("users_agentui"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(users_agentuiCmd)
}

func Cmdusers_agentui() *cobra.Command {
	return users_agentuiCmd
}
