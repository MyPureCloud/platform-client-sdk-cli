package profile

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("profile", "SWAGGER_OVERRIDE_/api/v2/profiles")
	profileCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("profile"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(profileCmd)
}

func Cmdprofile() *cobra.Command {
	return profileCmd
}
