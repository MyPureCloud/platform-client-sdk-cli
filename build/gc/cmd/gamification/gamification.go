package gamification

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("gamification", "SWAGGER_OVERRIDE_/api/v2/gamification")
	gamificationCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("gamification"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(gamificationCmd)
}

func Cmdgamification() *cobra.Command {
	return gamificationCmd
}
