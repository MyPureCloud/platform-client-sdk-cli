package users_screenmonitors

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("users_screenmonitors", "SWAGGER_OVERRIDE_/api/v2/users/{userId}/screenmonitors")
	users_screenmonitorsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("users_screenmonitors"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(users_screenmonitorsCmd)
}

func Cmdusers_screenmonitors() *cobra.Command {
	return users_screenmonitorsCmd
}
