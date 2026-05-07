package users_stations

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("users_stations", "SWAGGER_OVERRIDE_/api/v2/users/stations")
	users_stationsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("users_stations"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(users_stationsCmd)
}

func Cmdusers_stations() *cobra.Command {
	return users_stationsCmd
}
