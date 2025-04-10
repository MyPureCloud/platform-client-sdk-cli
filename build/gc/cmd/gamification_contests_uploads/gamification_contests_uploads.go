package gamification_contests_uploads

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("gamification_contests_uploads", "SWAGGER_OVERRIDE_/api/v2/gamification/contests/uploads")
	gamification_contests_uploadsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("gamification_contests_uploads"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(gamification_contests_uploadsCmd)
}

func Cmdgamification_contests_uploads() *cobra.Command {
	return gamification_contests_uploadsCmd
}
